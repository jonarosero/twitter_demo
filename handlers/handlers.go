package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jonarosero/twitter_demo/middlew"
	"github.com/jonarosero/twitter_demo/routers"
	"github.com/rs/cors"
)

//Manejadores setea el puerto, el handler y pone a escuchar al servidor
func Manejadores () {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultoRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultoRelacion))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	router.HandleFunc("/registroRol", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(routers.RegistroRol)))).Methods("POST")
	router.HandleFunc("/eliminarRol", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(routers.EliminarRol)))).Methods("DELETE")
	router.HandleFunc("/modificarRol", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(routers.ModificarRolUsuario)))).Methods("PUT")
	router.HandleFunc("/verRol", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerRol))).Methods("GET")
	router.HandleFunc("/listaRoles", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaRoles))).Methods("GET")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}