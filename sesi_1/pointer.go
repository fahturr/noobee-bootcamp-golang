package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s Student) CallMyName() {
	fmt.Printf("Hello, My Name is %s", s.Name)
}

func main() {
	var student1 = Student{
		Name:  "Diki",
		Class: "IX",
	}

	student1.SetMyName("Hafizh")
	student1.CallMyName()
}
