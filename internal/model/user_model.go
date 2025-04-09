package model

import (
	"time"
)

type User struct {
	ID          string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Role        string    `json:"role" gorm:"type:varchar(50);"`
	User_Status bool      `json:"status" gorm:"column:status;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

type UserResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	User_Status bool      `json:"status"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserAdminsResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	User_Status bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
