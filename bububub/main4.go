package main

import "fmt"

func main() {
	var v, t, s int

	fmt.Scan(&v)
	fmt.Scan(&t)
	fmt.Scan(&s)

	s = v * t

	fmt.Println(s)

}
