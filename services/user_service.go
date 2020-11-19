package services

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"ilmostro.org/gin-tutorial/configuration"
	"ilmostro.org/gin-tutorial/repository"
)

type StudentService interface {
	GetStudents() []repository.Student

	GetStudent(id string) repository.Student

	Save(student repository.Student)
}

const (
	StudentCacheKey = "students"
)

var template = configuration.GetConnection()

type RedisUserService struct {
}

func (r RedisUserService) GetStudents() []repository.Student {

	var students []repository.Student
	students = getStudentsByCache()
	if students != nil && len(students) > 0 {
		return students
	}
	students = getStudentsByDB()
	setStudentCache(students)

	return students
}

func (r RedisUserService) GetStudent(id string) repository.Student {
	return repository.Student{}
}

func (r RedisUserService) Save(student repository.Student) {

}

func getStudentsByCache() []repository.Student {
	value, err := redis.String(template.Do("GET", StudentCacheKey))
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
	_, _ = template.Do("SET", StudentCacheKey, marshal)
}

func getStudentsByDB() []repository.Student {
	return repository.GetAllUserFromDB()
}
