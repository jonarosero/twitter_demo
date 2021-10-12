package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
	"github.com/jonarosero/twitter_demo/models"
)

func RegistroRol(w http.ResponseWriter, r *http.Request){
	var rol models.Rol

	err := json.NewDecoder(r.Body).Decode(&rol)

	registro := models.Rol{
		ID: rol.ID,
		Rol: rol.Rol,
	}

	_, status, err := bd.RegistroRol(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un nuevo rol", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar un nuevo rol", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}