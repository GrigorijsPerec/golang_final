package main

import "fmt"

func divide2(n uint) int {
	if n%2 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	var n uint
	fmt.Scan(&n)
	result := divide2(n)
	fmt.Println(result)
}
