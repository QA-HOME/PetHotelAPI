package main

import (
	"log"
	"net/http"
	"v1/src/errors"
	"v1/src/handlers"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.Index)

	log.Default().Print("8081 is started")

	err := http.ListenAndServe(":8081", r)
	errors.CheckErr(err)
}
