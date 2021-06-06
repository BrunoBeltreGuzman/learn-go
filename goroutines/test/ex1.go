package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch() {
	url := "https://tecnodigital.vercel.app/"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}

func main() {
	for i := 0; i < 1000; i++ {
		go fetch()
	}

	for i := 0; i < 1000; i++ {
		go fetch()
	}

	var message string
	fmt.Scan(&message)
}
