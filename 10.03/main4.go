package main

import (
	"fmt"
	"strings"
)

func FandLstr(n int) {
	fmt.Println(strings.Repeat("x", n))
}

func MedStr(n int) {
	fmt.Println("x" + strings.Repeat(" ", n) + "x")
}

func main() {
	n := 7
	FandLstr(n)
	for i := 0; i < n; i++ {
		MedStr(n)
	}
	FandLstr(n)
}
