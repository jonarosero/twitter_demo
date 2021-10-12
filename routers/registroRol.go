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

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+ err.Error(),400)
		return
	}

	//registro := models.Rol{
		//Rol: rol.Rol,
	//}

	_, econtrado, _ := bd.ChequeoYaExisteRol(rol.Rol)
	if econtrado {
		http.Error(w, "Ya existe el rol"+rol.Rol, 400)
		return
	}

	_, status, err := bd.RegistroRol(rol)

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