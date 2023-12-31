package models

import (
    "time"
)

type User struct {
    ID uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
    Email string `json:"email" gorm:"unique;not null"`
    Password string `json:"password" gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
