package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	desc "ozonva/ova-competition-api/pkg/competition/api"
	"time"
)

func (s *Server) ListCompetitions(ctx context.Context, req *desc.ListCompetitionsRequest) (*desc.ListCompetitionsResponse, error) {
	log.Infof("Listing competitions: %v", req)
	return &desc.ListCompetitionsResponse{
		Competitions: []*desc.CompetitionResponse{
			{
				Id:         1,
				Name:       "Test competition 1",
				CreateDate: timestamppb.New(time.Now()),
				Status:     desc.CompetitionStatus_Finished,
			},
			{
				Id:         2,
				Name:       "Test competition 2",
				CreateDate: timestamppb.New(time.Now()),
				Status:     desc.CompetitionStatus_Pending,
			},
		},
	}, nil
}
