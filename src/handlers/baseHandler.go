package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
	"v1/src/errors"
)

func SetContentJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ValidateToken(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {

			tokenString := r.Header["Authorization"][0]
			tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					SendUnAuthWrite(w)
				}
				return SECRET, nil
			})

			if err != nil {
				SendUnAuthWrite(w)
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			log.Default().Print("Header.Token is empty")
			SendUnAuthWrite(w)
		}
	})
}

func SendUnAuthWrite(w http.ResponseWriter) {

	SetContentJson(w)

	log.Default().Print("unauthorized trying")

	w.WriteHeader(http.StatusUnauthorized)
	m := make(map[string]string)
	m["status"] = "unauthorized"
	m["message"] = "you are unauthorized"
	jsonM, _ := json.Marshal(m)

	_, err := w.Write(jsonM)
	errors.CheckErr(err)
}

func NotPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {

		SendUnAuthWrite(w)

		log.Panic("not post")
	}
}
