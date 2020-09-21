package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tkanos/gonfig"
)

var (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbName   = ""
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}

	fileName := fmt.Sprintf("./config.json")
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

func Connect() (*sql.DB, error) {
	configuration := GetConfig()

	user = configuration.DB_USERNAME
	password = configuration.DB_PASSWORD
	dbName = configuration.DB_NAME

	var connectionString string = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
