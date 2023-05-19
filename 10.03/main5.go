package main

import (
	"fmt"
	"strings"
)

func MedStr(n int) {
	fmt.Println(strings.Repeat("x", n))

}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		MedStr(n)
	}

}
