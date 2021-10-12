package bd

import (
	"context"
	"time"

	"github.com/jonarosero/twitter_demo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteRol recibe un rol de parámetro y chequea si ya existe en la BD
func ChequeoYaExisteRol (rol string) (models.Rol, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("rol")

	condicion := bson.M{"rol":rol}

	var resultado models.Rol

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}