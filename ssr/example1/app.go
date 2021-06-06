package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {

	//Server config
	const PORT string = ":8080"

	routers := mux.NewRouter()
	routers.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("web").HTTPBox()))

	//Server start
	log.Fatal(http.ListenAndServe(PORT, routers))
}
