package repo

import (
	"context"
	"ozonva/ova-competition-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Competition
type Repo interface {
	AddEntities(ctx context.Context, entities []models.Competition) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Competition, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Competition, error)
	UpdateEntity(ctx context.Context, entityId uint64, competition *models.Competition) error
	RemoveEntity(ctx context.Context, entityId uint64) error
}
