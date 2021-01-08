package main

import "fmt"

type student struct {
	name string
	age int
	email string
}

func main() {
	stds := []student{
		student{"Goku", 35, "Goku@super.saiya"},
		student{"Gohan", 15, "Gohan@super.saiya"},
		student{"Videl", 14, "Videl@super.saiya"}}

	avgAge := 0
	for _, v := range stds {
		avgAge += v.age
	}
	avgAge = avgAge / len(stds)
	fmt.Println("average age is", avgAge)
}