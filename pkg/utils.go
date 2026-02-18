package pkg

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type ErrorObject struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

const (
	COST_FACTOR = 14
)

func GenerateHashFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), COST_FACTOR)

	if err != nil {
		return "", errors.New("could not generate hash")
	}

	return string(bytes), nil
}

func CompareHashWithPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HttpAppErrorResponse(message string) *ErrorObject {
	return &ErrorObject{Error: true, Message: message}
}

func ErrorMessage(err error, fallback string) string {
	if err != nil {
		return err.Error()
	}
	return fallback
}
