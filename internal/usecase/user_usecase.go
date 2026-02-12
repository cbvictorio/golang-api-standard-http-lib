package usecase

import (
	"golang-api-standard-http-lib/internal/domain"
	repository "golang-api-standard-http-lib/internal/repository"
)

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
