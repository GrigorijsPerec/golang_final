package main

import (
	"fmt"
)

func removeChar(str string, char string) string {
	var result string
	for i := 0; i < len(str); i++ {
		if string(str[i]) != char {
			result += string(str[i])
		}
	}
	return result
}

func main() {
	var str string
	var char string
	fmt.Scan(&str)
	fmt.Scan(&char)

	newStr := removeChar(str, char)
	fmt.Println(newStr)
}
