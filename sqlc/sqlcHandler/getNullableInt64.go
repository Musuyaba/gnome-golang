package sqlcHandler

import "database/sql"

func GetNullableInt64(result sql.NullInt64) *int64 {
	if result.Valid {
		return &result.Int64
	}
	return nil // Return 0 for NULL values
}
