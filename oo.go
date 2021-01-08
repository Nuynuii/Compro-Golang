package main

import "fmt"

type student struct {
	name string
	age int
	email string
}

func main() {
	std := student{name: "nuy"}
	n := &std
	(*n).age = 19
	n.email = "nn.@waran.ya"

	fmt.Println(std)
}

