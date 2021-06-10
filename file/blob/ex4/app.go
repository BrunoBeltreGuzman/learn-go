package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("newfile.txt")
	reader := bufio.NewReader(file)
	data, _ := reader.Peek(50)
	fmt.Println(string(data))
	fmt.Println(reader)
	file.Close()

}
