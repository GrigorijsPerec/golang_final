package main

import (
	"fmt"
	"strings"
)

func MedStr(n int) {
	switch {
	case n == 1:
		fmt.Println(strings.Repeat("1", n))
	case n == 2:
		fmt.Println(strings.Repeat("2", n))
	case n == 3:
		fmt.Println(strings.Repeat("3", n))
	case n == 4:
		fmt.Println(strings.Repeat("4", n))
	case n == 5:
		fmt.Println(strings.Repeat("5", n))
	case n == 6:
		fmt.Println(strings.Repeat("6", n))
	case n == 7:
		fmt.Println(strings.Repeat("7", n))
	case n == 8:
		fmt.Println(strings.Repeat("8", n))
	case n == 9:
		fmt.Println(strings.Repeat("9", n))

	}

}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		MedStr(n)
	}

}
