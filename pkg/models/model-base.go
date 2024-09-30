// Package models @Author hubo 2024/9/29 22:51:00
package models

import "time"

type BaseModel struct {
	ID        uint `gorm:"primary_key" json:"ID"` // 主键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index" json:"-"`
}
