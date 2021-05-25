package main

import "fmt"

func main() {

	canal := make(chan string)

	go ok(canal)

	message := <-canal

	fmt.Println(message)
}

func ok(canal chan string) {

	canal <- "Hello World!!"
}
