package main

import (
	"fmt"
)

func main() {
	task1()
}

func task1() {
	array := []int{1, 3, 3, 3, 1, 2, 2, 5, 5}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] == array[i] {
				array = append(array[0:j], array[j+1:len(array)]...)
				j--
			}
		}n 
	}
	fmt.Println(array)
}

func task2() {
	array := []int{1, 2, 5, 16}
	array1 := []int{8, 9, 25}
	array2 := [7]int{}
}
