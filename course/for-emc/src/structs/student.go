package main

type Student struct {
	Human
	age int
	sex byte
}

func (s *Student) Hi() string {
	return "Hello"
}

func (s *Student) RenamePoint() {
	s.Name += "_new"
	return
}

func (s Student) Rename() {
	s.Name += "_new"
	return
}
