package mongoController

import (
	schema "github.com/Musuyaba/gnome-golang/mongo"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
	"github.com/Musuyaba/gnome-golang/sqlc/sqlcHandler"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateKoliObject(row generated.GetPiecesNotDoneWithWrapKoliPalleteRow, palleteObject schema.Pallete) schema.Box {
	koliObject := schema.Box{
		BoxId:         sqlcHandler.ConvertInt64Int32(sqlcHandler.GetNullableInt64(row.KoliID)),
		Data_Print_Id: sqlcHandler.GetNullableInt32(row.KoliDataPrintID),
		Parent_ID:     sqlcHandler.GetNullableInt32(row.KoliParentID),
		Serial_Level:  sqlcHandler.GetNullableInt32(row.KoliSerialLevel),
		Barcode:       sqlcHandler.GetNullableString(row.KoliBarcode),
		Serialisasi:   sqlcHandler.GetNullableString(row.KoliSerialisasi),
		Batch_No:      sqlcHandler.GetNullableString(row.KoliBatchNo),
		Process_Order: sqlcHandler.GetNullableString(row.KoliProcessOrder),
		Scanned:       row.KoliScanned.Time,
		Product:       sqlcHandler.GetNullableString(row.KoliProduct),
		Nie:           sqlcHandler.GetNullableString(row.KoliNie),
		Sku:           sqlcHandler.GetNullableString(row.KoliSku),
		Counter:       sqlcHandler.GetNullableInt32(row.KoliCounter),
		Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(row.KoliBerat)),
		Md:            row.KoliMd.Time,
		Ed:            row.KoliEd.Time,
		Username:      sqlcHandler.GetNullableString(row.KoliUsername),
		Station_name:  sqlcHandler.GetNullableString(row.KoliStationName),
		Grup:          sqlcHandler.GetNullableString(row.KoliGrup),
		Ipc:           sqlcHandler.GetNullableString(row.KoliIpc),
		Sample:        row.KoliSample.Time,
		Packer:        sqlcHandler.GetNullableString(row.KoliPacker),
		Upload_line:   sqlcHandler.GetNullableTime(row.KoliUploadLine),
		Status:        sqlcHandler.GetNullDataPrintStatus(row.KoliStatus),
		Sjp:           sqlcHandler.GetNullableString(row.KoliSjp),
		Done:          bson.RawValue{Type: bson.TypeNull},
		Synced:        bson.RawValue{Type: bson.TypeNull},
		Pallete:       palleteObject,
	}
	return koliObject
}
