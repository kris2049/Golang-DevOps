package main

import (
	"golang-im/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("zjz:123@tcp(127.0.0.1:3306)/IM"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{
		Name: "zjz",
	}
	db.Create(user)

	// Read
	res := db.First(user, 1) // 根据整型主键查找
	println(res)
	res = db.First(user, "name = ?", "zjz") // 查找 code 字段值为 D42 的记录
	println(res)

	// Update - 将 user 的 name 更新为 zhou
	db.Model(user).Update("Passwd", "1234")

	// db.Model(user).Update("name", "zhou")
	// Update - 更新多个字段

	// Delete - 删除 user
	// db.Delete(&user, 1)
}
