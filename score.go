package main

import "fmt"

func main() {
	var score float64 
	fmt.Scan(&score)
	if score >= 80 && score <=100 {
		fmt.Println("A")
	}
	if score >= 75 && score < 79 {
		fmt.Println("B+")
	}
	if score >=70 && score < 75 {
		fmt.Println("B")
	}
	if score >=65 && score < 70 {
		fmt.Println("C+")
	}
	if score >=60 && score < 65 {
		fmt.Println("C")
	}
	if score >=55 && score < 60 {
		fmt.Println("D+")
	}
	if score >=50 && score < 55 {
		fmt.Println("D")
	}
	if score >=0 && score < 50 {
		fmt.Println("F")
	}
}
