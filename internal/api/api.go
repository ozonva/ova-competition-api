package api

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"ozonva/ova-competition-api/internal/config"
	"ozonva/ova-competition-api/internal/repo"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

const (
	grpcPort   = ":1489"
	configPath = ".env"
)

type Server struct {
	desc.UnimplementedCompetitionServiceServer
	competitionRepo *repo.Repo
}

func newServer(dbConfig *config.PostgresConfig) (desc.CompetitionServiceServer, error) {
	db, err := repo.NewDb(dbConfig)
	if err != nil {
		return nil, err
	}

	competitionRepo := repo.NewRepo(db)
	return &Server{
		competitionRepo: &competitionRepo,
	}, nil
}

func RunGrpcServer() error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return errors.New(fmt.Sprintf("failed to read config: %v", err))
	}

	dbConfig := config.ParsePostgresConfigFromViper()

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()
	competitionServer, err := newServer(dbConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to start new server: %v", err))
	}
	desc.RegisterCompetitionServiceServer(grpcServer, competitionServer)

	if err := grpcServer.Serve(listen); err != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", err))
	}

	return nil
}
