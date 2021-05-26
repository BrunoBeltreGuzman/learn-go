package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetImage(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadFile("descarga.png")
	str := base64.StdEncoding.EncodeToString(b)
	json.NewEncoder(w).Encode(str)
}

func main() {
	router := gin.Default()
	router.Static("/image", "descarga.png")
	log.Fatal(http.ListenAndServe(":8000", router))
}
