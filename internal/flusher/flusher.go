package flusher

import (
	"context"
	"ozonva/ova-competition-api/internal/models"
	"ozonva/ova-competition-api/internal/repo"
	"ozonva/ova-competition-api/internal/utils"
)

type flusher struct {
	chunkSize       int
	competitionRepo repo.Repo
}

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	// Flush сбрасывает соревнования в хранилища с разбиением на батчи
	Flush(ctx context.Context, entities []models.Competition) []models.Competition
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	competitionRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:       chunkSize,
		competitionRepo: competitionRepo,
	}
}

func (f *flusher) Flush(ctx context.Context, competitions []models.Competition) []models.Competition {
	batches, err := utils.CompetitionSliceToBatches(competitions, f.chunkSize)
	if err != nil {
		return competitions
	}

	failedToFlush := make([]models.Competition, 0, len(competitions))
	for _, batch := range batches {
		err := f.competitionRepo.AddEntities(ctx, batch)
		if err != nil {
			failedToFlush = append(failedToFlush, batch...)
		}
	}

	if len(failedToFlush) > 0 {
		return failedToFlush
	}

	return nil
}
