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
