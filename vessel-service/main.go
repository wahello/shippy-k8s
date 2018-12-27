package main

import (
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	pb "github.com/cgault/shippy/vessel-service/proto/vessel"
)

const (
	serviceName    = "shippy.vessel"
	serviceVersion = "latest"
	defaultHost    = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatalf("Error connecting to datastore %s: %v", host, err)
	}
	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)
	srv := k8s.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
