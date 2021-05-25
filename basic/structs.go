package main

type Finalsalary func(int, int) int

func test() {

	type person struct {
		name   string
		salary Finalsalary
	}

	var bobo person

	bobo.name = "Bobo"
	bobo.salary = func(Ma int, pay int) int {
		return Ma * pay
	}

	println(bobo.name)
	println(bobo.salary(10, 20))

}
