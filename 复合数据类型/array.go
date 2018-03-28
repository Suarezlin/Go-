package main

import (
	"fmt"
)

func main() {
	var a [3]int = [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
