package main

import (
	"fmt"
	"strings"
)

func main() {
	var a strings.Builder
	a.WriteString("Hello")
	a.WriteString(" ")
	a.WriteString("World")
	fmt.Println(a.String())

}
