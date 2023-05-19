package main

import "fmt"

func main(){
	var system int
	var chislo string
	fmt.Println("Число:")
	fmt.Scan(&chislo)
	fmt.Println("Система исчисления")
	fmt.Scan(&system)
	chislo2 := []rune(chislo)
	if Proverka(system, chislo2) {
		fmt.Println("Число записано корректно")
	}else{
		fmt.Println("Число записано не корректно")
	}
}
func Proverka(system int, chislo []rune)bool { 
	for _, num := range chislo {
		if int(num) - '0' >= system {
			return false
		}
	}
	return true
}
