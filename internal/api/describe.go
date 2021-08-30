package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	desc "ozonva/ova-competition-api/pkg/competition/api"
	"time"
)

func (s *Server) DescribeCompetition(ctx context.Context, req *desc.DescribeCompetitionRequest) (*desc.CompetitionResponse, error) {
	log.Infof("Describing competition: %v", req)
	return &desc.CompetitionResponse{
		Id:         1,
		Name:       "Test competition",
		CreateDate: timestamppb.New(time.Now()),
		Status:     desc.CompetitionStatus_Pending,
	}, nil
}
