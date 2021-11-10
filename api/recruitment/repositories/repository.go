package repositories

import (
	"seafarer-backend/api/recruitment/interfaces"
	"seafarer-backend/domain/models"

	"gorm.io/gorm"
)

type RecruitmentsRepository struct {
	Postgres *gorm.DB
}

func NewRecruitmentsRepository(postgres *gorm.DB) interfaces.IRecruitmentsRepository {
	return &RecruitmentsRepository{Postgres: postgres}
}

func (repo RecruitmentsRepository) Add(model *models.Recruitments, tx *gorm.DB) (err error) {
	return tx.Omit("deleted_at", "available_on", "sign_on").Create(&model).Error
}

func (repo RecruitmentsRepository) UpdateByIDUser(idUser string, model models.Recruitments, tx *gorm.DB) (err error) {
	return tx.Model(models.NewRecruitments()).Where("user_id = ?", idUser).Updates(&model).Error
}

func (repo RecruitmentsRepository) FilterByStatusRecruitment(offset, limit int, orderBy, sort, search, status string) (model []models.RecruitmentsDetail, count int64, err error) {
	queryBuilder := repo.Postgres.Model(&models.User{})
	queryBuilder.Select("users.id, users.name, users.email, users.is_verified, users.updated_at, users.deleted_at, users.company_id, recruitments.user_id, recruitments.created_at, recruitments.updated_at, recruitments.status, recruitments.salary, recruitments.expect_salary , recruitments.position, recruitments.available_on, recruitments.sign_on, recruitments.letter, recruitments.ship")
	queryBuilder.Joins("JOIN recruitments ON recruitments.id = users.recruitment_id AND recruitments.status=?", status)
	queryBuilder.Where("users.deleted_at IS NULL").
		Where("users.is_verified = TRUE")

	if search != "" {
		queryBuilder.Where("users.name LIKE '%" + search + "%' OR users.email LIKE '%" + search + "%'")
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
