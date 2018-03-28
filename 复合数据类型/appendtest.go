package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 3)
	for i := 1; i <= 20; i++ {
		s = append(s, i)
		fmt.Printf("%d %d\n", len(s), cap(s))
	}
}
