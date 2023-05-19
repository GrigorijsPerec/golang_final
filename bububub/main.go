package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var s int = a * b
	var p int = a + b + a + b
	fmt.Println("S - ", s)
	fmt.Println("P - ", p)
}
