package main

import "fmt"

func main() {
	fmt.Print("input number : ")
	var number int
	n, err := fmt.Scan(&number)
	fmt.Println(`read number `, number, ` from standard input`)
	fmt.Println(`number of argument `, n)
	fmt.Println(`error `, err)
}

