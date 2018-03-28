package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct {
	X, Y int
}

func main() {
	// var dilbert Employee
	p := Point{1, 2}
	q := Point{1, 2}
	fmt.Println(p == q)
}
