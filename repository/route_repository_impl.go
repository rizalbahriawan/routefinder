package repository

import (
	"routefinder_golang/entity"

	"gorm.io/gorm"
)

type RouteRepositoryImpl struct {
	DB *gorm.DB
}

func NewRouteRepository(db *gorm.DB) *RouteRepositoryImpl {
	return &RouteRepositoryImpl{DB: db}
}

func (repository *RouteRepositoryImpl) GetAllRoute() ([]entity.Route, error) {
	var result = []entity.Route{}
	err := repository.DB.Order("id asc").Find(&result).Error
	return result, err
}

func (repository *RouteRepositoryImpl) FindRoute(id int64) (entity.Route, error) {
	var result entity.Route
	err := repository.DB.First(&result, id).Error
	return result, err
}

func (repository *RouteRepositoryImpl) CreateRoute(input entity.Route) error {
	result := repository.DB.Create(&input)

	return result.Error
}

func (repository *RouteRepositoryImpl) EditRoute(input entity.Route) error {
	result := repository.DB.Save(input)
	return result.Error
}

func (repository *RouteRepositoryImpl) DeleteRoute(id int64) error {
	var obj entity.Route
	result := repository.DB.Where("id = ? ", id).Delete(obj).Error
	return result
}
