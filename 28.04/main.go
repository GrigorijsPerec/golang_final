package main

import (
	"fmt"
)

func main() {
	orders := map[string]int{
		"Ethan":    2,
		"Ava":      3,
		"Lucas":    1,
		"Mia":      5,
		"Noah":     2,
		"Emma":     4,
		"Jackson":  1,
		"Sophia":   3,
		"Aiden":    2,
		"Isabella": 1,
	}

	fmt.Println("Total customers:", len(orders))

	totalOrders := 0
	for _, value := range orders {
		totalOrders += value
	}
	fmt.Println("Total orders:", totalOrders)

	for key, value := range orders {
		fmt.Printf("%s made %d order(s)\n", key, value)
	}
}
