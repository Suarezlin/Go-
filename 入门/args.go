package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	args := os.Args[1:] // os.Args 是一个 string 类型的切片
	for _, arg := range args {
		s += " " + arg
	}
	fmt.Println(s)
}
