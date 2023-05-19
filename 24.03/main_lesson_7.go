package main

import (
	"fmt"
)

func main() {

	//var slice int = 3
	var listRange int
	array := [100]int{0}
	fmt.Scan(&listRange)

	for i := 0; i < listRange; i++ {
		array[i] = i + 1
	}
	fmt.Println(array)

	for i := 0; i < len(array); i += 3 {
		if i < listRange {
			array[i] = 0
		} else {
			break
		}
	}

	fmt.Println(array)

}
