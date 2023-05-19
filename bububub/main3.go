package main

import "fmt"

func main() {
	var km, h float32
	var ms float32
	var mmin float32

	fmt.Scan(&km)
	fmt.Scan(&h)

	ms = (km / 1000) / (h / 3600)
	mmin = km * 16.6

	fmt.Println(ms)
	fmt.Println(mmin)
}
