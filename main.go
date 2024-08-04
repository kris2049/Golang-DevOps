package main

import (
	"golang-im/config"
	"golang-im/db"
	"golang-im/models"
	"golang-im/router"

	"gorm.io/gorm"
)

func main() {
	cfg := config.InitConfig()

	// 创建数据库实例
	database, err := db.NewDatabase(cfg)
	if err != nil {
		panic("failed to create database instance: " + err.Error())
	}

	// 连接数据库
	var db *gorm.DB
	db, err = database.Connect()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

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

	r := router.Router()

	r.Run(":8080")
}
