package routers

import (
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
)

func EliminarTweet (w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")

	if len(ID)<1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error" + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}