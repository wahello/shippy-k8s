package main

import (
	"context"
	"log"

	pb "github.com/cgault/shippy/auth-service/proto/auth"
	micro "github.com/micro/go-micro"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)
	srv.Init()
	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
