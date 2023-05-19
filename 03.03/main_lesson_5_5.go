package main

import (
	"fmt"
	"math"
)

func PerfectSquare(n uint) (ok uint, root2 int) {
	root := math.Sqrt(float64(n))
	if root == float64(int64(root)) {
		return 1, int(root)
	}
	return 0, 0
}

func main() {
	var n uint
	fmt.Scan(&n)
	is, root := PerfectSquare(n)
	fmt.Println(is, root)
}
