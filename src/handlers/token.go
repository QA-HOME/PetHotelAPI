package handlers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
	"v1/src/auth"
	"v1/src/errors"
)

var SECRET = []byte("super-secret-auth-key")

func GetToken(w http.ResponseWriter, r *http.Request) {

	if r.Header["X-Api-Key"] != nil {

		if auth.CheckPasswordHash(r.Header["X-Api-Key"][0]) {

			token, err := CreateToken()
			errors.CheckErr(err)

			_, err = fmt.Fprint(w, token)
			errors.CheckErr(err)
		} else {
			return
		}
	} else {
		log.Default().Print("x-api-key is empty")
		log.Default().Print(r.Header)
	}
}

func CreateToken() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)
	errors.CheckErr(err)

	return tokenStr, nil
}
