package bd

import (
	"context"
	"time"

	"github.com/jonarosero/twitter_demo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//BuscoPerfil busca un perfil en la BD
func BuscoRol(nombre string) (models.Rol, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("rol")

	condicion := bson.M{"rol":nombre}

	var resultado models.Rol

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	
	if err != nil {
		return resultado, err
	}
	return resultado, err
}