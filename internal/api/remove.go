package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) RemoveCompetition(_ context.Context, req *desc.RemoveCompetitionRequest) (*emptypb.Empty, error) {
	log.Infof("Removing competition: %v", req)
	err := (*s.competitionRepo).RemoveEntity(req.CompetitionId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
