package main

import "fmt"

func divide2(n uint) int {
	if n%2 == 0 {
		return 0
	} else {
		return 1
	}
}

func main() {
	var n uint
	fmt.Scan(&n)
	result := divide2(n)
	fmt.Println(result)
}
