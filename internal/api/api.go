package api

import (
	"google.golang.org/grpc"
	"log"
	"net"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

const (
	grpcPort = ":1489"
)

type Server struct {
	desc.UnimplementedCompetitionServiceServer
}

func newServer() desc.CompetitionServiceServer {
	return &Server{}
}

func RunGrpcServer() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterCompetitionServiceServer(s, newServer())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
