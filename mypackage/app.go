package main

import (
	"fmt"
	"mypackage/mypack"
)

func main() {
	var message string = mypack.MyFunc()
	fmt.Println(message)
}
