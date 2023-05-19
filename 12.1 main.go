package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	numbers, _ := os.Open("numbers.txt")
	defer numbers.Close()

	var counter ['9' - '0' + 1]uint
	n, in := 0, bufio.NewScanner(numbers)
	for in.Scan() {
		n++
		for _, ch := range in.Text() {
			if ch >= '0' && ch <= '9' {
				counter[ch-'0']++
			}
		}
	}

	fmt.Println("Всего ввели", n, "строк.")
	for i, cnt := range counter {
		if cnt > 0 {
			fmt.Printf("%c  %4d\n", '0'+i, cnt)
		}
	}
}
