package main

import (
	"log"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	k8s "github.com/micro/kubernetes/go/micro"

	pb "github.com/cgault/shippy/auth-service/proto/auth"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenService := &TokenService{repo}
	srv := k8s.NewService(
		micro.Name("shippy.auth"),
	)
	srv.Init()
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}