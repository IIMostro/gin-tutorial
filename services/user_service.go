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
	var student repository.Student
	student = getStudentCache(id)
	if student != (repository.Student{}) {
		return student
	}
	student = repository.GetStudentById(id)
	setStudent(student)
	return student
}

func getStudentCache(id string) repository.Student {
	key := fmt.Sprintf("student:%s", id)
	client := configuration.GetRedisClient()
	defer client.Close()
	value, _ := redis.String(client.Do("GET", key))

	var student repository.Student
	_ = json.Unmarshal([]byte(value), &student)
	return student
}

func setStudent(s repository.Student) {
	key := fmt.Sprintf("student:%d", s.Id)
	client := configuration.GetRedisClient()
	defer client.Close()

	marshal, err := json.Marshal(s)
	if err != nil {
		log.Printf("set redis cache error!, caluse: %f", err)
	}
	_, _ = client.Do("SET", key, marshal)
}

func (r RedisUserService) Save(student repository.Student) {
	template := configuration.GetRedisClient()
	channel := configuration.GetRabbitConnection()
	defer template.Close()
	defer channel.Close()
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
	template := configuration.GetRedisClient()
	defer template.Close()
	value, err := redis.String(template.Do("GET", StudentsCacheKey))
	var students []repository.Student
	if err != nil {
		return []repository.Student{}
	}
	err = json.Unmarshal([]byte(value), &students)
	return students
}

func setStudentCache(students []repository.Student) {
	template := configuration.GetRedisClient()
	defer template.Close()
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
