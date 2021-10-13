package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
	"github.com/jonarosero/twitter_demo/models"
)

func ModificarRolUsuario (w http.ResponseWriter, r *http.Request){

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(),400)
		return
	}

	var status bool
	status, err = bd.ModificoRolUsuario(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el rol del usuario "+err.Error(),400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el rol del usuario ",400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
}