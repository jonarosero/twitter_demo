package routers

import (
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
)

func EliminarRol (w http.ResponseWriter, r *http.Request){

	rolID := r.URL.Query().Get("id")

	_, econtrado, mensaje := bd.ChequeoUsuarioAdmin(IDUsuario)
	if !econtrado {
		http.Error(w, mensaje, 400)
		return
	}

	if len(rolID)<1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := bd.BorroRol(rolID)
	if err != nil {
		http.Error(w, "Ocurrio un error" + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}