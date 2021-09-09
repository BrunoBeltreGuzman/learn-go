package main

import "fmt"

/*
	Closures
	Go admite funciones anónimas , que pueden formar cierres . Las funciones anónimas son útiles cuando desea definir una función en línea sin tener que nombrarla.
*/

func getFunc() func() int {
	var a int = 2
	return func() int {
		return a + 2
	}
}

func runFunc(f func()) {
	f()
}

func main() {

	var f func() int = getFunc()
	var r int = f()
	fmt.Println(r)

	var newFunc func() = func() {
		fmt.Println("The func")
	}
	runFunc(newFunc)
}
