package handlers

import (
	"encoding/json"
	"net/http"
	"v1/src/errors"
)

func Index(w http.ResponseWriter, r *http.Request) {

	m := make(map[string]string)
	m["message"] = "home page"
	m["title"] = "home page"

	jsonM, _ := json.Marshal(m)

	var _, err = w.Write(jsonM)

	errors.CheckErr(err)
}
