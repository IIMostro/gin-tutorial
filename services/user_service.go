package services

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/streadway/amqp"
	"ilmostro.org/gin-tutorial/configuration"
	"ilmostro.org/gin-tutorial/repository"
	"log"
	"time"
)

type StudentService interface {
	GetStudents() []repository.Student

	GetStudent(id string) repository.Student

	Save(student repository.Student)
}

const (
	StudentsCacheKey = "students"
	StudentCacheKey  = "student:%d"
)

var template = configuration.GetConnection()
var channel = configuration.Channel

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
	student.CreateTime = time.Now().String()
	repository.Save(&student)
	cacheKey := fmt.Sprintf(StudentCacheKey, student.Id)
	cacheValue, _ := json.Marshal(student)
	_, err := template.Do("SET", cacheKey, cacheValue)
	if err != nil {
		log.Printf("set redis cache error, cause: %f", err)
	}
	_ = channel.Publish("go-tutorial-user",
		"insert",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        cacheValue,
		})
}

func getStudentsByCache() []repository.Student {
	value, err := redis.String(template.Do("GET", StudentsCacheKey))
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
	_, err = template.Do("SET", StudentsCacheKey, marshal)
	if err != nil {
		log.Printf("put redis cache err, cause: %f", err)
	}
}

func getStudentsByDB() []repository.Student {
	return repository.GetAllUserFromDB()
}
