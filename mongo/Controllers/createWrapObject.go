package mongocontroller

import (
	mongoschema "github.com/Musuyaba/gnome-golang/mongo"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
	"github.com/Musuyaba/gnome-golang/sqlc/sqlcHandler"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateWrapObject(row generated.GetPiecesNotDoneWithWrapKoliPalleteRow, koliObject mongoschema.Box) mongoschema.Inner {
	wrapObject := mongoschema.Inner{
		InnerId:       sqlcHandler.ConvertInt64Int32(sqlcHandler.GetNullableInt64(row.WrapID)),
		Data_Print_Id: sqlcHandler.GetNullableInt32(row.WrapDataPrintID),
		Parent_ID:     sqlcHandler.GetNullableInt32(row.WrapParentID),
		Serial_Level:  sqlcHandler.GetNullableInt32(row.WrapSerialLevel),
		Barcode:       sqlcHandler.GetNullableString(row.WrapBarcode),
		Serialisasi:   sqlcHandler.GetNullableString(row.WrapSerialisasi),
		Batch_No:      sqlcHandler.GetNullableString(row.WrapBatchNo),
		Process_Order: sqlcHandler.GetNullableString(row.WrapProcessOrder),
		Scanned:       row.WrapScanned.Time,
		Product:       sqlcHandler.GetNullableString(row.WrapProduct),
		Nie:           sqlcHandler.GetNullableString(row.WrapNie),
		Sku:           sqlcHandler.GetNullableString(row.WrapSku),
		Counter:       sqlcHandler.GetNullableInt32(row.WrapCounter),
		Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(row.WrapBerat)),
		Md:            row.WrapMd.Time,
		Ed:            row.WrapEd.Time,
		Username:      sqlcHandler.GetNullableString(row.WrapUsername),
		Station_name:  sqlcHandler.GetNullableString(row.WrapStationName),
		Grup:          sqlcHandler.GetNullableString(row.WrapGrup),
		Ipc:           sqlcHandler.GetNullableString(row.WrapIpc),
		Sample:        row.WrapSample.Time,
		Packer:        sqlcHandler.GetNullableString(row.WrapPacker),
		Upload_line:   sqlcHandler.GetNullableTime(row.WrapUploadLine),
		Status:        sqlcHandler.GetNullDataPrintStatus(row.WrapStatus),
		Sjp:           sqlcHandler.GetNullableString(row.WrapSjp),
		Done:          bson.RawValue{Type: bson.TypeNull},
		Synced:        bson.RawValue{Type: bson.TypeNull},
		Koli:          koliObject,
	}
	return wrapObject
}
