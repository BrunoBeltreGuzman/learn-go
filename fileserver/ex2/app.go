package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	url := "http://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(body))
	//download the file in browser
}

func main() {
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
