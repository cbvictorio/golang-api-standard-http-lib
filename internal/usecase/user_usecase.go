package usecase

import (
	"golang-api-standard-http-lib/internal/domain"
	repository "golang-api-standard-http-lib/internal/repository"
	"golang-api-standard-http-lib/pkg"
	"time"

	"github.com/google/uuid"
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

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{userRepository: r}
}

func (userService *UserService) GetByEmail(email string) (*domain.User, error) {
	result, err := userService.userRepository.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (userService *UserService) CreateUser(userInput *NewUserDTO) error {
	password, err := pkg.GenerateHashFromPassword(userInput.Password)

	if err != nil {
		return err
	}

	user := &domain.User{
		ID:        uuid.NewString(),
		Name:      userInput.Name,
		Email:     userInput.Email,
		Role:      domain.RoleCustomer,
		CreatedAt: time.Now(),
		Password:  password,
	}

	result := userService.userRepository.Create(user)

	return result
}
