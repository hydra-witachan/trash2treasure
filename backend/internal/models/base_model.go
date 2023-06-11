package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `gorm:"column:"id";primaryKey;default:UUID()" json:"id"`
	CreatedAt time.Time      `gorm:"column:"created_at";autoCreateTime;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:"updated_at";autoUpdateTime;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:"deleted_at";index" json:"deletedAt"`
}
