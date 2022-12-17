package util

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB接続
func DbConnect() (*gorm.DB, error) {
	user := os.Getenv("FOOTBALL_DB_USER")
	pw := os.Getenv("FOOTBALL_DB_PASSWORD")
	db_name := os.Getenv("FOOTBALL_DB_NAME")
	var dsn string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, pw, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(dsn + "failed to connect database")
		return nil, errors.WithStack(ERR_USER_SYSTEM_ERROR)
	}
	return db, nil
}
