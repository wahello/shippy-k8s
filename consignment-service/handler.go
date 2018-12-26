package main

import (
	"log"

	pb "github.com/cgault/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/cgault/shippy/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselService
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id
	err = repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
