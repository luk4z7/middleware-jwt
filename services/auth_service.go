package services

import (
	"api.jwt.auth/api/parameters"
	"api.jwt.auth/core/authentication"
	"api.jwt.auth/services/models"
	"encoding/json"
	"net/http"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token := parameters.TokenAuthentication{authBackend.GenerateToken()}
		response, _ := json.Marshal(token)
		return http.StatusOK, response
	}

	return http.StatusUnauthorized, []byte("")
}

func RefreshToken() []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token := parameters.TokenAuthentication{authBackend.GenerateToken()}
	response, err := json.Marshal(token)
	if err != nil {
		panic(err)
	}
	return response
}

/*func Logout() {

}*/