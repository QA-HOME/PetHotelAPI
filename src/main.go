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

	log.Default().Print("Server localhost:8081 is started")

	defer func(addr string, handler http.Handler) {
		err := http.ListenAndServe(addr, handler)
		errors.CheckErr(err)
	}(":8081", r)
}
