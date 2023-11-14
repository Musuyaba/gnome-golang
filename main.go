package main

import (
	"context"
	"fmt"
	"log"
	"time"

	mongocontroller "github.com/Musuyaba/gnome-golang/mongo/Controllers"
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

	for i := 0; i < 1; i++ {
		startTime := time.Now()
		runMigrate(dbMongoInstance, "qr-flatten-pieces")
		executionTime := time.Since(startTime)
		fmt.Println("Execution time:", executionTime)
	}

	defer initializers.MongoInstance.Disconnect(ctx)

	// // routers.ApiRoutes(server)
	// // routers.PublicRoutes(server)

	// // log.Fatal(server.RunTLS(config.HOSTNAME+":"+config.PORT, "./certs/generated/server.crt", "./certs/generated/server.key"))
}

func runMigrate(dbMongoInstance *mongo.Database, collection_name string) {
	rowDataPrint, err := sqlcControllers.FirstAndupdateDataPrint(ctx, initializers.MySqlcInstance)
	if err != nil {
		log.Fatal(err)
	}

	palleteObject := mongocontroller.CreatePalleteObject(rowDataPrint)
	koliObject := mongocontroller.CreateKoliObject(rowDataPrint, palleteObject)
	wrapObject := mongocontroller.CreateWrapObject(rowDataPrint, koliObject)
	pieceObject := mongocontroller.CreatePieceObject(rowDataPrint, wrapObject)

	collection := dbMongoInstance.Collection(collection_name)
	if _, err := mongocontroller.InsertPiece(collection, ctx, pieceObject); err != nil {
		log.Fatal(err)
	}
}
