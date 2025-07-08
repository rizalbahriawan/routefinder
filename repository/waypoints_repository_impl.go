package repository

import (
	"routefinder_golang/entity"

	"gorm.io/gorm"
)

type WaypointsRepositoryImpl struct {
	DB *gorm.DB
}

func NewWaypointsRepository(db *gorm.DB) *WaypointsRepositoryImpl {
	return &WaypointsRepositoryImpl{DB: db}
}

func (repository *WaypointsRepositoryImpl) GetAllWaypoints() ([]entity.Waypoints, error) {
	var result = []entity.Waypoints{}
	err := repository.DB.Order("id asc").Find(&result).Error
	return result, err
}

func (repository *WaypointsRepositoryImpl) FindWaypoints(id int64) (entity.Waypoints, error) {
	var result entity.Waypoints
	err := repository.DB.First(&result, id).Error
	return result, err
}

func (repository *WaypointsRepositoryImpl) CreateWaypoints(input entity.Waypoints) error {
	result := repository.DB.Create(&input)

	return result.Error
}

func (repository *WaypointsRepositoryImpl) EditWaypoints(input entity.Waypoints) error {
	result := repository.DB.Save(input)
	return result.Error
}

func (repository *WaypointsRepositoryImpl) DeleteWaypoints(id int64) error {
	var obj entity.Waypoints
	result := repository.DB.Where("id = ? ", id).Delete(obj).Error
	return result
}
