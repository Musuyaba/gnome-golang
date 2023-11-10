package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

	for i := 0; i < 1; i++ {

		sqlc_query := generated.New(initializers.MySqlcInstance)
		startTime := time.Now()
		rowPallete, err := sqlc_query.GetPiecesNotDoneWithWrapKoliPallete(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Data Print ID: " + strconv.Itoa(int(rowPallete.PieceID)) + "\n")

		err = sqlc_query.UpdatePieceToDone(ctx, rowPallete.PieceID)
		if err != nil {
			log.Fatal(err)
		}

		db_mongo := initializers.MongoInstance.Database(config.DATABASE_MONGODB_DATA_NODE_A_NAME)
		collection := db_mongo.Collection("qr-flatten")

		palleteObject := schema.Pallete{
			PalleteId:     sqlcHandler.ConvertInt64Int32(sqlcHandler.GetNullableInt64(rowPallete.PalleteID)),
			Data_Print_Id: sqlcHandler.GetNullableInt32(rowPallete.PalleteDataPrintID),
			Parent_ID:     sqlcHandler.GetNullableInt32(rowPallete.PalleteParentID),
			Serial_Level:  sqlcHandler.GetNullableInt32(rowPallete.PalleteSerialLevel),
			Barcode:       sqlcHandler.GetNullableString(rowPallete.PalleteBarcode),
			Serialisasi:   sqlcHandler.GetNullableString(rowPallete.PalleteSerialisasi),
			Batch_No:      sqlcHandler.GetNullableString(rowPallete.PalleteBatchNo),
			Process_Order: sqlcHandler.GetNullableString(rowPallete.PalleteProcessOrder),
			Scanned:       rowPallete.PalleteScanned.Time,
			Product:       sqlcHandler.GetNullableString(rowPallete.PalleteProduct),
			Nie:           sqlcHandler.GetNullableString(rowPallete.PalleteNie),
			Sku:           sqlcHandler.GetNullableString(rowPallete.PalleteSku),
			Counter:       sqlcHandler.GetNullableInt32(rowPallete.PalleteCounter),
			Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(rowPallete.PalleteBerat)),
			Md:            rowPallete.PalleteMd.Time,
			Ed:            rowPallete.PalleteEd.Time,
			Username:      sqlcHandler.GetNullableString(rowPallete.PalleteUsername),
			Station_name:  sqlcHandler.GetNullableString(rowPallete.PalleteStationName),
			Grup:          sqlcHandler.GetNullableString(rowPallete.PalleteGrup),
			Ipc:           sqlcHandler.GetNullableString(rowPallete.PalleteIpc),
			Sample:        rowPallete.PalleteSample.Time,
			Packer:        sqlcHandler.GetNullableString(rowPallete.PalletePacker),
			Upload_line:   sqlcHandler.GetNullableTime(rowPallete.PalleteUploadLine),
			Status:        sqlcHandler.GetNullDataPrintStatus(rowPallete.PalleteStatus),
			Sjp:           sqlcHandler.GetNullableString(rowPallete.PalleteSjp),
			Done:          bson.RawValue{Type: bson.TypeNull},
			Synced:        bson.RawValue{Type: bson.TypeNull},
		}

		koliObject := schema.Box{
			BoxId:         sqlcHandler.ConvertInt64Int32(sqlcHandler.GetNullableInt64(rowPallete.KoliID)),
			Data_Print_Id: sqlcHandler.GetNullableInt32(rowPallete.KoliDataPrintID),
			Parent_ID:     sqlcHandler.GetNullableInt32(rowPallete.KoliParentID),
			Serial_Level:  sqlcHandler.GetNullableInt32(rowPallete.KoliSerialLevel),
			Barcode:       sqlcHandler.GetNullableString(rowPallete.KoliBarcode),
			Serialisasi:   sqlcHandler.GetNullableString(rowPallete.KoliSerialisasi),
			Batch_No:      sqlcHandler.GetNullableString(rowPallete.KoliBatchNo),
			Process_Order: sqlcHandler.GetNullableString(rowPallete.KoliProcessOrder),
			Scanned:       rowPallete.KoliScanned.Time,
			Product:       sqlcHandler.GetNullableString(rowPallete.KoliProduct),
			Nie:           sqlcHandler.GetNullableString(rowPallete.KoliNie),
			Sku:           sqlcHandler.GetNullableString(rowPallete.KoliSku),
			Counter:       sqlcHandler.GetNullableInt32(rowPallete.KoliCounter),
			Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(rowPallete.KoliBerat)),
			Md:            rowPallete.KoliMd.Time,
			Ed:            rowPallete.KoliEd.Time,
			Username:      sqlcHandler.GetNullableString(rowPallete.KoliUsername),
			Station_name:  sqlcHandler.GetNullableString(rowPallete.KoliStationName),
			Grup:          sqlcHandler.GetNullableString(rowPallete.KoliGrup),
			Ipc:           sqlcHandler.GetNullableString(rowPallete.KoliIpc),
			Sample:        rowPallete.KoliSample.Time,
			Packer:        sqlcHandler.GetNullableString(rowPallete.KoliPacker),
			Upload_line:   sqlcHandler.GetNullableTime(rowPallete.KoliUploadLine),
			Status:        sqlcHandler.GetNullDataPrintStatus(rowPallete.KoliStatus),
			Sjp:           sqlcHandler.GetNullableString(rowPallete.KoliSjp),
			Done:          bson.RawValue{Type: bson.TypeNull},
			Synced:        bson.RawValue{Type: bson.TypeNull},
			Pallete:       palleteObject,
		}

		wrapObject := schema.Inner{
			InnerId:       sqlcHandler.ConvertInt64Int32(sqlcHandler.GetNullableInt64(rowPallete.WrapID)),
			Data_Print_Id: sqlcHandler.GetNullableInt32(rowPallete.WrapDataPrintID),
			Parent_ID:     sqlcHandler.GetNullableInt32(rowPallete.WrapParentID),
			Serial_Level:  sqlcHandler.GetNullableInt32(rowPallete.WrapSerialLevel),
			Barcode:       sqlcHandler.GetNullableString(rowPallete.WrapBarcode),
			Serialisasi:   sqlcHandler.GetNullableString(rowPallete.WrapSerialisasi),
			Batch_No:      sqlcHandler.GetNullableString(rowPallete.WrapBatchNo),
			Process_Order: sqlcHandler.GetNullableString(rowPallete.WrapProcessOrder),
			Scanned:       rowPallete.WrapScanned.Time,
			Product:       sqlcHandler.GetNullableString(rowPallete.WrapProduct),
			Nie:           sqlcHandler.GetNullableString(rowPallete.WrapNie),
			Sku:           sqlcHandler.GetNullableString(rowPallete.WrapSku),
			Counter:       sqlcHandler.GetNullableInt32(rowPallete.WrapCounter),
			Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(rowPallete.WrapBerat)),
			Md:            rowPallete.WrapMd.Time,
			Ed:            rowPallete.WrapEd.Time,
			Username:      sqlcHandler.GetNullableString(rowPallete.WrapUsername),
			Station_name:  sqlcHandler.GetNullableString(rowPallete.WrapStationName),
			Grup:          sqlcHandler.GetNullableString(rowPallete.WrapGrup),
			Ipc:           sqlcHandler.GetNullableString(rowPallete.WrapIpc),
			Sample:        rowPallete.WrapSample.Time,
			Packer:        sqlcHandler.GetNullableString(rowPallete.WrapPacker),
			Upload_line:   sqlcHandler.GetNullableTime(rowPallete.WrapUploadLine),
			Status:        sqlcHandler.GetNullDataPrintStatus(rowPallete.WrapStatus),
			Sjp:           sqlcHandler.GetNullableString(rowPallete.WrapSjp),
			Done:          bson.RawValue{Type: bson.TypeNull},
			Synced:        bson.RawValue{Type: bson.TypeNull},
			Koli:          koliObject,
		}

		pieceObject := schema.Piece{
			PieceId:       sqlcHandler.ConvertInt64Int32(&rowPallete.PieceID),
			Data_Print_Id: sqlcHandler.GetNullableInt32(rowPallete.PieceDataPrintID),
			Parent_ID:     sqlcHandler.GetNullableInt32(rowPallete.PieceParentID),
			Serial_Level:  sqlcHandler.GetNullableInt32(rowPallete.PieceSerialLevel),
			Barcode:       sqlcHandler.GetNullableString(rowPallete.PieceBarcode),
			Serialisasi:   &rowPallete.PieceSerialisasi,
			Batch_No:      sqlcHandler.GetNullableString(rowPallete.PieceBatchNo),
			Process_Order: &rowPallete.PieceProcessOrder,
			Scanned:       rowPallete.PieceScanned.Time,
			Product:       sqlcHandler.GetNullableString(rowPallete.PieceProduct),
			Nie:           sqlcHandler.GetNullableString(rowPallete.PieceNie),
			Sku:           sqlcHandler.GetNullableString(rowPallete.PieceSku),
			Counter:       sqlcHandler.GetNullableInt32(rowPallete.PieceCounter),
			Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(rowPallete.PieceBerat)),
			Md:            rowPallete.PieceMd.Time,
			Ed:            rowPallete.PieceEd.Time,
			Username:      sqlcHandler.GetNullableString(rowPallete.PieceUsername),
			Station_name:  sqlcHandler.GetNullableString(rowPallete.PieceStationName),
			Grup:          sqlcHandler.GetNullableString(rowPallete.PieceGrup),
			Ipc:           sqlcHandler.GetNullableString(rowPallete.PieceIpc),
			Sample:        sqlcHandler.GetNullableTime(rowPallete.PieceSample),
			Packer:        sqlcHandler.GetNullableString(rowPallete.PiecePacker),
			Upload_line:   rowPallete.PieceUploadLine,
			Status:        schema.Status(rowPallete.PieceStatus),
			Sjp:           sqlcHandler.GetNullableString(rowPallete.PieceSjp),
			Done:          bson.RawValue{Type: bson.TypeNull},
			Synced:        bson.RawValue{Type: bson.TypeNull},
			Wrap:          wrapObject,
		}

		result, err := collection.InsertOne(ctx, &pieceObject)
		if err != nil {
			log.Fatal(err)
		}

		// Print the inserted document ID.
		fmt.Println(result.InsertedID)

		executionTime := time.Since(startTime)
		fmt.Println("Execution time:", executionTime)
	}
	// defer initializers.MongoInstance.Disconnect(ctx)

	// // routers.ApiRoutes(server)
	// // routers.PublicRoutes(server)

	// // log.Fatal(server.RunTLS(config.HOSTNAME+":"+config.PORT, "./certs/generated/server.crt", "./certs/generated/server.key"))
}
