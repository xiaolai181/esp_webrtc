package models

import "fmt"

func Createtable(table string, maps interface{}) bool {
	result := db.Table(table).Create(maps)
	if result != nil {
		fmt.Println("插入错误")
		return false
	} else {
		fmt.Println("插入正确")
		return true
	}
}
