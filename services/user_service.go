package services

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"ilmostro.org/gin-tutorial/configuration"
	"ilmostro.org/gin-tutorial/repository"
)

var template = configuration.GetConnection()

func GetStudents() []repository.Student {

	var students []repository.Student
	students = getStudentsByCache()
	if students != nil && len(students) > 0 {
		return students
	}
	students = getStudentsByDB()
	setStudentCache(students)

	return students
}

func getStudentsByCache() []repository.Student {
	value, err := redis.String(template.Do("GET", "students"))
	var students []repository.Student
	if err != nil {
		return []repository.Student{}
	}
	err = json.Unmarshal([]byte(string(value)), &students)
	return students
}

func setStudentCache(students []repository.Student) {
	marshal, err := json.Marshal(students)
	if err != nil {
		return
	}
	_, _ = template.Do("SET", "students", marshal)
}

func getStudentsByDB() []repository.Student {
	return repository.GetAllUserFromDB()
}
