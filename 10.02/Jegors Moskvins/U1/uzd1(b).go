package main

import "fmt" 

func main() {
	var system int
	var chislo string
	var key = 'A'
	mapa := make(map[rune]int)
	for value := 10; value != 36; value++ {
		mapa[key] += value
		key++
	}
	fmt.Println("Число:")
	fmt.Scan(&chislo)
	fmt.Println("Система исчисления - не меньше 10 и не больше 36")
	fmt.Scan(&system)
	chislo2 := []rune(chislo)
	if Proverka(system, chislo2, mapa) {
		fmt.Println("Число записано корректно")
	}else{
		fmt.Println("Число записано не корректно")
	}
}
func Proverka(system int, chislo []rune, mapa map[rune]int)bool { 
	for _, num := range chislo {
		if int(num)<10 {
			if int(num) - '0' >= system {
				return false
			}
		} else if int(num)>'A' {
			if mapa[num] >= system {
				return false
			}
		}
	}
	return true
}
