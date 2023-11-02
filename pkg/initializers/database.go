package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DbInstance *gorm.DB
	err        error
)

func ConnectDbPostgreSQL(config *Config) {
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DATABASE_POSTGRESQL_HOST, config.DATABASE_POSTGRESQL_USER, config.DATABASE_POSTGRESQL_PASSWORD, config.DATABASE_POSTGRESQL_NAME, config.DATABASE_POSTGRESQL_PORT)

	DbInstance, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
		log.Fatalln(err)
	}

}
