package main

import (
	"fmt"
	"math"
)

func main() {
	var n, s, h float64
	fmt.Scan(&n)
	fmt.Scan(&s)
	h = math.Pow(n, s)
	fmt.Println(h)

}
