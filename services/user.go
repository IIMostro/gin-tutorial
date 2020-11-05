package services

import (
	"fmt"
)

type Person interface {
	Eat() string

	Run() string
}

type Student struct {
	Name string

	Age int
}

func (s Student) Eat() string {
	result := fmt.Sprintf("user name: %s, age: %d, eating", s.Name, s.Age)
	return result
}

func (s Student) Run() string {
	result := fmt.Sprintf("user name: %s, age: %d, running", s.Name, s.Age)
	return result
}
