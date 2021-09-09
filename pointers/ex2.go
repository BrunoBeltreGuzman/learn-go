package main

import "fmt"

//Cambiar valor de puntero
func main() {

	var name *string = nil

	cambiarValor(name)

	fmt.Println("Valor:", &name)

}

func cambiarValor(v *string) {
	var tem string = ""
	v = &tem
}
