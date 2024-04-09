package initialize

import (
	User "com.dxj/app/model"
	"com.dxj/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitDatabase(config *config.Config) *gorm.DB {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	var dns string
	dns = config.Mysql.Username + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Url + ")/" + config.Mysql.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//打印SQL
		Logger: ormLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("==========数据库连接成功============！")
	//数据库迁移
	db.AutoMigrate(&User.User{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	//fmt.Println("===db的值=====", db)
	return db
}
