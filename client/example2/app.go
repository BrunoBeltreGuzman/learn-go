package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Server config
	const PORT string = "8080"

	http.Handle("/", http.FileServer(http.Dir("./static/")))

	//Server start
	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		fmt.Println("Starting server at port" + PORT)
		fmt.Println("http://localhost:" + PORT + "/")
	}
}
