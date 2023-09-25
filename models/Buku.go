package models

import (
	"time"
)

type Bukus struct {
    ID uint64   `json:"id" gorm:"primary_key"`
    Judul string `json:"judul" binding:"required"`
    Tahun string `json:"tahun" binding:"required"`
    UserID uint `json:"user_id" binding:"required" gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}