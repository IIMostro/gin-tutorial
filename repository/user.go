package repository

import (
	"fmt"
	"log"
)

type Person interface {
	Eat() string

	Run() string
}

type Student struct {
	Model

	Id int

	Name string

	Age int
}

func (s Student) Eat() string {
	result := fmt.Sprintf("user id:%d, name: %s, age: %d, eating", s.Id, s.Name, s.Age)
	return result
}

func (s Student) Run() string {
	result := fmt.Sprintf("user id:%d, name: %s, age: %d, running", s.Id, s.Name, s.Age)
	return result
}

func GetAllUserFromDB() []Student {

	sql := "select id, `name`, age from user"
	result, err := Connection.Query(sql)
	if err != nil {
		log.Fatalf("query user error!, cause:%v", err)
	}

	var students []Student

	for result.Next() {
		var id, age int
		var name string

		err := result.Scan(&id, &name, &age)
		if err != nil {
			log.Printf("result scan rows error!, cause: %v", err)
			continue
		}
		student := Student{
			Id:   id,
			Name: name,
			Age:  age,
		}
		students = append(students, student)
	}

	return students
}
