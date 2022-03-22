package main

import (
	"fmt"
	"github.com/stack-labs/stack/broker"
	"github.com/stack-labs/stack/plugin/registry/consul"
)

func main() {
	stdZhang := &Student{
		Human: Human{Name: "张三"},
		age:   18,
		sex:   1,
	}

	stdZhang.eat("苹果")
	stdZhang.Name = "张老三"
	stdZhang.eat("苹果")

	fmt.Printf("name is %s\n", stdZhang.Name)

	stdLi := new(Student)
	stdLi.Name = "李四"
	stdLi.Rename()
	fmt.Printf("after rename, name is %s\n", stdLi.Name)
	stdLi.RenamePoint()
	fmt.Printf("after rename with point, name is %s\n", stdLi.Name)

	fmt.Printf("name is %s\n", stdLi.Name)
	changePoint(stdLi)
	fmt.Printf("after change, name is %s\n", stdLi.Name)

	stdWang := Student{}
	stdWang.Name = "王二"
	fmt.Printf("name is %s\n", stdWang.Name)
	changeName(stdWang)
	fmt.Printf("after change, name is %s\n", stdWang.Name)

	stdLi.Rename()
	fmt.Printf("after new, name is %s\n", stdLi.Name)
	stdLi.RenamePoint()
	fmt.Printf("after new with point, name is %s\n", stdLi.Name)
}

func changeName(s Student) {
	s.Name = "王老二"
}

func changePoint(s *Student) {
	s.Name = "李老四"
}
