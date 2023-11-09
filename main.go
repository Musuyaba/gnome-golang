package main

import (
	"context"
	"fmt"
	"log"
	"time"

	schema "github.com/Musuyaba/gnome-golang/mongo"
	"github.com/Musuyaba/gnome-golang/pkg/initializers"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
	"github.com/Musuyaba/gnome-golang/sqlc/sqlcHandler"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
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

	initializers.ConnectDbGormPostgreSQL(&config)
	initializers.ConnectDbSqlcMySQL(&config)
	initializers.ConnectDbMongo(&config, ctx)
	defer initializers.MongoInstance.Disconnect(ctx)
	// server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	sqlc_query := generated.New(initializers.MySqlcInstance)
	startTime := time.Now()
	rowPallete, err := sqlc_query.GetPalleteNotDone(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// err = sqlc_query.UpdatePalleteToDone(ctx, rowPallete.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db_mongo := initializers.MongoInstance.Database(config.DATABASE_MONGODB_DATA_NODE_A_NAME)
	collection := db_mongo.Collection("my-collection")

	pallete := schema.Pallete{
		Data_Print_Id: sqlcHandler.GetNullableInt32(rowPallete.DataPrintID),
		Parent_ID:     sqlcHandler.GetNullableInt32(rowPallete.ParentID),
		Serial_Level:  sqlcHandler.GetNullableInt32(rowPallete.SerialLevel),
		Barcode:       sqlcHandler.GetNullableString(rowPallete.Barcode),
		Serialisasi:   &rowPallete.Serialisasi,
		Batch_No:      sqlcHandler.GetNullableString(rowPallete.BatchNo),
		Process_Order: &rowPallete.ProcessOrder,
		Scanned:       rowPallete.Scanned.Time,
		Product:       sqlcHandler.GetNullableString(rowPallete.Product),
		Nie:           sqlcHandler.GetNullableString(rowPallete.Nie),
		Sku:           sqlcHandler.GetNullableString(rowPallete.Sku),
		Counter:       sqlcHandler.GetNullableInt32(rowPallete.Counter),
		Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(rowPallete.Berat)),
		Md:            rowPallete.Md.Time,
		Ed:            rowPallete.Ed.Time,
		Username:      sqlcHandler.GetNullableString(rowPallete.Username),
		Station_name:  sqlcHandler.GetNullableString(rowPallete.StationName),
		Grup:          sqlcHandler.GetNullableString(rowPallete.Grup),
		Ipc:           sqlcHandler.GetNullableString(rowPallete.Ipc),
		Sample:        rowPallete.Sample.Time,
		Packer:        sqlcHandler.GetNullableString(rowPallete.Packer),
		Upload_line:   rowPallete.UploadLine,
		Status:        schema.Status(rowPallete.Status),
		Sjp:           sqlcHandler.GetNullableString(rowPallete.Sjp),
		Done:          bson.RawValue{Type: bson.TypeNull},
		Synced:        bson.RawValue{Type: bson.TypeNull},
	}

	result, err := collection.InsertOne(ctx, &pallete)
	if err != nil {
		log.Fatal(err)
	}

	// Print the inserted document ID.
	fmt.Println(result.InsertedID)

	executionTime := time.Since(startTime)
	fmt.Println("Execution time:", executionTime)

	// // routers.ApiRoutes(server)
	// // routers.PublicRoutes(server)

	// // log.Fatal(server.RunTLS(config.HOSTNAME+":"+config.PORT, "./certs/generated/server.crt", "./certs/generated/server.key"))
}
