// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package generated

import (
	"context"
)

const getPalleteNotDone = `-- name: GetPalleteNotDone :one
SELECT
	id, data_print_id, parent_id, serial_level, barcode, serialisasi, batch_no, process_order, scanned, product, nie, sku, counter, berat, md, ed, username, station_name, grup, ipc, sample, packer, upload_line, status, sjp, done, sync
FROM
	data_print
WHERE
	Serial_Level = 4
	AND done = 0
	LIMIT 1
`

func (q *Queries) GetPalleteNotDone(ctx context.Context) (DataPrint, error) {
	row := q.db.QueryRowContext(ctx, getPalleteNotDone)
	var i DataPrint
	err := row.Scan(
		&i.ID,
		&i.DataPrintID,
		&i.ParentID,
		&i.SerialLevel,
		&i.Barcode,
		&i.Serialisasi,
		&i.BatchNo,
		&i.ProcessOrder,
		&i.Scanned,
		&i.Product,
		&i.Nie,
		&i.Sku,
		&i.Counter,
		&i.Berat,
		&i.Md,
		&i.Ed,
		&i.Username,
		&i.StationName,
		&i.Grup,
		&i.Ipc,
		&i.Sample,
		&i.Packer,
		&i.UploadLine,
		&i.Status,
		&i.Sjp,
		&i.Done,
		&i.Sync,
	)
	return i, err
}

const updatePalleteToDone = `-- name: UpdatePalleteToDone :exec
UPDATE data_print
SET done = 1
WHERE
	ID = ?
`

func (q *Queries) UpdatePalleteToDone(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, updatePalleteToDone, id)
	return err
}
