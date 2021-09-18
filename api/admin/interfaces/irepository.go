package interfaces

import "seafarer-backend/domain/models"

type IAdminRepository interface {
	ReadByEmail(email string, user *models.Admin) (err error)
}
