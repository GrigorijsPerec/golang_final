package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type student struct {
	lastname  string
	firstname string
	class     string
}

func main() {
	file, err := os.Open("list.csv")
	if err != nil {
		fmt.Println("Error while opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	students := make(map[string]int)

	for scanner.Scan() {
		tempStudent := strings.Split(scanner.Text(), ",")
		st := student{
			firstname: tempStudent[0],
			lastname:  tempStudent[1],
			class:     tempStudent[2],
		}
		students[st.lastname]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scanning the file:", err)
		return
	}

	duplicatesFound := false
	for lastname, count := range students {
		if count > 1 {
			fmt.Println(lastname, count)
			duplicatesFound = true
		}
	}

	if !duplicatesFound {
		fmt.Println("No")
	}
	for first := 0; first< len(students); first++{
		for second := first+1; second<len(students; second++){
			
		}
	}
}
