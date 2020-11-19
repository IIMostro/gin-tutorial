package repository

import (
	"fmt"
	"ilmostro.org/gin-tutorial/configuration"
)

var connection = configuration.Connection

type Student struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`

	Name string `json:"name"`

	Age int `json:"age"`

	CreateTime string `json:"create_time" gorm:"now()"`
}

func (Student) TableName() string {
	return "user"
}

func (s Student) Eat() string {
	result := fmt.Sprintf("user id:%d, name: %s, age: %d, eating", s.Id, s.Name, s.Age)
	return result
}

func (s Student) Run() string {
	result := fmt.Sprintf("user id:%d, name: %s, age: %d, running", s.Id, s.Name, s.Age)
	return result
}

func init() {
	table := connection.HasTable("user")
	if !table {
		connection.CreateTable(&Student{})
		connection.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Student{})
	}
}

func GetAllUserFromDB() []Student {
	var users []Student
	connection.Select("id, name, age, create_time").Find(&users)
	return users
}

func GetStudentById(Id string) Student {
	student := Student{}
	connection.Where("id = ?", Id).First(&student)
	return student
}

func Save(student Student) {
	connection.Create(&student)
}
