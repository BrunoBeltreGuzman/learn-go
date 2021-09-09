package main

import "fmt"

//Suma de punteros
func main2() {

	var aTem int = 10
	var bTem int = 15

	/*
		& = valor en memoria
	*/
	var a *int = &aTem
	var b *int = &bTem
	tem := *a + *b + *b
	r := suma(a, b)

	fmt.Println("Suma: ", r)
	fmt.Println("Tem: ", tem)

}

func suma(a *int, b *int) *int {
	tem := *a + *b
	return &tem
}
