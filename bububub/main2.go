package main

import "fmt"

func main() {
	var a, b, c int
	var v, saa int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	v = a * b * c
	saa = 2 * (a*b + a*c + b*c)

	fmt.Println(v)
	fmt.Println(saa)
}
