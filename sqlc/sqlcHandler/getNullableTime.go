package sqlcHandler

import (
	"database/sql"
	"time"
)

func GetNullableTime(result sql.NullTime) *time.Time {
	if result.Valid {
		return &result.Time
	}
	return nil // Return 0 for NULL values
}
