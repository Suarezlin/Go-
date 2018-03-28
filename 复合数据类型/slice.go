package main

import (
	"fmt"
)

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func traverse1(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func traverse2(s []string) {
	for i, e := range s {
		fmt.Printf("(%d %s)\n", i, e)
	}
}

func main() {
	var runes []rune
	for _, e := range "Hello, 世界" {
		runes = append(runes, e)
	}
	fmt.Println(runes)
}
