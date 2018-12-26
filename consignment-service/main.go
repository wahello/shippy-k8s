// Package main provides ...
package main

import (
	"context"
	"log"

	pb "github.com/cgault/shippy/consignment-service/proto/consignment"

	micro "github.com/micro/go-micro"
)

const (
	port string = ":50051"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type repository struct {
	consignments []*pb.Consignment
}

func (r *repository) Create(c *pb.Consignment) (*pb.Consignment, error) {
	u := append(r.consignments, c)
	r.consignments = u
	return c, nil
}

func (r *repository) GetAll() []*pb.Consignment {
	return r.consignments
}

type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	c, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = c
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	repo := &repository{}
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
