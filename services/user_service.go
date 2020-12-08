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

	GetStudent(id string) *repository.Student

	Save(student repository.Student)
}

const (
	StudentsCacheKey = "students"
	StudentCacheKey  = "student:%d"
)

type RedisUserService struct {
	connection redis.Conn
}

func NewRedisUserService() *RedisUserService {
	return &RedisUserService{connection: configuration.GetRedisClient()}
}

func (r RedisUserService) GetStudents() []repository.Student {
	var students []repository.Student
	students = getStudentsByCache(r.connection)
	if students != nil && len(students) > 0 {
		return students
	}
	students = getStudentsByDB()
	setStudentCache(students, r.connection)
	return students
}

func (r RedisUserService) GetStudent(id string) *repository.Student {
	var student *repository.Student
	student = getStudentCache(id, r.connection)
	if student != nil {
		return student
	}
	student = repository.GetStudentById(id)
	errors := make(chan error)
	setStudent(student, r.connection, errors)
	r.connection.Close()
	return student
}

func (r RedisUserService) Save(student repository.Student) {
	student.CreateTime = time.Now().String()
	repository.Save(&student)
	cacheValue, _ := json.Marshal(student)

	asyncErrorChan := make(chan error)

	go setStudent(&student, r.connection, asyncErrorChan)
	go sendStudentMessage(cacheValue, asyncErrorChan)

	errors := <-asyncErrorChan
	if errors != nil {
		errorMsg := fmt.Errorf("async save user error, cause:%w", errors)
		panic(errorMsg)
	}
}

func getStudentCache(id string, connection redis.Conn) *repository.Student {
	key := fmt.Sprintf("student:%s", id)
	value, _ := redis.String(connection.Do("GET", key))
	var student *repository.Student
	if &value == nil || value == "" {
		return student
	}
	_ = json.Unmarshal([]byte(value), &student)
	return student
}

func setStudent(s *repository.Student, connection redis.Conn, asyncErrorChan chan error) {
	key := fmt.Sprintf(StudentCacheKey, s.Id)
	marshal, err := json.Marshal(s)
	if err != nil {
		asyncErrorChan <- err
		err := fmt.Errorf("set redis cache error!, caluse: %w", err)
		panic(err)
	}
	_, err = connection.Do("SET", key, marshal)
	asyncErrorChan <- err
}

func getStudentsByCache(connection redis.Conn) []repository.Student {
	value, err := redis.String(connection.Do("GET", StudentsCacheKey))
	var students []repository.Student
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(value), &students)
	return students
}

func setStudentCache(students []repository.Student, connection redis.Conn) {
	marshal, err := json.Marshal(students)
	if err != nil {
		return
	}
	_, err = connection.Do("SET", StudentsCacheKey, marshal)
	if err != nil {
		log.Printf("put redis cache err, cause: %f", err)
	}
}

func getStudentsByDB() []repository.Student {
	return repository.GetAllUserFromDB()
}

func sendStudentMessage(cacheValue []byte, asyncErrorChan chan error) {
	channel := configuration.GetRabbitConnection()
	defer channel.Close()
	err := channel.Publish("go-tutorial-user",
		"insert",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        cacheValue,
		})
	asyncErrorChan <- err
}
