package repository

import (
	"fmt"
	"log"
	"time"
)

type Student struct {
	Model

	Name string `json:"name"`

	Age int `json:"age"`
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
		log.Printf("query user error!, cause:%v", err)
		return nil
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
		model := Model{Id: id, CreatedOn: time.Now().Nanosecond()}
		student := Student{
			Name: name,
			Age:  age,
		}
		student.Model = model
		students = append(students, student)
	}

	return students
}

func GetStudentById(Id string) Student {

	sql := fmt.Sprintf("select id, `name`, age from user where id = %s", Id)
	result, err := Connection.Query(sql)
	if err != nil {
		log.Printf("query user error!, cause:%v", err)
		return Student{}
	}

	var id, age int
	var name string
	for result.Next() {
		err = result.Scan(&id, &name, &age)
		if err != nil {
			return Student{}
		}
	}

	model := Model{Id: id, CreatedOn: time.Now().Nanosecond()}
	student := Student{
		Name: name,
		Age:  age,
	}
	student.Model = model

	return student
}
