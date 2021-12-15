package models

import (
	"gorm.io/gorm"
	"time"
)

type BizModel struct {
	CreatedAt time.Time      `gorm:"column:ft_create_time"`
	UpdatedAt time.Time      `gorm:"column:ft_update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:ft_delete_time"`
}
