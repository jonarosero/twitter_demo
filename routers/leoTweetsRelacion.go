package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jonarosero/twitter_demo/bd"
)

func LeoTweetsSeguidores (w http.ResponseWriter, r *http.Request){
	
	if len (r.URL.Query().Get("pagina"))<1{
		http.Error(w, "Debe enviar la página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parametro página mayor a 0" , http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)

	if correcto == false {
		http.Error(w, "error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(respuesta)

}