package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/cgault/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/cgault/shippy/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
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
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
