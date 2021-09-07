package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) ListCompetitions(ctx context.Context, req *desc.ListCompetitionsRequest) (*desc.ListCompetitionsResponse, error) {
	log.Infof("Listing competitions: %v", req)
	defer (*s.metrics).ListCompetitions()

	competitions, err := (*s.competitionRepo).ListEntities(ctx, uint64(req.Limit), uint64(req.Offset))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := make([]*desc.CompetitionResponse, 0, len(competitions))
	for _, competition := range competitions {
		res = append(res, &desc.CompetitionResponse{
			Id:         competition.Id,
			Name:       competition.Name,
			CreateDate: timestamppb.New(competition.StartTime),
			Status:     desc.CompetitionStatus(int32(competition.Status())),
		})
	}

	return &desc.ListCompetitionsResponse{
		Competitions: res,
	}, nil
}
