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
func GetUser(id int) (user Esp_User) {
	db.Where("id = ?", id).First(&user)
	fmt.Println(user.ID)
	return
}
