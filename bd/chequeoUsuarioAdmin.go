package bd

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/jonarosero/twitter_demo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoUsuarioAdmin recibe un id de parámetro y chequea si tiene privilegios de rol de admin
func ChequeoUsuarioAdmin(id string) (models.Usuario, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	colRol := db.Collection("rol")

	ID := hex.EncodeToString([]byte(id))
	

	condicion := bson.M{"_id":ID}
	condicionRol := bson.M{"rol":"admin"}

	var resultadoUsuario models.Usuario
	var resultadoRol models.Rol

	err := col.FindOne(ctx, condicion).Decode(&resultadoUsuario)

	errRol:= colRol.FindOne(ctx, condicionRol).Decode(&resultadoRol)

	
	if err != nil {
		return resultadoUsuario, false, "No se encuentra el usuario"
	}
	if errRol != nil {
		return resultadoUsuario, false, "No existe el rol de ADMINISTRADOR"
	}

	if (resultadoUsuario.RolId != resultadoRol.ID.String()){
		return resultadoUsuario, false, "No es un ADMINISTRADOR"
	}

	return resultadoUsuario, true, ID
}