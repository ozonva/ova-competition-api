package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) CreateCompetition(ctx context.Context, req *desc.CreateCompetitionRequest) (*desc.CompetitionResponse, error) {
	log.Infof("Creating competition: %v", req)
	return &desc.CompetitionResponse{
		Id:         req.Id,
		Name:       req.Name,
		CreateDate: req.CreateDate,
		Status:     desc.CompetitionStatus_Planned,
	}, nil
}
