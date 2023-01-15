package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
	"v1/src/auth"
	"v1/src/errors"
)

var SECRET = []byte("super-secret-auth-key")

func GetToken(w http.ResponseWriter, r *http.Request) {

	SetContentJson(w)

	if r.Header["X-Api-Key"] != nil {

		if auth.CheckPasswordHash(r.Header["X-Api-Key"][0]) {

			token, err := CreateToken()
			errors.CheckErr(err)

			m := make(map[string]string)
			m["token"] = token
			m["status"] = "success"

			jsonS, _ := json.Marshal(m)

			_, err = w.Write(jsonS)
			errors.CheckErr(err)

		} else {
			SendUnAuthWrite(w)
		}
	} else {
		SendUnAuthWrite(w)
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
