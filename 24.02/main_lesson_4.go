package main

import (
	"fmt"
)

func task1() {

	a := 0
	b := 1

	for b <= 1000000 {
		fmt.Println(b)

		a, b = b, a+b

	}
}

func task2() {

}

func main() {
	task1()
}
