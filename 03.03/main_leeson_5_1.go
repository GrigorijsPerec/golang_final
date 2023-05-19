package main

import "fmt"

func main() {
	var n uint
	fmt.Scan(&n)
	result := power2(n)
	fmt.Println(result)
}

func power2(n uint) int {
	if n == 0 {
		return -1
	}
	count := 0
	for n%2 == 0 {
		n /= 2
		count++
	}
	if n == 1 {
		return count
	}
	return -1
}
