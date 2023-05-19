package main

import "fmt"

func main() {
	name := "  SDFFGJHFGJ DFGSDP F] ]JK GHJHH  "
	runes := []rune(name)
	indexFrom := 0
	indexTo := len(runes)

	for _, c := range runes {
		if c == 32 {
			indexFrom++
		} else {
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == 32 {
			indexTo--
		} else {
			break
		}
	}

	nonSpace := runes[indexFrom:indexTo]
	fmt.Printf("%c", nonSpace)
}
