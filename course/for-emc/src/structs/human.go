package main

import "fmt"

type Human struct {
	Name string
}

func (h *Human) eat(food string) {
	fmt.Println(h.Name, "is eating", food)
}
