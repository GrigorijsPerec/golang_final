package main

import "fmt"

func prime(n uint) uint {
	if n/n == n/1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	var n uint
	fmt.Scan(&n)
	result := prime(n)
	fmt.Println(result)
}
