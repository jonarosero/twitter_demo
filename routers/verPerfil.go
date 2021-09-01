package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
)

//VerPerfil permite extraer los valores del Perfil
func VerPerfil (w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un registro ", 400)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}