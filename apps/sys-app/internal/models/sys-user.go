// Package models @Author hubo 2024/9/29 23:18:00
package models

import "github.com/Hakunari/supertools-go/pkg/models"

type SysUser struct {
	models.BaseModel
	Username string `json:"username" gorm:"index;comment:用户名"`
}
