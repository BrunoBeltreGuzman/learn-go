package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetImage(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadFile("descarga.png")
	str := base64.StdEncoding.EncodeToString(b)
	json.NewEncoder(w).Encode(str)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/image", GetImage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
