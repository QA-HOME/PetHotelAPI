package handlers

import (
	"net/http"
)

func SetContentJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
