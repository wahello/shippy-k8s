package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	k8s "github.com/micro/kubernetes/go/micro"
	"golang.org/x/net/context"

	authService "github.com/cgault/shippy/auth-service/proto/auth"
	pb "github.com/cgault/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/cgault/shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

var (
	srv micro.Service
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	srv = k8s.NewService(
		micro.Name("shippy.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)
	vesselClient := vesselProto.NewVesselServiceClient("shippy.vessel", srv.Client())
	srv.Init()
	pb.RegisterConsignmentServiceHandler(srv.Server(), &service{session, vesselClient})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)
		authClient := authService.NewAuthClient("shippy.user", srv.Client())
		_, err := authClient.ValidateToken(ctx, &authService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
