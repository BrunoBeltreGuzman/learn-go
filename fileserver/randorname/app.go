package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func TempFileName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}

func main() {

	name := TempFileName()
	fmt.Println(name)
}
