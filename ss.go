package main

import "fmt"

type student struct {
	name string
	age int
	email string
}

func main() {
	var a student
	a.name = "Nuy"
	a.age = 19
	a.email = "Nuy@waran.nn"

	b := student{"non", 12, "non@waran.nn"}

	c := student{email: "non@waran.ya"}

	d := student{age: 14}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}