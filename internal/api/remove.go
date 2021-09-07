package api

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"ozonva/ova-competition-api/internal/kafka"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) RemoveCompetition(ctx context.Context, req *desc.RemoveCompetitionRequest) (*emptypb.Empty, error) {
	log.Infof("Removing competition: %v", req)
	defer (*s.metrics).RemoveCompetition()

	existingCompetition, err := (*s.competitionRepo).DescribeEntity(ctx, req.CompetitionId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if existingCompetition == nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("failed to find competition %d", req.CompetitionId))
	}

	err = (*s.competitionRepo).RemoveEntity(ctx, req.CompetitionId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err := (*s.kafkaClient).Send(ctx, kafka.NewCompetitionEvent(kafka.CompetitionDeleted, existingCompetition)); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
