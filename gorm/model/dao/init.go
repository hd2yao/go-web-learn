package dao

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var _DB *gorm.DB

func DB() *gorm.DB {
	return _DB
}

func init() {
	_DB = initDB()
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_web_scout")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}
