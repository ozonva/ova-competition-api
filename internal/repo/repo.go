package repo

import (
	"context"
	"ozonva/ova-competition-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Competition
type Repo interface {
	// AddEntities сохраняет несколько соревнований
	AddEntities(ctx context.Context, entities []models.Competition) error
	// ListEntities выводит список соревнований размера limit со смещением offset
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Competition, error)
	// DescribeEntity выводит информацию о заданном соревновании
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Competition, error)
	// UpdateEntity обновляет информацию об имеющемся соревновании
	UpdateEntity(ctx context.Context, entityId uint64, competition *models.Competition) error
	// RemoveEntity удаляет заданное соревнование
	RemoveEntity(ctx context.Context, entityId uint64) error
}
