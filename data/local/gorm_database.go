package local

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"vita-message-service/data/entity"
)

func GetGormDb() *gorm.DB {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBURL") + ":" + os.Getenv("DBPORT"),
		DBName:               os.Getenv("DBNAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := gorm.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("cannot connect to database ", "mysql")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("gorm is connected to the database ", "mysql")
	}

	db.AutoMigrate(&entity.User{})

	return db
}