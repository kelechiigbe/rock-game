package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelechiigbe/rock-game/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterGameRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
