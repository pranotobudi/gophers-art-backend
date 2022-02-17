package config

import (
	"github.com/revel/revel"
)

type DbConnection struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	SSLMode  string
}

// func DbConfig() DbConnection {
// 	dbConfig := DbConnection{
// 		Host:     os.Getenv("DB_HOST"),
// 		Port:     os.Getenv("DB_PORT"),
// 		DbName:   os.Getenv("DB_NAME"),
// 		Username: os.Getenv("DB_USERNAME"),
// 		Password: os.Getenv("DB_PASSWORD"),
// 		SSLMode:  os.Getenv("DB_SSL_MODE"),
// 	}

// 	return dbConfig
// }

func DbConfig() DbConnection {
	host, _ := revel.Config.String("DB_HOST")
	port, _ := revel.Config.String("DB_PORT")
	dbName, _ := revel.Config.String("DB_NAME")
	username, _ := revel.Config.String("DB_USERNAME")
	password, _ := revel.Config.String("DB_PASSWORD")
	sslMode, _ := revel.Config.String("DB_SSL_MODE")

	dbConfig := DbConnection{
		Host:     host,
		Port:     port,
		DbName:   dbName,
		Username: username,
		Password: password,
		SSLMode:  sslMode,
	}

	return dbConfig
}

type AppEnvironment struct {
	Port    string
	AppEnv  string
	AppHost string
}

func AppConfig() AppEnvironment {
	appHost, _ := revel.Config.String("APP_HOST")
	port, _ := revel.Config.String("PORT")
	appEnv, _ := revel.Config.String("APP_ENV")

	appConfig := AppEnvironment{
		AppHost: appHost,
		Port:    port,
		AppEnv:  appEnv,
	}

	return appConfig
}
