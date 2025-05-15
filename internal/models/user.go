package models

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name" form:"name" binding:"required"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" form:"email" binding:"required"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password" form:"password" binding:"required"`
	PhoneNo   string    `gorm:"type:varchar(15);not null" json:"phone_no" form:"phone_no" binding:"required"`
	ImageName string    `gorm:"type:varchar(255)" json:"image_name" form:"image_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
