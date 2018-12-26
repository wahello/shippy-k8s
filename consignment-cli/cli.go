package main

import (
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
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
	var token string
	log.Println(os.Args)
	token = os.Args[1]
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
