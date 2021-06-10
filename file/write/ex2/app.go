package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func newFile(filename string, content string) bool {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
		return false

	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
		return false

	}

	fmt.Println("done")

	return true
}

func TryErro(err error, response http.ResponseWriter) {
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid File Name")
		return
	}
}

func GetDesign(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	filename := params["filename"]
	fmt.Println(filename)

	if filename == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid File Name")
		return
	}

	data, err := ioutil.ReadFile(filename)
	TryErro(err, response)
	response.WriteHeader(http.StatusOK)
	fmt.Fprint(response, string(data))
	return
}

func NewDesign(response http.ResponseWriter, request *http.Request) {

	error := request.ParseForm()
	if error != nil {
		log.Fatal(error)
		return
	}

	filename := request.FormValue("filename")
	content := request.FormValue("content")
	fmt.Println(filename)
	fmt.Println(content)

	if filename == "" && content == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400, Invalid Inputs")
		return
	}

	boo := newFile(filename, content)
	if boo == true {
		result := map[string]string{
			"code":     "200",
			"message":  "The request has succeeded.",
			"response": "File Create successfully",
		}
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(result)
		return
	} else {
		result := map[string]string{
			"code":     "500",
			"message":  "Error 500",
			"response": "Error orcurred in create file",
		}
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(result)
		return
	}

}

func main() {
	routers := mux.NewRouter()

	//index
	routers.HandleFunc("/getdesign/{filename}", GetDesign).Methods("GET")
	routers.HandleFunc("/newdesign", NewDesign).Methods("POST")

	server := &http.Server{
		Addr:           ":8000",
		Handler:        routers,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server starting successfully")
	fmt.Println("http://localhost:8000")
	log.Fatal(server.ListenAndServe())
}
