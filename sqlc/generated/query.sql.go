// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const getPiecesNotDoneWithWrapKoliPallete = `-- name: GetPiecesNotDoneWithWrapKoliPallete :one
SELECT
	piece.ID AS piece_ID,
	piece.Data_Print_ID AS piece_Data_Print_ID,
	piece.Parent_ID AS piece_Parent_ID,
	piece.Serial_Level AS piece_Serial_Level,
	piece.Barcode AS piece_Barcode,
	piece.Serialisasi AS piece_Serialisasi,
	piece.Batch_No AS piece_Batch_No,
	piece.Process_Order AS piece_Process_Order,	
	piece.Scanned AS piece_Scanned,
	piece.product AS piece_product,
	piece.nie AS piece_nie,
	piece.sku AS piece_sku,
	piece.Counter AS piece_Counter,
	piece.Berat AS piece_Berat,
	piece.md AS piece_md,
	piece.ed AS piece_ed,
	piece.Username AS piece_Username,
	piece.station_name AS piece_station_name,
	piece.grup AS piece_grup,
	piece.ipc AS piece_ipc,
	piece.Sample AS piece_Sample,
	piece.packer AS piece_packer,
	piece.upload_line AS piece_upload_line,
	piece.status AS piece_status,
	piece.sjp AS piece_sjp,
	piece.done AS piece_done,
	piece.synced AS piece_synced,
	wrap.ID AS wrap_ID,
	wrap.Data_Print_ID AS wrap_Data_Print_ID,
	wrap.Parent_ID AS wrap_Parent_ID,
	wrap.Serial_Level AS wrap_Serial_Level,
	wrap.Barcode AS wrap_Barcode,
	wrap.Serialisasi AS wrap_Serialisasi,
	wrap.Batch_No AS wrap_Batch_No,
	wrap.Process_Order AS wrap_Process_Order,	
	wrap.Scanned AS wrap_Scanned,
	wrap.product AS wrap_product,
	wrap.nie AS wrap_nie,
	wrap.sku AS wrap_sku,
	wrap.Counter AS wrap_Counter,
	wrap.Berat AS wrap_Berat,
	wrap.md AS wrap_md,
	wrap.ed AS wrap_ed,
	wrap.Username AS wrap_Username,
	wrap.station_name AS wrap_station_name,
	wrap.grup AS wrap_grup,
	wrap.ipc AS wrap_ipc,
	wrap.Sample AS wrap_Sample,
	wrap.packer AS wrap_packer,
	wrap.upload_line AS wrap_upload_line,
	wrap.status AS wrap_status,
	wrap.sjp AS wrap_sjp,
	wrap.done AS wrap_done,
	wrap.synced AS wrap_synced,
	koli.ID AS koli_ID,
	koli.Data_Print_ID AS koli_Data_Print_ID,
	koli.Parent_ID AS koli_Parent_ID,
	koli.Serial_Level AS koli_Serial_Level,
	koli.Barcode AS koli_Barcode,
	koli.Serialisasi AS koli_Serialisasi,
	koli.Batch_No AS koli_Batch_No,
	koli.Process_Order AS koli_Process_Order,	
	koli.Scanned AS koli_Scanned,
	koli.product AS koli_product,
	koli.nie AS koli_nie,
	koli.sku AS koli_sku,
	koli.Counter AS koli_Counter,
	koli.Berat AS koli_Berat,
	koli.md AS koli_md,
	koli.ed AS koli_ed,
	koli.Username AS koli_Username,
	koli.station_name AS koli_station_name,
	koli.grup AS koli_grup,
	koli.ipc AS koli_ipc,
	koli.Sample AS koli_Sample,
	koli.packer AS koli_packer,
	koli.upload_line AS koli_upload_line,
	koli.status AS koli_status,
	koli.sjp AS koli_sjp,
	koli.done AS koli_done,
	koli.synced AS koli_synced,
	pallete.ID AS pallete_ID,
	pallete.Data_Print_ID AS pallete_Data_Print_ID,
	pallete.Parent_ID AS pallete_Parent_ID,
	pallete.Serial_Level AS pallete_Serial_Level,
	pallete.Barcode AS pallete_Barcode,
	pallete.Serialisasi AS pallete_Serialisasi,
	pallete.Batch_No AS pallete_Batch_No,
	pallete.Process_Order AS pallete_Process_Order,	
	pallete.Scanned AS pallete_Scanned,
	pallete.product AS pallete_product,
	pallete.nie AS pallete_nie,
	pallete.sku AS pallete_sku,
	pallete.Counter AS pallete_Counter,
	pallete.Berat AS pallete_Berat,
	pallete.md AS pallete_md,
	pallete.ed AS pallete_ed,
	pallete.Username AS pallete_Username,
	pallete.station_name AS pallete_station_name,
	pallete.grup AS pallete_grup,
	pallete.ipc AS pallete_ipc,
	pallete.Sample AS pallete_Sample,
	pallete.packer AS pallete_packer,
	pallete.upload_line AS pallete_upload_line,
	pallete.status AS pallete_status,
	pallete.sjp AS pallete_sjp,
	pallete.done AS pallete_done,
	pallete.synced AS pallete_synced
