package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GormMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(db:3306)/miniproject2?charset=utf8mb4&parseTime=True&loc=UTC"), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
		return nil, err
	}
	return db, nil
}
