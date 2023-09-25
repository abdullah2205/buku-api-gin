package models

import (
	"time"
)

type Bukus struct {
    ID uint   `json:"id" gorm:"primary_key"`
    Judul string `json:"judul"`
    Tahun string `json:"tahun"`
    UserID uint `json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}