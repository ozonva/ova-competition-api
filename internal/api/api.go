package api

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"io"
	"net"
	"ozonva/ova-competition-api/internal/config"
	"ozonva/ova-competition-api/internal/kafka"
	"ozonva/ova-competition-api/internal/metrics"
	"ozonva/ova-competition-api/internal/repo"
	"ozonva/ova-competition-api/internal/tracer"
	desc "ozonva/ova-competition-api/pkg/competition/api"
)

const (
	grpcPort   = ":1489"
	configPath = ".env"
)

type Server struct {
	desc.UnimplementedCompetitionServiceServer
	competitionRepo *repo.Repo
	kafkaClient     *kafka.Client
	metrics         *metrics.CompetitionMetrics
}

func newServer(dbConfig *config.PostgresConfig, kafkaConfig *config.KafkaConfig, competitionMetrics *metrics.CompetitionMetrics) (desc.CompetitionServiceServer, error) {
	db, err := repo.NewDb(dbConfig)
	if err != nil {
		return nil, err
	}

	competitionRepo := repo.NewRepo(db)
	kafkaClient := kafka.NewClient(fmt.Sprintf("%s:%d", kafkaConfig.KafkaHost, kafkaConfig.KafkaPort), kafkaConfig.Topic)

	return &Server{
		competitionRepo: &competitionRepo,
		kafkaClient:     &kafkaClient,
		metrics:         competitionMetrics,
	}, nil
}

func RunGrpcServer() error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return errors.New(fmt.Sprintf("failed to read config: %v", err))
	}

	dbConfig := config.ParsePostgresConfigFromViper()
	tracerConfig := config.ParseTracerConfigFromViper()
	kafkaConfig := config.ParseKafkaConfigFromViper()
	metricsConfig := config.ParseMetricsConfigFromViper()

	competitionMetrics := metrics.NewMetrics("competitions_api", "Server")

	grpcServer := grpc.NewServer()
	competitionServer, err := newServer(dbConfig, kafkaConfig, &competitionMetrics)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to start new server: %v", err))
	}

	closer, err := tracer.InitTracer("Competitions service", tracerConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to initialize tracer: %v", err))
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Errorf("failed to close tracer: %v", err)
		}
	}(closer)

	desc.RegisterCompetitionServiceServer(grpcServer, competitionServer)

	go func() {
		err := metrics.RunMetricsServer(metricsConfig)
		if err != nil {
			log.Errorf("failed to run metrics server: %v", err)
		}
	}()

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}
	if err := grpcServer.Serve(listen); err != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", err))
	}

	return nil
}
