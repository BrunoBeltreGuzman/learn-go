package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	var name string = "Bobo"
	go nombreLento(name)
	fmt.Println(name)
	time.Sleep(11111 * time.Millisecond)
}

func nombreLento(name string) {
	var letras []string = strings.Split(name, "")

	for index, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(index, ") ", letra)
	}
}
