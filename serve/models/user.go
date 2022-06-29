package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Esp_User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Role     int    `gorm:"default:1"`
}

func Create_user(name string, password string, email string) bool {
	result := db.Create(&Esp_User{
		Name:     name,
		Password: password,
		Email:    email,
	})
	if result != nil {
		fmt.Println("插入错误")
		return false
	} else {
		fmt.Println("插入正确")
		return true
	}
}
func GetUserById(id int) (user Esp_User) {
	db.Where("id = ?", id).First(&user)
	fmt.Println(user.ID)
	return
}

func GetUserByName(name string) (user Esp_User) {
	db.Where("name = ?", name).First(&user)
	fmt.Println(user.ID)
	return
}

func Vaild_User(name string, password string) bool {
	var user Esp_User
	db.Where("name = ? and password = ?", name, password).First(&user)
	return user.ID > 0
}

func Deleted_user(id int) bool {
	result := db.Delete(&Esp_User{}, "id = ?", id)
	if result != nil {
		fmt.Println("删除错误")
		return false
	} else {
		fmt.Println("删除正确")
		return true
	}
}
