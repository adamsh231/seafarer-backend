package repositories

import (
	"seafarer-backend/api/user/interfaces"
	"seafarer-backend/domain/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Postgres *gorm.DB
}

func NewUserRepository(postgres *gorm.DB) interfaces.IUserRepository {
	return &UserRepository{Postgres: postgres}
}

func (repo UserRepository) Add(model *models.User, tx *gorm.DB) (err error) {
	return tx.Omit("deleted_at").Create(model).Error
}

func (repo UserRepository) ReadByEmail(email string, model *models.User) (err error) {
	return repo.Postgres.Where("email = ?", email).First(model).Error
}

func (repo UserRepository) UpdateVerifiedByEmail(email string, tx *gorm.DB) (err error) {
	return tx.Model(models.NewUser()).Where("email = ?", email).Update("is_verified", true).Error
}

func (repo UserRepository) UpdatePasswordByEmail(email string, password string, tx *gorm.DB) (err error) {
	return tx.Model(models.NewUser()).Where("email = ?", email).Update("password", password).Error
}

func (repo UserRepository) Filter(offset, limit int, orderBy, sort, search string) (model []models.User, count int64, err error) {
	var modelUsers = models.NewUser()

	queryBuilder := repo.Postgres.Model(&modelUsers)
	queryBuilder.Where("deleted_at IS NULL").
		Where("is_verified = TRUE")

	if search != "" {
		queryBuilder.Where("name LIKE '%" + search + "%' OR email LIKE '%" + search + "%'")
	}

	countQuery := queryBuilder

	queryBuilder.Order(orderBy + ` ` + sort)
	queryBuilder.Offset(offset).Limit(limit)
	err = queryBuilder.Scan(&model).Error

	if err != nil {
		return model, count, err
	}

	// hitung total data
	countQuery.Offset(-1).Limit(-1).Count(&count)
	return model, count, err
}
