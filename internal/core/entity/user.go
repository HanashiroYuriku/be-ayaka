package entity

import "time"

type User struct {
	ID string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}