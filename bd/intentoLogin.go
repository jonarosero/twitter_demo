package bd

import (
	"github.com/jonarosero/twitter_demo/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLOgin realiza el chequeo de login a la bd
func IntentoLogin(email string, password string) (models.Usuario, bool){
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true

}