-- name: GetPalleteNotDone :one
SELECT
	*
FROM
	data_print
WHERE
	Serial_Level = 4
	AND done = 0
	LIMIT 1;

-- name: UpdatePalleteToDone :exec
UPDATE data_print
SET done = 1
WHERE
	ID = ?;