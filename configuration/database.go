package configuration

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Connection *gorm.DB

func getConnection() (*gorm.DB, error) {

	connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Properties.Database.User, Properties.Database.Password,
		Properties.Database.Server, Properties.Database.Port, Properties.Database.DatabaseName)

	log.Printf("get connection url: %s", connectionUrl)

	open, err := gorm.Open("mysql", connectionUrl)
	if err != nil {
		return nil, err
	}
	open.DB().SetMaxIdleConns(Properties.Database.Pool.MaxIdleConnection)
	open.DB().SetMaxOpenConns(Properties.Database.Pool.MaxConnection)

	return open, nil
}

func init() {
	conn, err := getConnection()
	if err != nil {
		log.Fatalf("get database connection error, %f", err)
	}

	Connection = conn
}
