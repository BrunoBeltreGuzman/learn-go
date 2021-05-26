package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

func TempFileName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}

func fetch() {
	fileUrl := "https://golangcode.com/logo.svg"

	var path string = "filepath/" + TempFileName() + ".svg"
	err := DownloadFile(path, fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + path)
}

func main() {
	fetch()
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
