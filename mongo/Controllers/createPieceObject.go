package mongocontroller

import (
	schema "github.com/Musuyaba/gnome-golang/mongo"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
	"github.com/Musuyaba/gnome-golang/sqlc/sqlcHandler"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePieceObject(row generated.GetPiecesNotDoneWithWrapKoliPalleteRow, wrapObject schema.Inner) schema.Piece {
	pieceObject := schema.Piece{
		PieceId:       sqlcHandler.ConvertInt64Int32(&row.PieceID),
		Data_Print_Id: sqlcHandler.GetNullableInt32(row.PieceDataPrintID),
		Parent_ID:     sqlcHandler.GetNullableInt32(row.PieceParentID),
		Serial_Level:  sqlcHandler.GetNullableInt32(row.PieceSerialLevel),
		Barcode:       sqlcHandler.GetNullableString(row.PieceBarcode),
		Serialisasi:   &row.PieceSerialisasi,
		Batch_No:      sqlcHandler.GetNullableString(row.PieceBatchNo),
		Process_Order: &row.PieceProcessOrder,
		Scanned:       row.PieceScanned.Time,
		Product:       sqlcHandler.GetNullableString(row.PieceProduct),
		Nie:           sqlcHandler.GetNullableString(row.PieceNie),
		Sku:           sqlcHandler.GetNullableString(row.PieceSku),
		Counter:       sqlcHandler.GetNullableInt32(row.PieceCounter),
		Berat:         sqlcHandler.ConvertStringFloat32(sqlcHandler.GetNullableString(row.PieceBerat)),
		Md:            row.PieceMd.Time,
		Ed:            row.PieceEd.Time,
		Username:      sqlcHandler.GetNullableString(row.PieceUsername),
		Station_name:  sqlcHandler.GetNullableString(row.PieceStationName),
		Grup:          sqlcHandler.GetNullableString(row.PieceGrup),
		Ipc:           sqlcHandler.GetNullableString(row.PieceIpc),
		Sample:        sqlcHandler.GetNullableTime(row.PieceSample),
		Packer:        sqlcHandler.GetNullableString(row.PiecePacker),
		Upload_line:   row.PieceUploadLine,
		Status:        schema.Status(row.PieceStatus),
		Sjp:           sqlcHandler.GetNullableString(row.PieceSjp),
		Done:          bson.RawValue{Type: bson.TypeNull},
		Synced:        bson.RawValue{Type: bson.TypeNull},
		Wrap:          wrapObject,
	}
	return pieceObject
}
