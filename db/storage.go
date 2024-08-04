package db

import (
	"fmt"
	"golang-im/config"

	"gorm.io/gorm"
)

type Database interface {
	Connect() (*gorm.DB, error)
	GetDSN() string
}

// registry用来存储数据库类型及其构造函数
var registry = make(map[string]func(cfg *config.Config) Database)

// Registry用来注册新的数据库类型
func Registry(dbType string, factory func(cfg *config.Config) Database) {
	registry[dbType] = factory
}

// NewDatabase 用来创建数据库实例
func NewDatabase(cfg *config.Config) (Database, error) {
	if factory, ok := registry[cfg.DBType]; ok {
		return factory(cfg), nil
	}
	return nil, fmt.Errorf("unsupport database: %s", cfg.DBType)
}
