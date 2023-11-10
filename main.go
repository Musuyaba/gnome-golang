package main

import (
	"context"
	"fmt"
	"log"
	"time"

	mongoController "github.com/Musuyaba/gnome-golang/mongo/Controllers"
	"github.com/Musuyaba/gnome-golang/pkg/initializers"
	"github.com/Musuyaba/gnome-golang/sqlc/sqlcControllers"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// server *gin.Engine
	ctx context.Context
)

func init() {
	ctx = context.Background()
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// initializers.ConnectDbGormPostgreSQL(&config)
	initializers.ConnectDbSqlcMySQL(&config)
	initializers.ConnectDbMongo(&config, ctx)

	// server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	dbMongoInstance := initializers.MongoInstance.Database(config.DATABASE_MONGODB_DATA_NODE_A_NAME)
	collection := dbMongoInstance.Collection("qr-flatten")

	for i := 0; i < 1; i++ {
		startTime := time.Now()
		runMigrate(collection)
		executionTime := time.Since(startTime)
		fmt.Println("Execution time:", executionTime)
	}

	defer initializers.MongoInstance.Disconnect(ctx)

	// // routers.ApiRoutes(server)
	// // routers.PublicRoutes(server)

	// // log.Fatal(server.RunTLS(config.HOSTNAME+":"+config.PORT, "./certs/generated/server.crt", "./certs/generated/server.key"))
}

func runMigrate(collection *mongo.Collection) {
	rowDataPrint, err := sqlcControllers.FirstAndupdateDataPrint(ctx, initializers.MySqlcInstance)
	if err != nil {
		log.Fatal(err)
	}

	palleteObject := mongoController.CreatePalleteObject(rowDataPrint)
	koliObject := mongoController.CreateKoliObject(rowDataPrint, palleteObject)
	wrapObject := mongoController.CreateWrapObject(rowDataPrint, koliObject)
	pieceObject := mongoController.CreatePieceObject(rowDataPrint, wrapObject)

	if _, err := mongoController.InsertPiece(collection, ctx, pieceObject); err != nil {
		log.Fatal(err)
	}
}
