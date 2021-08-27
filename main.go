package main

import (
	"log"
	"github.com/jonarosero/twitter_demo/bd"
	"github.com/jonarosero/twitter_demo/handlers"

)

func main (){
	if bd.ChequeoConnection() == 0{
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}