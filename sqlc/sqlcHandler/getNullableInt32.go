package sqlcHandler

import "database/sql"

func GetNullableInt32(result sql.NullInt32) *int32 {
	if result.Valid {
		return &result.Int32
	}
	return nil // Return 0 for NULL values
}
