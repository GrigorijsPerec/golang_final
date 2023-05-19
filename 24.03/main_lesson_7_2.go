package main

import (
	"fmt"
)

func main() {

	array := [10]int{0}
	//b := 3

	for i := 0; i < len(array); i++ {
		var c int
		fmt.Scan(&c)
		array[i] = c
	}
	for i := 0; i <= 10; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
