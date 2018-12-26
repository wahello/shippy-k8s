package main

import (
	"log"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"

	pb "github.com/cgault/shippy/consignment-service/proto/consignment"
)

func main() {
	cmd.Init()
	client := pb.NewShippingService("go.micro.srv.consignment", microclient.DefaultClient)
	consignment := &pb.Consignment{
		Description: "This is a test consignment",
		Weight:      55000,
		Containers: []*pb.Container{
			{
				CustomerId: "cust001",
				UserId:     "user001",
				Origin:     "Manchester, United Kingdom",
			},
			{
				CustomerId: "cust002",
				UserId:     "user001",
				Origin:     "Derby, United Kingdom",
			},
			{
				CustomerId: "cust005",
				UserId:     "user001",
				Origin:     "Sheffield, United Kingdom",
			},
		},
	}
	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}
	log.Printf("created: %t", r.Created)
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
