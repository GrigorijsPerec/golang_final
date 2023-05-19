package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Открыть файл для чтения
	file, err := os.Open("numbers.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Создать срезы для хранения положительных и отрицательных чисел
	var pos, neg []int

	// Считывать числа из файла и добавлять их в соответствующий срез
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		if n >= 0 {
			pos = append(pos, n)
		} else {
			neg = append(neg, n)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Открыть файл для записи
	outfile, err := os.Create("output.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()

	// Записать положительные числа, затем отрицательные числа
	for _, n := range pos {
		fmt.Fprintln(outfile, n)
	}
	for _, n := range neg {
		fmt.Fprintln(outfile, n)
	}
}
