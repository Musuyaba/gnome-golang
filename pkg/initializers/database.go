package initializers

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgreGormInstance *gorm.DB
	MySqlcInstance      *sql.DB
	MongoInstance       *mongo.Client
	err                 error
)

func ConnectDbGormPostgreSQL(config *Config) {
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DATABASE_POSTGRESQL_HOST, config.DATABASE_POSTGRESQL_USER, config.DATABASE_POSTGRESQL_PASSWORD, config.DATABASE_POSTGRESQL_NAME, config.DATABASE_POSTGRESQL_PORT)

	PostgreGormInstance, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
		log.Fatalln(err)
	}
}

func ConnectDbSqlcMySQL(config *Config) {
	MySqlcInstance, err = sql.Open("mysql", config.DATABASE_MYSQL_USER+":"+config.DATABASE_MYSQL_PASSWORD+"@("+config.DATABASE_MYSQL_HOST+":"+config.DATABASE_MYSQL_PORT+")/"+config.DATABASE_MYSQL_NAME+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDbMongo(config *Config, ctx context.Context) {
	MongoInstance, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+config.DATABASE_MONGODB_USER+":"+config.DATABASE_MONGODB_PASSWORD+"@"+config.DATABASE_MONGODB_HOST+":"+config.DATABASE_MONGODB_PORT))
	if err != nil {
		log.Fatal(err)
	}
}
