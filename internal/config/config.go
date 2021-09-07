package config

import "github.com/spf13/viper"

type PostgresConfig struct {
	// Адрес БД
	DbHost string
	// Порт БД
	DbPort int
	// Название БД
	DbName string
	// Имя пользователя БД
	DbUserName string
	// Пароль пользователя БД
	DbPassword string
}

type KafkaConfig struct {
	KafkaHost string
	KafkaPort int
	Topic     string
}

type MetricsConfig struct {
	MetricsPath string
	MetricsPort int
}

type TracerConfig struct {
	// Адрес Jaeger
	JaegerHost string
	// Порт подключения к Jaeger
	JaegerPort int
}

func ParsePostgresConfigFromViper() *PostgresConfig {
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetInt("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	dbUserName := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	return &PostgresConfig{
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUserName: dbUserName,
		DbPassword: dbPassword,
	}
}

func ParseTracerConfigFromViper() *TracerConfig {
	return &TracerConfig{
		JaegerHost: viper.GetString("JAEGER_HOST"),
		JaegerPort: viper.GetInt("JAEGER_PORT"),
	}
}

func ParseKafkaConfigFromViper() *KafkaConfig {
	return &KafkaConfig{
		KafkaHost: viper.GetString("KAFKA_HOST"),
		KafkaPort: viper.GetInt("KAFKA_PORT"),
		Topic:     viper.GetString("KAFKA_TOPIC"),
	}
}

func ParseMetricsConfigFromViper() *MetricsConfig {
	return &MetricsConfig{
		MetricsPath: viper.GetString("METRICS_PATH"),
		MetricsPort: viper.GetInt("METRICS_PORT"),
	}
}
