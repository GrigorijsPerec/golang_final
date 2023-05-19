package main

import (
	"fmt"
	"math"
	"sort"
)

func task1() {
	var choise int
	println("Enter choise\n")
	fmt.Scan(&choise)
	if choise == 1 {
		var v, t, s int

		fmt.Scan(&v)
		fmt.Scan(&t)
		fmt.Scan(&s)

		s = v * t

		fmt.Println(s)
	} else if choise == 2 {
		var km, h float32
		var ms float32
		var mmin float32

		fmt.Scan(&km)
		fmt.Scan(&h)

		ms = (km / 1000) / (h / 3600)
		mmin = km * 16.6

		fmt.Println(ms)
		fmt.Println(mmin)
	} else if choise == 3 {
		var a, b, c int
		var v, saa int

		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)

		v = a * b * c
		saa = 2 * (a*b + a*c + b*c)

		fmt.Println(v)
		fmt.Println(saa)
	} else if choise == 4 {
		var a int
		fmt.Scan(&a)
		var b int
		fmt.Scan(&b)
		var s int = a * b
		var p int = a + b + a + b
		fmt.Println("S - ", s)
		fmt.Println("P - ", p)
	}
}

func task2() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a == b {
		print(0)
	} else if a >= b {
		print(a)
	} else if b >= a {
		print(b)
	}
}

func task3() {
	var a, b, c, p, S float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a+b > c && a+c > b && c+b > a {
		p = (a + b + c) / 2
		S = math.Sqrt(p * (p - a) * (p - b) * (p - c))
		fmt.Println(p)
		fmt.Println(S)
	} else {
		fmt.Println("Non a triangle")
	}
}

func task4() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if math.Pow(a, 2) > math.Pow(b, 2)+math.Pow(c, 2) {
		fmt.Println(">90")
	} else if math.Pow(a, 2) == math.Pow(b, 2)+math.Pow(c, 2) {
		fmt.Println("90")
	} else if math.Pow(a, 2) < math.Pow(b, 2)+math.Pow(c, 2) {
		fmt.Println("<90")
	}
}

func task5() {

	var nums [5]float64
	for i := 0; i < 5; i++ {
		fmt.Scan(&nums[i])
	}

	sort.Float64s(nums[:])
	min := nums[0]
	max := nums[4]
	median := nums[2]

	fmt.Print(min)
	fmt.Print(max)
	fmt.Print(median)
}
func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
}
