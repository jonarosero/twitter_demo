package middlew

import (
	"net/http"

	"github.com/jonarosero/twitter_demo/bd"
	"github.com/jonarosero/twitter_demo/routers"
)

//ValidoAdmin es el middleware que me permite conocer el estado del rol de los usuarios
func ValidoAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, econtrado, mensaje := bd.ChequeoUsuarioAdmin(routers.IDUsuario)
		if !econtrado {
			http.Error(w, mensaje, 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
