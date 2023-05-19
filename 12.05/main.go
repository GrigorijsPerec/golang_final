package main

import (
	"fmt"
)

const (
	width  = 5
	height = 4
)

func main() {
	var arr [5][5]int

	for row := 0; row < width; row++ {
		for col := 0; col < height; col++ {
			if col%2 == 0 && row%2 == 0 {
				arr[row][col] = 1
			} else if col%2 == 0 && row%2 == 1 {
				arr[row][col] = 0
			} else if col%2 == 1 && row%2 == 0 {
				arr[row][col] = 0
			} else {
				arr[row][col] = 1
			}

		}
	}
	for row := 0; row < width; row++ {
		for col := 0; col < height; col++ {
			fmt.Print(arr[row][col], " ")
		}
		fmt.Println()
	}

}