FROM
	data_print piece
LEFT JOIN data_print wrap on wrap.Data_Print_ID = piece.Parent_ID
LEFT JOIN data_print koli on koli.Data_Print_ID = wrap.Parent_ID
LEFT JOIN data_print pallete on pallete.Data_Print_ID = koli.Parent_ID
WHERE
	piece.Serial_Level = 1 
	AND piece.done is NULL
	LIMIT 1
`

type GetPiecesNotDoneWithWrapKoliPalleteRow struct {
	PieceID             int64
	PieceDataPrintID    sql.NullInt32
	PieceParentID       sql.NullInt32
	PieceSerialLevel    sql.NullInt32
	PieceBarcode        sql.NullString
	PieceSerialisasi    string
	PieceBatchNo        sql.NullString
	PieceProcessOrder   string
	PieceScanned        sql.NullTime
	PieceProduct        sql.NullString
	PieceNie            sql.NullString
	PieceSku            sql.NullString
	PieceCounter        sql.NullInt32
	PieceBerat          sql.NullString
	PieceMd             sql.NullTime
	PieceEd             sql.NullTime
	PieceUsername       sql.NullString
	PieceStationName    sql.NullString
	PieceGrup           sql.NullString
	PieceIpc            sql.NullString
	PieceSample         sql.NullTime
	PiecePacker         sql.NullString
	PieceUploadLine     time.Time
	PieceStatus         DataPrintStatus
	PieceSjp            sql.NullString
	PieceDone           sql.NullInt32
	PieceSynced         sql.NullInt32
	WrapID              sql.NullInt64
	WrapDataPrintID     sql.NullInt32
	WrapParentID        sql.NullInt32
	WrapSerialLevel     sql.NullInt32
	WrapBarcode         sql.NullString
	WrapSerialisasi     sql.NullString
	WrapBatchNo         sql.NullString
	WrapProcessOrder    sql.NullString
	WrapScanned         sql.NullTime
	WrapProduct         sql.NullString
	WrapNie             sql.NullString
	WrapSku             sql.NullString
	WrapCounter         sql.NullInt32
	WrapBerat           sql.NullString
	WrapMd              sql.NullTime
	WrapEd              sql.NullTime
	WrapUsername        sql.NullString
	WrapStationName     sql.NullString
	WrapGrup            sql.NullString
	WrapIpc             sql.NullString
	WrapSample          sql.NullTime
	WrapPacker          sql.NullString
	WrapUploadLine      sql.NullTime
	WrapStatus          NullDataPrintStatus
	WrapSjp             sql.NullString
	WrapDone            sql.NullInt32
	WrapSynced          sql.NullInt32
	KoliID              sql.NullInt64
	KoliDataPrintID     sql.NullInt32
	KoliParentID        sql.NullInt32
	KoliSerialLevel     sql.NullInt32
	KoliBarcode         sql.NullString
	KoliSerialisasi     sql.NullString
	KoliBatchNo         sql.NullString
	KoliProcessOrder    sql.NullString
	KoliScanned         sql.NullTime
	KoliProduct         sql.NullString
	KoliNie             sql.NullString
	KoliSku             sql.NullString
	KoliCounter         sql.NullInt32
	KoliBerat           sql.NullString
	KoliMd              sql.NullTime
	KoliEd              sql.NullTime
	KoliUsername        sql.NullString
	KoliStationName     sql.NullString
	KoliGrup            sql.NullString
	KoliIpc             sql.NullString
	KoliSample          sql.NullTime
	KoliPacker          sql.NullString
	KoliUploadLine      sql.NullTime
	KoliStatus          NullDataPrintStatus
	KoliSjp             sql.NullString
	KoliDone            sql.NullInt32
	KoliSynced          sql.NullInt32
	PalleteID           sql.NullInt64
	PalleteDataPrintID  sql.NullInt32
	PalleteParentID     sql.NullInt32
	PalleteSerialLevel  sql.NullInt32
	PalleteBarcode      sql.NullString
	PalleteSerialisasi  sql.NullString
	PalleteBatchNo      sql.NullString
	PalleteProcessOrder sql.NullString
	PalleteScanned      sql.NullTime
	PalleteProduct      sql.NullString
	PalleteNie          sql.NullString
	PalleteSku          sql.NullString
	PalleteCounter      sql.NullInt32
	PalleteBerat        sql.NullString
	PalleteMd           sql.NullTime
	PalleteEd           sql.NullTime
	PalleteUsername     sql.NullString
	PalleteStationName  sql.NullString
	PalleteGrup         sql.NullString
	PalleteIpc          sql.NullString
	PalleteSample       sql.NullTime
	PalletePacker       sql.NullString
	PalleteUploadLine   sql.NullTime
	PalleteStatus       NullDataPrintStatus
	PalleteSjp          sql.NullString
	PalleteDone         sql.NullInt32
	PalleteSynced       sql.NullInt32
}

func (q *Queries) GetPiecesNotDoneWithWrapKoliPallete(ctx context.Context) (GetPiecesNotDoneWithWrapKoliPalleteRow, error) {
	row := q.db.QueryRowContext(ctx, getPiecesNotDoneWithWrapKoliPallete)
	var i GetPiecesNotDoneWithWrapKoliPalleteRow
	err := row.Scan(
		&i.PieceID,
		&i.PieceDataPrintID,
		&i.PieceParentID,
		&i.PieceSerialLevel,
		&i.PieceBarcode,
		&i.PieceSerialisasi,
		&i.PieceBatchNo,
		&i.PieceProcessOrder,
		&i.PieceScanned,
		&i.PieceProduct,
		&i.PieceNie,
		&i.PieceSku,
		&i.PieceCounter,
		&i.PieceBerat,
		&i.PieceMd,
		&i.PieceEd,
		&i.PieceUsername,
		&i.PieceStationName,
		&i.PieceGrup,
		&i.PieceIpc,
		&i.PieceSample,
		&i.PiecePacker,
		&i.PieceUploadLine,
		&i.PieceStatus,
		&i.PieceSjp,
		&i.PieceDone,
		&i.PieceSynced,
		&i.WrapID,
		&i.WrapDataPrintID,
		&i.WrapParentID,
		&i.WrapSerialLevel,
		&i.WrapBarcode,
		&i.WrapSerialisasi,
		&i.WrapBatchNo,
		&i.WrapProcessOrder,
		&i.WrapScanned,
		&i.WrapProduct,
		&i.WrapNie,
		&i.WrapSku,
		&i.WrapCounter,
		&i.WrapBerat,
		&i.WrapMd,
		&i.WrapEd,
		&i.WrapUsername,
		&i.WrapStationName,
		&i.WrapGrup,
		&i.WrapIpc,
		&i.WrapSample,
		&i.WrapPacker,
		&i.WrapUploadLine,
		&i.WrapStatus,
		&i.WrapSjp,
		&i.WrapDone,
		&i.WrapSynced,
		&i.KoliID,
		&i.KoliDataPrintID,
		&i.KoliParentID,
		&i.KoliSerialLevel,
		&i.KoliBarcode,
		&i.KoliSerialisasi,
		&i.KoliBatchNo,
		&i.KoliProcessOrder,
		&i.KoliScanned,
		&i.KoliProduct,
		&i.KoliNie,
		&i.KoliSku,
		&i.KoliCounter,
		&i.KoliBerat,
		&i.KoliMd,
		&i.KoliEd,
		&i.KoliUsername,
		&i.KoliStationName,
		&i.KoliGrup,
		&i.KoliIpc,
		&i.KoliSample,
		&i.KoliPacker,
		&i.KoliUploadLine,
		&i.KoliStatus,
		&i.KoliSjp,
		&i.KoliDone,
		&i.KoliSynced,
		&i.PalleteID,
		&i.PalleteDataPrintID,
		&i.PalleteParentID,
		&i.PalleteSerialLevel,
		&i.PalleteBarcode,
		&i.PalleteSerialisasi,
		&i.PalleteBatchNo,
		&i.PalleteProcessOrder,
		&i.PalleteScanned,
		&i.PalleteProduct,
		&i.PalleteNie,
		&i.PalleteSku,
		&i.PalleteCounter,
		&i.PalleteBerat,
		&i.PalleteMd,
		&i.PalleteEd,
		&i.PalleteUsername,
		&i.PalleteStationName,
		&i.PalleteGrup,
		&i.PalleteIpc,
		&i.PalleteSample,
		&i.PalletePacker,
		&i.PalleteUploadLine,
		&i.PalleteStatus,
		&i.PalleteSjp,
		&i.PalleteDone,
		&i.PalleteSynced,
	)
	return i, err
}

const updatePieceToDone = `-- name: UpdatePieceToDone :exec
UPDATE data_print 
SET done = 1 
WHERE
	ID = ?
`

func (q *Queries) UpdatePieceToDone(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, updatePieceToDone, id)
	return err
}