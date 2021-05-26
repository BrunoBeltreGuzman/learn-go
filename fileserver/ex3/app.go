package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	out, err := os.Create("output.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get("http://example.com/")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)

}
