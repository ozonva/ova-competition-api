package repo

import (
	"ozonva/ova-competition-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Competition
type Repo interface {
	AddEntities(entities []models.Competition) error
	ListEntities(limit, offset uint64) ([]models.Competition, error)
	DescribeEntity(entityId uint64) (*models.Competition, error)
	RemoveEntity(entityId uint64) error
}
