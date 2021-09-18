package helpers

import (
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"seafarer-backend/domain/constants/messages"
)

type HashHelper struct{}

func NewHashHelper() HashHelper {
	return HashHelper{}
}

func (helper HashHelper) HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(hash), err
}

func (helper HashHelper) CheckHashString(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}

func (helper HashHelper) GenerateOTP(count int) (otp string, err error) {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, count)

	n, err := io.ReadAtLeast(rand.Reader, b, count)
	if err != nil {
		return otp, err
	}
	if n != count {
		return otp, errors.New(messages.BytesReadErrorMessage)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b), err
}
