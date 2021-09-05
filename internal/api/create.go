package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ozonva/ova-competition-api/internal/models"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) CreateCompetition(_ context.Context, req *desc.CreateCompetitionRequest) (*desc.CompetitionResponse, error) {
	log.Infof("Creating competition: %v", req)
	err := (*s.competitionRepo).AddEntities([]models.Competition{
		models.NewCompetition(req.Id, req.Name, req.CreateDate.AsTime()),
	})

	if err != nil {
		return nil, err
	}

	competition, err := (*s.competitionRepo).DescribeEntity(req.Id)
	if err != nil {
		return nil, err
	}

	return &desc.CompetitionResponse{
		Id:         competition.Id,
		Name:       competition.Name,
		CreateDate: timestamppb.New(competition.StartTime),
		Status:     desc.CompetitionStatus(int32(competition.Status())),
	}, nil
}
