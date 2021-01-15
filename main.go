package main

import (
	"fmt"
	"os"
)

func main() {
	// create file
	file , err := os.Create("nuynuii.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("Hello")
	file.Close()
}