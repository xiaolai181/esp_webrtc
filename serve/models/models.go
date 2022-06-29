package models

import (
	"fmt"
	"log"

	setting "esp_webrtc/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var (
		err                          error
		dbName, user, password, host string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	db.AutoMigrate(&Esp_User{})
	var count int64
	// Model(&User{})查询用户表, Count(&count) 将用户表的数据赋值给count字段。
	if err := db.Model(&Esp_User{}).Count(&count).Error; err == nil && count == 0 {
		log.Println("User 表不存在，创建...")
		//新增
		db.Create(&Esp_User{
			Name: "admin",
			//邮箱
			Email: "admin@qq.com",
			//密码
			Password: "123123",
			//角色 管理员
			Role: 0,
		})
	}

}

func DBCreate(value interface{}) {
	db.Create(value)
}

func GetDB() *gorm.DB {
	return db
}
