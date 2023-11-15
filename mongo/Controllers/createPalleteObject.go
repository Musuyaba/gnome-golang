package mongocontroller

import (
	mongoschema "github.com/Musuyaba/gnome-golang/mongo"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
	"github.com/Musuyaba/gnome-golang/sqlc/sqlcHandler"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePalleteObject(row generated.GetPiecesNotDoneWithWrapKoliPalleteRow) mongoschema.Pallete {
	palleteObject := mongoschema.Pallete{
		PalleteId:     sqlcHandler.ConvertInt64Int32(sqlcHandler.GetNullableInt64(row.PalleteID)),
		Data_Print_Id: sqlcHandler.GetNullableInt32(row.PalleteDataPrintID),
		Parent_ID:     sqlcHandler.GetNullableInt32(row.PalleteParentID),
		Serial_Level:  sqlcHandler.GetNullableInt32(row.PalleteSerialLevel),
		Barcode:       sqlcHandler.GetNullableString(row.PalleteBarcode),
		Serialisasi:   sqlcHandler.GetNullableString(row.PalleteSerialisasi),
		Batch_No:      sqlcHandler.GetNullableString(row.PalleteBatchNo),
		Process_Order: sqlcHandler.GetNullableString(row.PalleteProcessOrder),
		Scanned:       row.PalleteScanned.Time,
		Product:       sqlcHandler.GetNullableString(row.PalleteProduct),
		Nie:           sqlcHandler.GetNullableString(row.PalleteNie),
		Sku:           sqlcHandler.GetNullableString(row.PalleteSku),
		Counter:       sqlcHandler.GetNullableInt32(row.PalleteCounter),
		Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(row.PalleteBerat)),
		Md:            row.PalleteMd.Time,
		Ed:            row.PalleteEd.Time,
		Username:      sqlcHandler.GetNullableString(row.PalleteUsername),
		Station_name:  sqlcHandler.GetNullableString(row.PalleteStationName),
		Grup:          sqlcHandler.GetNullableString(row.PalleteGrup),
		Ipc:           sqlcHandler.GetNullableString(row.PalleteIpc),
		Sample:        row.PalleteSample.Time,
		Packer:        sqlcHandler.GetNullableString(row.PalletePacker),
		Upload_line:   sqlcHandler.GetNullableTime(row.PalleteUploadLine),
		Status:        sqlcHandler.GetNullDataPrintStatus(row.PalleteStatus),
		Sjp:           sqlcHandler.GetNullableString(row.PalleteSjp),
		Done:          bson.RawValue{Type: bson.TypeNull},
		Synced:        bson.RawValue{Type: bson.TypeNull},
	}
	return palleteObject
}
