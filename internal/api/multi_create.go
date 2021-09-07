package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ozonva/ova-competition-api/internal/models"
	"ozonva/ova-competition-api/internal/utils"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

func (s *Server) MultiCreateCompetitions(ctx context.Context, req *desc.MultiCreateCompetitionsRequest) (*desc.MultiCreateCompetitionsResponse, error) {
	log.Infof("Creating multiple competitions: %v", req)
	defer (*s.metrics).MultiCreated()

	competitions := make([]models.Competition, 0, len(req.Competitions))
	for _, c := range req.Competitions {
		competitions = append(competitions, models.NewCompetition(c.Id, c.Name, c.CreateDate.AsTime()))
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "MultiCreateCompetitions")
	defer span.Finish()

	competitionSlices, err := utils.CompetitionSliceToBatches(competitions, int(req.BatchSize))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, competitionsSlice := range competitionSlices {
		err := s.insertWithTracing(ctx, span, competitionsSlice)
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &desc.MultiCreateCompetitionsResponse{
		CreatedCompetitions: uint64(len(competitions)),
	}, nil
}

func (s *Server) insertWithTracing(ctx context.Context, parentSpan opentracing.Span, competitions []models.Competition) error {
	batchSpan := opentracing.StartSpan("MultiCreateCompetitionsBatch", opentracing.ChildOf(parentSpan.Context()))
	defer batchSpan.Finish()
	if err := (*s.competitionRepo).AddEntities(ctx, competitions); err != nil {
		return err
	}
	return nil
}
