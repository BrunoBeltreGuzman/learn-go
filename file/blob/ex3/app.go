package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func File(response http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadFile("MyDisenoFinal2.idg")
	check(err)
	fmt.Fprint(response, string(data))
}

func Blob(response http.ResponseWriter, request *http.Request) {
	file, _ := os.Open("MyDisenoFinal2.idg")
	reader := bufio.NewReader(file)
	data, _ := reader.Peek(50)
	fmt.Println(string(data))
	fmt.Println(reader)
	file.Close()
	fmt.Fprint(response, reader)
}

func main() {

	http.HandleFunc("/file", File)
	http.HandleFunc("/blob", Blob)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}

}
