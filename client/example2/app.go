package main

import (
	"log"
	"net/http"
)

func index(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "pages/index.html")
}

func main() {

	//Server config
	const PORT string = ":8080"

	//Routers
	http.HandleFunc("/", index)

	http.Handle("/", http.FileServer(http.Dir("css/")))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//Server start
	log.Fatal(http.ListenAndServe(PORT, nil))
}
