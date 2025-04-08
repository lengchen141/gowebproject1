package models

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	Id           int `gorm:"primaryKey"`
	Username     string
	Password     string
	CreatedAtStr string `gorm:"-" json:"CreatedAt"`
	UpdatedAtStr string `gorm:"-" json:"UpdatedAt"`
}

func (table *UserInfo) TableName() string {
	return "t_user"
}

func (u *UserInfo) BeforeMarshalJSON() (err error) {
	u.CreatedAtStr = u.CreatedAt.In(time.UTC).Format("2006-01-02T15:04:05Z")
	u.UpdatedAtStr = u.UpdatedAt.In(time.UTC).Format("2006-01-02T15:04:05Z")
	return nil
}
