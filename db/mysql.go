package db

import (
	"fmt"
	"golang-im/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化MySQL数据库
type MySQL struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (m *MySQL) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.DBName)
}

func (m *MySQL) Connect() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(m.GetDSN()), &gorm.Config{})
}

func init() {
	Registry("mysql", func(cfg *config.Config) Database {
		return &MySQL{
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			DBName:   cfg.DBName,
		}
	})
}
