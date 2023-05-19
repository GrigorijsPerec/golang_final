package main

import ("fmt")

func main() {
	var chislo string
	p16 := map[string]string {
		"0000":"0",
		"0001":"1", 
		"0010":"2", 
		"0011":"3", 
		"0100":"4", 
		"0101":"5", 
		"0110":"6", 
		"0111":"7", 	
		"1000":"8", 
		"1001":"9", 
		"1010":"A",
		"1011":"B",
		"1100":"C",
		"1101":"D",
		"1110":"E",
		"1111":"F",
}	
	
	//Программа переводит число из двоичной системы в 16 ричную
	fmt.Println("Введите число:")
	fmt.Scan(&chislo)
	if len(chislo)%4!=0{
		Dobavka(chislo)
		Perevod(chislo,p16)
	} else{
		Perevod(chislo,p16)
	}
}
func Dobavka(chislo string)string {
	var chislo2 string
	for y:=16-len(chislo);y!=0;y-- {
		chislo2 += "0"
	}
	chislo2 += chislo
	fmt.Println("Число:",chislo2)
	return chislo2
}
func Perevod(chislo string, p16 map[string]string) {
	s := ""
	chislo2 := []rune(chislo)
	var slais04 = chislo2[:3]
	var slais08 = chislo2[:7]
	var slais12 = chislo2[:11]
	var slais16 = chislo2[:15]
	s += p16[string(slais04)]
	s += p16[string(slais08)]
	s += p16[string(slais12)]
	s += p16[string(slais16)]
	fmt.Println(s)
	
}
	

