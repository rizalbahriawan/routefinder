package repository

import (
	"routefinder_golang/entity"
	"routefinder_golang/request"

	"gorm.io/gorm"
)

type MountainRepositoryImpl struct {
	DB *gorm.DB
}

func NewMountainRepository(db *gorm.DB) *MountainRepositoryImpl {
	return &MountainRepositoryImpl{DB: db}
}

func (repository *MountainRepositoryImpl) GetAllMountain() ([]entity.Mountain, error) {
	var result = []entity.Mountain{}
	err := repository.DB.Preload("Route").Order("id asc").Find(&result).Error
	return result, err
}

func (repository *MountainRepositoryImpl) FindMountain(id int64) (entity.Mountain, error) {
	var result entity.Mountain
	err := repository.DB.Preload("Route").First(&result, id).Error
	return result, err
}

func (repository *MountainRepositoryImpl) CreateMountain(input request.MountainRequest) error {
	result := repository.DB.Create(&input)

	return result.Error
}

func (repository *MountainRepositoryImpl) EditMountain(input request.MountainRequest) error {
	result := repository.DB.Save(input)
	return result.Error
}

func (repository *MountainRepositoryImpl) DeleteMountain(id int64) error {
	var mountain request.MountainRequest
	result := repository.DB.Where("id = ? ", id).Delete(mountain).Error
	return result
}
