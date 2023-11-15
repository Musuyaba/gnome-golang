package mongocontroller

import (
	"context"
	"fmt"
	"log"

	mongoschema "github.com/Musuyaba/gnome-golang/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPiece(collection *mongo.Collection, ctx context.Context, pieceObject mongoschema.Piece) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, &pieceObject)
	if err != nil {
		log.Fatal(err)
	}

	// Print the inserted document ID.
	fmt.Println(result.InsertedID)
	return result, err
}
