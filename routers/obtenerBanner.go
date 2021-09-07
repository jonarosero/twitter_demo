package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/jonarosero/twitter_demo/bd"
)

func ObtenerBanner (w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")

	if len (ID) <1 {
		http.Error(w, "Debe enviar el paramétro id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err !=nil{
		http.Error(w, "Usuario no encontrado", 404)
		return
	}

	OpenFile, err := os.Open("uploads/banners/"+ perfil.Banner)

	if err != nil {
		http.Error(w, "Imagen no encontrada", 404)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}