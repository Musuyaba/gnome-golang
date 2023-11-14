package sqlcControllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/Musuyaba/gnome-golang/pkg/initializers"
	"github.com/Musuyaba/gnome-golang/sqlc/generated"
)

func FirstAndupdateDataPrint(ctx context.Context, MySqlcInstance *sql.DB) (generated.GetPiecesNotDoneWithWrapKoliPalleteRow, error) {
	sqlcQuery := generated.New(initializers.MySqlcInstance)
	rowDataPrint, err := sqlcQuery.GetPiecesNotDoneWithWrapKoliPallete(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = sqlcQuery.UpdatePieceToDone(ctx, rowDataPrint.PieceID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Piece ID: " + strconv.Itoa(int(rowDataPrint.PieceID)) + "\n")
	return rowDataPrint, err
}
