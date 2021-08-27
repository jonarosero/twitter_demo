package bd

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexión a la base de datos
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user:twittorpass@cluster0.pnsxo.mongodb.net/twittor?authSource=admin")

//ConectarBD permite conectar a la BD de mongo
func ConectarBD() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)

	if  err !=nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa a la BD")
	return client
}
//ChequeoConnection es el Ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}