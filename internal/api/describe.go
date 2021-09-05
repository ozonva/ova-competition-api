package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) DescribeCompetition(_ context.Context, req *desc.DescribeCompetitionRequest) (*desc.CompetitionResponse, error) {
	log.Infof("Describing competition: %v", req)
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
