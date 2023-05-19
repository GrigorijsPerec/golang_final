package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var d, q, i int
	fmt.Scan(&d)
	if d%2 != 0 {
		return
	}
	for i = 2; q < 10; i++ {
		if isPrime(i) && isPrime(d+i) {
			fmt.Println(i, d+i)
			q++
		}
	}
}
