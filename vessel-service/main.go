package main

import (
	"fmt"
	"os"

	pb "github.com/cgault/shippy/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
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
	shard1 := os.Getenv("MONGO_SHARD_1")
	shard2 := os.Getenv("MONGO_SHARD_2")
	session, err := CreateSession(shard1, shard2)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	fmt.Printf("Connected to %v!\n", session.LiveServers())
	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)
	srv := k8s.NewService(
		micro.Name("shippy.vessel"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
