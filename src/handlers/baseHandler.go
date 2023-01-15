package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"v1/src/errors"
)

func SetContentJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ValidateToken(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					m := make(map[string]string)
					m["status"] = "unauthorized"
					m["message"] = "you are unauthorized"
					jsonM, _ := json.Marshal(m)

					_, err := w.Write(jsonM)
					errors.CheckErr(err)
				}
				return SECRET, nil
			})

			if err != nil {

				w.WriteHeader(http.StatusUnauthorized)
				m := make(map[string]string)
				m["status"] = "unauthorized"
				m["message"] = "you are unauthorized"
				jsonM, _ := json.Marshal(m)

				_, err := w.Write(jsonM)
				errors.CheckErr(err)
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			m := make(map[string]string)
			m["status"] = "unauthorized"
			m["message"] = "you are unauthorized"
			jsonM, _ := json.Marshal(m)

			_, err := w.Write(jsonM)
			errors.CheckErr(err)
		}
	})
}
