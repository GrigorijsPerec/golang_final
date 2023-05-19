package main

import (
	"fmt"
	"strings"
)

func main() {

	task2()
}

func task0() {
	str := "dfj nhfjfjk dgj  g df jsgp[h sfjho gjm'hgj ipdhgjk 'ghkjl ghp[]] "
	count := len(strings.Fields(str))
	fmt.Println(count)
}

func task1() {
	fmt.Println(strings.Trim(" ffderfwj  qefweuuu   eh rfwehnf uj jefjwej ama", "amam "))
}

func task2() {
	s := []string{"a", "s"}

	for i := 0; i < len(s); i++ {
		str := strings.ReplaceAll("sdfsadfasdf sdf ds a gf", s[i], "")
		fmt.Println(str)
	}
}

func task3() {

}
