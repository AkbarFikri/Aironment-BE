package model

import "time"

type UserCreateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required,alpha"`
}

type UserCreateResponse struct {
	ID string `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}
