package bd

import (
	"context"
	"time"

	"github.com/jonarosero/twitter_demo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroRol(r models.Rol) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("rol")

	registro := models.Rol{
		ID:  primitive.NewObjectID(),
		Rol: r.Rol,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
