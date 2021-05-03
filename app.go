package main

import (
	"fmt"
)

func main2() {

	//Hello world
	fmt.Println("Hello world")

	//String
	var name string = "Bobo"
	fmt.Println(name)
	var calificacion string = "A"
	fmt.Println(calificacion)

	//Numbers
	//Int
	var num1 int = 10
	var num2 int = 3
	var suma = num1 + num2
	fmt.Println(suma)

	//var num5 byte = 6233333 Error
	var num5 byte = 255
	fmt.Println(num5)

	//Float
	var ip float32 = 3.14
	fmt.Println(ip)
	var ipLong float64 = 3.1400000000000000000000000000000000000000
	fmt.Println(ipLong)

	//Bools
	var isBool bool = true
	if isBool {
		fmt.Println("is Bool", isBool)
	} else {
		fmt.Println("is not Bool", isBool)
	}
	if isBool == true {
		fmt.Println("is Bool", isBool)
	} else {
		fmt.Println("is not Bool", isBool)
	}

	//Const
	const PI = 3.14
	fmt.Println(PI)

	//Funciones
	func1()

	//For

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j, ")", "Bobo")
	}

	for x := 0; x < 60; x++ {
		if x == 6 {
			continue
		}
		if x == 50 {
			break
		}
		fmt.Println(x, ")", "Ok")
	}

	var isTrue bool = 5 > 1

	if isTrue {
		fmt.Println("is True")
	} else {
		fmt.Println("is False")
	}

	if isTrue == false {
		fmt.Println("is True")
	} else {
		fmt.Println("is False")
	}

	if 5 > 3 {
		fmt.Println("is True")
	} else {
		fmt.Println("is False")
	}

	//switch
	var swit int = 2

	switch swit {

	case 1:
		fmt.Println("Case 1")

	case 2:
		fmt.Println("Case 2")

	case 3:
		fmt.Println("Case 3")

	default:
		fmt.Println("Case default")
	}

	//Array
	var numbers [5]int
	numbers[0] = 10
	fmt.Println(numbers[0])

	floats := [5]int{1, 2, 3, 4, 5}
	fmt.Println(floats[0])
	numbers[1] = floats[4]

	//Slices
	s := make([]string, 3)

	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("get:", s[2])
	fmt.Println("set:", s)
	fmt.Println("len:", len(s))

	my := []int{1, 2, 3, 4}
	fmt.Println(my)

}

func func1() {
	fmt.Println("app")
}
