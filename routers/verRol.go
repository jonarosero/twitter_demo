package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
)

//VerPerfil permite extraer los valores del Rol
func VerRol (w http.ResponseWriter, r *http.Request) {

	rol := r.URL.Query().Get("rol")


	informacion, err := bd.BuscoRol(rol)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un rol ", 400)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}