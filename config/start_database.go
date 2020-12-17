package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewDatabase(conf *Config) *sql.DB {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.DB_USER, conf.DB_PASS, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
	db, err := sql.Open("mysql", source)
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	log.Println("Database is starting ...")
	return db
}
