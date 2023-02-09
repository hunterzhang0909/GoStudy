package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primaryKey comment '主键'"`
	CreatedAt time.Time      `gorm:"type:datetime(3) comment '创建时间'"`
	UpdatedAt time.Time      `gorm:"type:datetime(3) comment '更新时间'"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:datetime(3)  comment '删除时间'"`
}
