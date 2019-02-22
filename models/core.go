package models

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var db *gorm.DB
var redisdb redis.Conn

func init() {
	var e error
	if e = os.MkdirAll("data", 0777); e != nil {
		panic("创建db目录失败" + e.Error())
	}
	db, e = gorm.Open("sqlite3", "data/data.db")
	if e != nil {
		beego.Info("打开数据库失败", e)
		panic("创建数据库失败" + e.Error())
	}
	db.AutoMigrate(&User{}, &Note{}, &Comment{}, &Praise{})
	//如果数据库为空，新增一条管理员账号
	var count int
	e = db.Model(&User{}).Count(&count).Error
	beego.Info(e)
	if e == nil && count == 0 {
		db.Create(&User{
			Name:   "admin",
			Email:  "admin@163.com",
			Pwd:    "123456",
			Avatar: "static/images/info-img.png",
			Role:   0,
		})
	}

}
func RedisDb() redis.Conn {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	return conn
}
