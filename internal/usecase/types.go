package usecase

import (
	repository "golang-api-standard-http-lib/internal/repository"
)

type UserCredentialsDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type NewUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserService struct {
	userRepository *repository.UserRepository
}
