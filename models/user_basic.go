package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Passwd        string
	Phone         string
	Email         string
	Identify      string
	ClientIP      string
	ClientPort    string
	LoginTime     time.Time `gorm:"default:null"`
	HeartbeatTime time.Time `gorm:"default:null"`
	LogoutTime    time.Time `gorm:"default:null"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
