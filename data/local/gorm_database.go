package local

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	gmsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
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
		Loc:                  time.Local,
	}

	db, err := gorm.Open(gmsql.Open(cfg.FormatDSN()), &gorm.Config{})

	if err != nil {
		fmt.Println("cannot connect to database ", "mysql")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("gorm is connected to the database ", "mysql")
	}

	err = db.AutoMigrate(
		entity.User{},
		entity.Message{},
		entity.CacheMessage{},
		entity.Energy{},
		entity.Setting{})

	if err != nil {
		log.Fatal("migration error:", err)
	}

	return db
}
