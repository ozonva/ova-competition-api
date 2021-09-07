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

func (s *Server) UpdateCompetition(ctx context.Context, req *desc.UpdateCompetitionRequest) (*desc.CompetitionResponse, error) {
	log.Infof("Updating competition: %v", req)
	defer (*s.metrics).UpdateCompetition()

	comp := models.NewCompetition(req.Competition.Id, req.Competition.Name, req.Competition.CreateDate.AsTime())
	err := (*s.competitionRepo).UpdateEntity(ctx, req.Competition.Id, &comp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err := (*s.kafkaClient).Send(ctx, kafka.NewCompetitionEvent(kafka.CompetitionUpdated, &comp)); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CompetitionResponse{
		Id:         comp.Id,
		Name:       comp.Name,
		CreateDate: timestamppb.New(comp.StartTime),
		Status:     desc.CompetitionStatus(int32(comp.Status())),
	}, nil
}
