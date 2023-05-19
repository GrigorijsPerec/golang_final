package main

import (
	"fmt"
)

const (
	width  = 4
	height = 8
)

/*
[

	[0 1 2 3 4 ]
	[5 6 7 8 9 ]
	[10 11 12 13 14 ]
	[15 16 17 18 19 ]
	[20 21 22 23 24 ]

]
*/
func main() {
	var arr [width][height]int

	for row := 0; row < width; row++ {
		for col := 0; col < height; col++ {
			if row%2 == 1 {
				arr[row][col] = 1
			}
			if col%2 == 1 {
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
