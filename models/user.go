package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `gorm:"unique" json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
