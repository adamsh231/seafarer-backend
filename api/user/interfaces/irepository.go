package interfaces

import (
	"seafarer-backend/domain/models"
)

type IUserRepository interface {
	ReadByEmail(email string, user *models.User) (err error)
}
