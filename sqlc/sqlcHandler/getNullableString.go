package sqlcHandler

import "database/sql"

func GetNullableString(result sql.NullString) *string {
	if result.Valid {
		return &result.String
	}
	return nil // Return 0 for NULL values
}
