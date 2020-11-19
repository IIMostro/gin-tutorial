package configuration

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

var pool *redis.Pool

func GetConnection() redis.Conn {
	connection := pool.Get()
	defer connection.Close()
	return connection
}

func getClient() (redis.Conn, error) {
	addr := fmt.Sprintf("%s:%d", Properties.Redis.Host, Properties.Redis.Port)
	log.Printf("redis client url: %s", addr)
	password := redis.DialPassword(Properties.Redis.Password)
	database := redis.DialDatabase(Properties.Redis.DB)
	return redis.Dial("tcp", addr, password, database)
}

func getPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     Properties.Redis.Pool.MaxIdle,
		IdleTimeout: Properties.Redis.Pool.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return getClient()
		},
	}
}

func init() {
	pool = getPool()
}
