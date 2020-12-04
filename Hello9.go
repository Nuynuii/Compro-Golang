package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Waranya", "a", "x", 2))
	fmt.Println(strings.Replace("Waranya", "a", "x", -1))
}
