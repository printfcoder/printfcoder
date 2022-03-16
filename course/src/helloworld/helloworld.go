package main

import "fmt"

type Abc struct {
	A int
	B int
	C int
}

func main() {
	b := make([]int, 0)
	c := make(map[int]int, 0)
	var s Abc
	s.A = 1
	s.B = 2
	fmt.Printf("%d %d %d\n", len(b), cap(b), len(c))
	fmt.Println("Hello World")
}
