package main

import (
	"fmt"
	"math"
	"sort"
)

func task1() {
	var s, v, v1, t float64

	fmt.Scan(&s)
	fmt.Scan(&v)
	fmt.Scan(&v1)
	t = s / (v + v1)

	fmt.Println(t)
}

func task2() {
	s := []int{4, 2, 3, 6, 7, 4, 2, 5, 7, 4, 3, 6, 8, 2, 2, 2, 2, 2}
	sort.Ints(s)
	fmt.Println(s)
}

func task3() {
	var day, hours, min, sec float64

	fmt.Scan(&sec)
	day = sec / 86400
	hours = sec / 3600
	min = sec / 60

	println(day, hours, min, sec)
}

func task4() {
	var year int = 3

	if year%4 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}
}

func task5() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	d := math.Pow(b, 2) - 4*a*c
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)

	fmt.Println("Roots:", x1, x2)
}

func task6() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	d := math.Pow(b, 2) - 4*a*c
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)

	fmt.Println("Roots:", x1, x2)
}

func task7() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	d := math.Pow(b, 2) - 4*a*c
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)

	fmt.Println("Roots:", x1, x2)
}

func main() {
	//task1()
	task2()
	//task3()
	//task4()
	//task5()
	//task6()
}
