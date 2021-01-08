package main

import (
	"unicode"
)
func countLowercaseLetters(text string) int {
	/*
		คืนค่าจำนวนตัวอักษรพิมพ์เล็กที่อยู่ในตัวแปร text
	*/
	count :=0
	for _, char := range text {
		if unicode.IsLower(char) {
			count++
		}
	}
	return count
}

func countUppercaseLetters(text string) int {
	/*
		คืนค่าจำนวนตัวอักษรพิมพ์ใหญ่ที่อยู่ในตัวแปร text
	*/
	count := 0
	for _, char := range text {
		if unicode.IsUpper(char) {
			count++
		}
	}
	return count
}

func countUpperAndLowercaseLetters(text string) (numberOfUppercaseLetters int, numberOfLowercaseLetters int) {
	for _, chars := range text {
		if unicode.IsLower(chars) {
			numberOfLowercaseLetters++
		}
		if unicode.IsUpper(chars) {
			numberOfUppercaseLetters++
		}
	}
	return numberOfUppercaseLetters, numberOfLowercaseLetters
}
func countNumbers(text string) (count int) {
	/*คืนค่าจำนวนตัวอักษรตัวเลขที่อยู่ในตัวแปร text*/
	for _, chars := range text {
		if unicode.IsNumber(chars) {
			count++
		}
	}
	return count
}

func countSymbol(text string) (count int) {
	/*คืนค่าจำนวนตัวอักษรตัวเลขที่อยู่ในตัวแปร text*/
	for _, chars := range text {
		if unicode.IsSymbol(chars) || unicode.IsPunct(chars) {
			count++
		}
	}
		return count 
}

func countMiddleNumbersOrSymbols(text string) (count int) {
	/*
		คืนค่าจำนวนตัวตัวเลขหรืออักษรสัญลักษณ์ (punctuation character and symbolic character)ที่ในตัวแปร text โดยต้องไม่เป็นตัวอักษรตัวแรกและตัวสุดท้าย
	*/
	/*for index, value := range text {
		if index != 0 && index != len(text)-1 {
			if unicode.IsSymbol(value) || unicode.IsSymbol(value) || unicode.IsPunct(value){
				count++
			}
		}
	}*/
	chars := []rune(text)
	for _, value := range chars[1 : len(chars)-1] {
		if unicode.IsNumber(value) || unicode.IsSymbol(value) || unicode.IsPunct(value) {
			count++
		}
	}
	return count
}

func isContainsSymbolNumberLowerUpper(text string, length int) bool {
	/*
		หากตัวแปร text ประกอบไปด้วย
		- punctuation character หรือ symbolic character
		- Numbers
		- Lowercase
		- Uppercase
		- Minimum [length] characters
		จะคืนค่า true
	*/
	if countNumbers(text) < 1 {
		return false
	}
	if countSymbol(text) < 1 {
		return false
	}
	if countUppercaseLetters(text) < 1 {
		return false
	}
	if countLowercaseLetters(text) < 1 {
		return false
	}
	if len(text) < length {
		return false
	}
	return true
}

