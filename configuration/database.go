package configuration

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func GetConnection() *gorm.DB {

	connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Properties.Database.User, Properties.Database.Password,
		Properties.Database.Server, Properties.Database.Port, Properties.Database.DatabaseName)

	open, err := gorm.Open("mysql", connectionUrl)
	if err != nil {
		log.Fatalf("get database connection error, %f", err)
	}
	open.DB().SetMaxIdleConns(Properties.Database.Pool.MaxIdleConnection)
	open.DB().SetMaxOpenConns(Properties.Database.Pool.MaxConnection)
	if err != nil {
		log.Fatalf("get database connection error, %f", err)
	}
	return open
}
