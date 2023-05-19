package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	var a, b, c, s float64
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	a = math.Abs(x1 - x2)
	fmt.Println(a)
	b = math.Abs(y1 - y2)
	fmt.Println(b)
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	s = a * b
	fmt.Println(c, s)
}
