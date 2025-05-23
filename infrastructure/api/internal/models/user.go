package models

import "time"

type User struct {
    ID           int       `json:"id" db:"id"`
    Username     string    `json:"username" db:"username"`
    Email        string    `json:"email" db:"email"`
    PasswordHash string    `json:"password_hash" db:"password_hash"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateUserRequest struct {
    Username     string `json:"username" binding:"required"`
    Email        string `json:"email" binding:"required,email"`
    PasswordHash string `json:"password_hash" binding:"required"`
} 