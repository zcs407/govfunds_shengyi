package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DBSQL *gorm.DB

func InitDB() {
	//Customize the datetime
	gorm.NowFunc = func() time.Time {
		return time.Now().Round(time.Second)
	}
	//dsn := "ander" + ":" + "Ander110!!" + "@tcp(" + "127.0.0.1:3316" + ")/" + "govfundsdb" + "?charset=utf8&parseTime=true&loc=Local"
	dsnOnline := "root" + ":" + "zhengfu2018" + "@tcp(" + "127.0.0.1:3306" + ")/" + "bdm26610270_db" + "?charset=utf8&parseTime=true&loc=Local"
	db, err := gorm.Open("mysql", dsnOnline)
	//db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.DB()
	db.LogMode(true)

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//设置最大空闲连接池
	db.DB().SetMaxIdleConns(10)
	//设置最大打开连接池
	db.DB().SetMaxOpenConns(20)
	//连接最大超时时间
	db.DB().SetConnMaxLifetime(time.Hour)
	db.SingularTable(true)
	DBSQL = db
}
