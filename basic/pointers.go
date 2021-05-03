package main

import (
	"fmt"
)

func main3() {

	var data string = "Bruno"
	var name *string //declaracion de puntero

	fmt.Println(name)
	fmt.Println(data)

	name = &data // Asignacion de valor en mamoria

	fmt.Println(*name) //print valor en mamoria

	updateValue(name, "Beltre")
	fmt.Println(*name) //print new valor en mamoria

	updateVariable(*name, "Guzman")
	fmt.Println(*name)
}

func updateValue(variable *string, data string) {
	*variable = data // update value en memoria
}

func updateVariable(variable string, data string) {
	variable = data // (se crea nueva variable)
}
