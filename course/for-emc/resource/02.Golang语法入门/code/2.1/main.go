package main

import "fmt"

type Abc struct {
	A int
	B int
	C int
}

func main() {
	a := 1
	b := make([]int, 0)
	var s Abc
	s.A = 1
	s.B = 2
	fmt.Printf("%d %d, %d", len(b), cap(b), s.A)
	fmt.Println("Hello World", a)

	stdZhang := &Student{
		name: "张三",
		age:  18,
		sex:  1,
	}

	fmt.Printf("name is %s", stdZhang.name)

	stdLi := new(Student)
	stdLi.name = "李四"
	fmt.Printf("name is %s", stdLi.name)

	stdWang := Student{
		name: "王二",
	}
	fmt.Printf("name is %s", stdWang.name)
}

func changeName(s Student) {
	s.name = "nameName"
}
