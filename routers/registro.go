package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
	"github.com/jonarosero/twitter_demo/models"
)

//Registro es la funcion para crear en la BD el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+ err.Error(),400)
		return
	}

	if len(t.Email)==0{
		http.Error(w, "Email de usuario requerido", 400)
		return
	}
	if len(t.Password) < 6{
		http.Error(w, "Password debe ser mayor a 6 caractéres", 400)
		return
	}

	_, econtrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if econtrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar al usuario"+ err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logro insertar el registro de usuario"+ err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}