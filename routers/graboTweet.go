package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jonarosero/twitter_demo/bd"
	"github.com/jonarosero/twitter_demo/models"
)

func GraboTweet (w http.ResponseWriter, r *http.Request){
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if  err != nil {
		http.Error(w, "Ocurrio un error al insertar el tweet"+err.Error(), 400)
		return
	}

	if !status{
		http.Error(w, "No se ha logrado insertar el tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)
}