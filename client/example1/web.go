package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.html")
}

type Raw struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func myApi(w http.ResponseWriter, req *http.Request) {

	var data Raw
	data.Id = 1
	data.Name = "Bobo"

	json.NewEncoder(w).Encode(data)

}

func main0() {

	//Server config
	const PORT string = ":8080"

	//Routers
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/home", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Home")
	})

	http.HandleFunc("/api", myApi)

	//Server start
	http.ListenAndServe(PORT, nil)
}
