package main

import "fmt"

func main() {
	name := "abccbaa"
	rev_name := Reverse1(name)
	if rev_name == name {
		fmt.Println(true)

	} else {
		fmt.Println(false)
	}
}

func Reverse1(s string) (res string) {
	for _, x := range s {
		res = string(x) + res
	}
	return
}
