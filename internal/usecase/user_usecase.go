package usecase

import (
	"errors"
	"golang-api-standard-http-lib/internal/domain"
	"golang-api-standard-http-lib/internal/repository"
	"golang-api-standard-http-lib/pkg"
	"time"

	"github.com/google/uuid"
)

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{userRepository: r}
}

func (userService *UserService) DoesEmailExist(email string) (bool, error) {
	result, err := userService.userRepository.GetByEmail(email)

	if err != nil {
		return false, err
	}

	if result != nil {
		return true, nil
	}

	return false, nil
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

func (userService *UserService) AuthenticateUser(email string, password string) (*domain.User, error) {
	user, err := userService.userRepository.GetByEmail(email)

	// something went wrong while retrieving the user
	if err != nil {
		return nil, err
	}

	// user does not exist
	if user == nil {
		return nil, nil
	}

	// validating password flow
	if !pkg.CompareHashWithPassword(user.Password, password) {
		return nil, errors.New("Invalid username or password")
	}

	return user, nil
}
