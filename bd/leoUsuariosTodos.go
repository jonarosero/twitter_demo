package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jonarosero/twitter_demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos (ID string, page int64, search string, tipo string) ([] *models.Usuario, bool ){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	var results [] *models.Usuario 

	findOptions := options.Find()
	findOptions.SetSkip((page-1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"regex":`(?i)`+search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx){
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID=ID
		r.UsuarioRelacionID=s.ID.Hex()

		incluir = false

		encontrado, _ = ConsultoRelacion(r)

		if tipo == "new" && !encontrado{
			incluir = true
		}
		if tipo == "follow" && encontrado{
			incluir = true
		}

		if r.UsuarioID == ID {
			incluir = false
		}

		if incluir {
			s.Password=""
			s.Biografia=""
			s.SitioWeb=""
			s.Email=""
			s.Banner=""
			s.Ubicacion=""

			results = append(results, &s)
		}
	}

	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
	
}