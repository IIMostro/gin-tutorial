package configuration

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetConnection() (*sql.DB, error) {

	properties := GetProperties().Database
	connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", properties.User, properties.Password,
		properties.Server, properties.Port, properties.DatabaseName)
	connection, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		log.Printf("get mysql connection error!, cause:#%v", err)
		return nil, err
	}

	connection.SetMaxOpenConns(properties.Pool.MaxConnection)
	connection.SetMaxIdleConns(properties.Pool.MaxIdleConnection)

	return connection, nil
}
