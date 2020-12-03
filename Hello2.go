package main

import "fmt"

func main() {
	fmt.Print(`\n \t Backticks`)
	fmt.Print("\n \t Double quote")
	fmt.Println()
	fmt.Println(`\n \t Backticks`[5])
	fmt.Println(len(`\n \t Backticks`))
	fmt.Println("Hello" + "World")

}
