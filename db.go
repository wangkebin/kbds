package main

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
)

func Connect(connstr string) (*gorm.DB, error) {
	var dbErr error
	//var db *gorm.DB

	once.Do(func() {
		db, dbErr = gorm.Open(
			mysql.Open(connstr),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Error),
			})

		if dbErr != nil {
			return
		}
		sqldb, err := db.DB()
		if err != nil {
			dbErr = err
			return
		}
		sqldb.SetMaxIdleConns(6)
		sqldb.SetMaxOpenConns(100)

	})

	if dbErr != nil {
		return nil, dbErr
	}
	return db, nil
}
