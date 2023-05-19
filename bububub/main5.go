package main

import (
	"fmt"
	"math"
)

func main() {
	var c, r, p, s float64
	fmt.Scan(&r)
	p = 3.14

	c = 2 * p * r
	s = p * math.Pow(r, 2)

	fmt.Println(s)
	fmt.Println(c)
}
