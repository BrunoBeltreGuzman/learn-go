package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
)

func handleServe(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("descarga.png")))
}

func main() {
	// http.HandleFunc("/serve/", handleServe)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
