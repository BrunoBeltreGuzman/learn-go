package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetDetail(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadFile("file.json")

	rawIn := json.RawMessage(string(b))
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(rawIn, &objmap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(objmap)

	json.NewEncoder(w).Encode(objmap)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/detail", GetDetail).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
