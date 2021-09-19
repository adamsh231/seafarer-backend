package repositories

import (
	"seafarer-backend/api/storage/interfaces"
	"seafarer-backend/domain/models"
	"time"

	"gorm.io/gorm"
)

type FileRepository struct {
	Postgres *gorm.DB
}

func NewFileRepository(postgres *gorm.DB) interfaces.IFileRepository {
	return &FileRepository{Postgres: postgres}
}

func (repo FileRepository) IsExist(userID string, name string) (isExist bool, err error) {
	var count int64
	if err = repo.Postgres.Model(models.NewFile()).Where("user_id = ?", userID).Where("name = ?", name).Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return false, err
	}

	if count > int64(0) {
		return true, err
	}

	return false, err
}

func (repo FileRepository) Read(id string, file *models.File) (err error) {
	return repo.Postgres.Where("id = ?", id).Where("deleted_at IS NULL").First(file).Error
}

func (repo FileRepository) Browse(userID string, offset, limit int, search, orderBy, sort string) (files []models.File, count int64, err error) {

	query := repo.Postgres.Table(models.NewFile().TableName()).
		Where("user_id = ?", userID)
	query.Where(`name LIKE '%` + search + `%'`)
	query.Where("deleted_at IS NULL")

	totalQuery := query
	totalQuery.Count(&count)

	query.Order(orderBy + ` ` + sort)
	query.Offset(offset).Limit(limit)
	err = query.Scan(&files).Error
	return files, count, err
}

func (repo FileRepository) Add(file *models.File, tx *gorm.DB) (err error) {
	return tx.Omit("id", "deleted_at").Create(file).Error
}

func (repo FileRepository) Delete(id string, tx *gorm.DB) (err error) {
	return tx.Model(models.File{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}
