package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `gorm:"primaryKey;default:UUID()" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
