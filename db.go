package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func initializeConnection(cfg *config) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.dbUser, cfg.dbPassword, cfg.dbHost, cfg.dbName,
	)
	dbConn, err := gorm.Open(
		mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		panic(err)
	}
	db = dbConn
	if err != nil {
		//_ = Log(errors.New("Fatal error, main(): database connection cannot be established." + err.Error()))
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetConnMaxLifetime(time.Minute)
}

func Conn(cfg *config) *gorm.DB {
	if db == nil {
		initializeConnection(cfg)
	}
	return db
}
