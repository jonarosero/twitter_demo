package bd

import (
	"context"
	"time"

	"github.com/jonarosero/twitter_demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ChequeoUsuarioAdmin recibe un id de parámetro y chequea si tiene privilegios de rol de admin
func ChequeoUsuarioAdmin(id string) (models.Usuario, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	colRol := db.Collection("rol")

	var resultadoUsuario models.Usuario
	var resultadoRol models.Rol
	
	objID, err := primitive.ObjectIDFromHex(id)

	if err!= nil {
		return resultadoUsuario, false, "No se pudo transformar el id"
	}

	condicion := bson.M{"_id":objID}
	condicionRol := bson.M{"rol":"admin"}

	

	errUsuario := col.FindOne(ctx, condicion).Decode(&resultadoUsuario)

	errRol:= colRol.FindOne(ctx, condicionRol).Decode(&resultadoRol)

	
	if errUsuario != nil {
		return resultadoUsuario, false, "No se encuentra el usuario"
	}
	if errRol != nil {
		return resultadoUsuario, false, "No existe el rol de ADMINISTRADOR"
	}

	idUsuarioRol, error := primitive.ObjectIDFromHex(resultadoUsuario.RolId)

	if error != nil {
		return resultadoUsuario, false, "Error al encontrar usuario"
	}
	if (idUsuarioRol != resultadoRol.ID){
		return resultadoUsuario, false, "No es un ADMINISTRADOR"
	}

	return resultadoUsuario, true, id
}