package main

import "fmt"

func counter(str string, cifr string) int {
	count := 0

	for i := 0; i < len(str); i++ {
		if string(str[i]) == cifr {
			count++
		}
	}
	return count
}

func main() {
	var str string
	var cifr string
	fmt.Scan(&str)
	fmt.Scan(&cifr)
	count := counter(str, cifr)

	fmt.Print(count)

}
