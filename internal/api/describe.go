package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) DescribeCompetition(ctx context.Context, req *desc.DescribeCompetitionRequest) (*desc.CompetitionResponse, error) {
	log.Infof("Describing competition: %v", req)
	defer (*s.metrics).DescribeCompetition()

	competition, err := (*s.competitionRepo).DescribeEntity(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CompetitionResponse{
		Id:         competition.Id,
		Name:       competition.Name,
		CreateDate: timestamppb.New(competition.StartTime),
		Status:     desc.CompetitionStatus(int32(competition.Status())),
	}, nil
}
