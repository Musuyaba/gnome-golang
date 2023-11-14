package sqlcHandler

import (
	schema "github.com/Musuyaba/gnome-golang/mongo"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
)

func GetNullDataPrintStatus(result generated.NullDataPrintStatus) *schema.Status {
	if result.Valid {
		text, err := result.Value()
		if err != nil {
			return nil
		}
		if str, ok := text.(string); ok {
			status := schema.Status(str)
			return &status
		}
	}
	return nil // Return 0 for NULL values
}
