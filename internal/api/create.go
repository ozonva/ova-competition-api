package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ozonva/ova-competition-api/internal/kafka"
	"ozonva/ova-competition-api/internal/models"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) CreateCompetition(ctx context.Context, req *desc.CreateCompetitionRequest) (*desc.CompetitionResponse, error) {
	log.Infof("Creating competition: %v", req)
	defer (*s.metrics).CreateCompetition()

	err := (*s.competitionRepo).AddEntities(ctx, []models.Competition{
		models.NewCompetition(req.Id, req.Name, req.CreateDate.AsTime()),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	competition, err := (*s.competitionRepo).DescribeEntity(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err := (*s.kafkaClient).Send(ctx, kafka.NewCompetitionEvent(kafka.CompetitionCreated, competition)); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CompetitionResponse{
		Id:         competition.Id,
		Name:       competition.Name,
		CreateDate: timestamppb.New(competition.StartTime),
		Status:     desc.CompetitionStatus(int32(competition.Status())),
	}, nil
}
