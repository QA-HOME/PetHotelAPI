package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"v1/src/errors"
)

func Index(w http.ResponseWriter, r *http.Request) {

	SetContentJson(w)

	log.Default().Print("success")

	m := make(map[string]string)
	m["status"] = "success"
	m["message"] = "home page"

	jsonM, _ := json.Marshal(m)

	var _, err = w.Write(jsonM)

	errors.CheckErr(err)
}
