-- name: GetPalleteNotDone :one
SELECT
	*
FROM
	data_print
WHERE
	Serial_Level = 4
	AND done is NULL
	LIMIT 1;

-- name: UpdatePalleteToDone :exec
UPDATE data_print
SET done = 0
WHERE
	ID = ?;
