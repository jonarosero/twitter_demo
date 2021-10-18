package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
)

func ListaRoles(w http.ResponseWriter, r *http.Request) {

	result, status := bd.LeoRolesTodos()
	if !status {
		http.Error(w, "Error al leer los roles", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}