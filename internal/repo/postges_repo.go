package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"ozonva/ova-competition-api/internal/config"
	"ozonva/ova-competition-api/internal/models"
)

type repo struct {
	db *sqlx.DB
}

func NewDb(dbConfig *config.PostgresConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbUserName, dbConfig.DbPassword, dbConfig.DbName)
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open postgres connection: %v", err))
	}

	return db, nil
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{db: db}
}

func (r *repo) AddEntities(ctx context.Context, entities []models.Competition) error {
	queryBuilder := squirrel.
		Insert("competitions").
		Columns("id", "name", "start_time", "status")
	for _, competition := range entities {
		queryBuilder = queryBuilder.Values(competition.Id, competition.Name, competition.StartTime, competition.Status())
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	_, execErr := r.db.ExecContext(ctx, query, args...)
	if execErr != nil {
		return execErr
	}

	return nil
}

func (r *repo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Competition, error) {
	result := make([]models.Competition, 0, limit)
	sql, args, err := squirrel.
		Select("id", "name", "start_time", "status").
		From("competitions").
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.SelectContext(ctx, &result, sql, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repo) DescribeEntity(ctx context.Context, entityId uint64) (*models.Competition, error) {
	var result *models.Competition
	query, args, err := squirrel.
		Select("id", "name", "start_time", "status").
		From("competitions").
		Where(squirrel.Eq{"id": entityId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.GetContext(ctx, &result, query, args)
	return result, err
}

func (r *repo) RemoveEntity(ctx context.Context, entityId uint64) error {
	sql, args, err := squirrel.
		Delete("competitions").
		Where(squirrel.Eq{"id": entityId}).
		ToSql()
	if err != nil {
		return err
	}

	_, execErr := r.db.ExecContext(ctx, sql, args...)
	if execErr != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateEntity(ctx context.Context, entityId uint64, competition *models.Competition) error {
	sql, args, err := squirrel.
		Update("competitions").
		Set("name", competition.Name).
		Set("start_time", competition.StartTime).
		Set("status", competition.Status()).
		Where(squirrel.Eq{"id": entityId}).
		ToSql()
	if err != nil {
		return err
	}

	_, execErr := r.db.ExecContext(ctx, sql, args...)
	if execErr != nil {
		return err
	}

	return nil
}
