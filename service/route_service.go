package service

import (
	"routefinder_golang/entity"
	"routefinder_golang/repository"
)

type RouteService interface {
	GetAllRoute() ([]entity.Route, error)
	FindRoute(id int64) (entity.Route, error)
	CreateRoute(input entity.Route) error
	EditRoute(id int64, input entity.Route) error
	DeleteRoute(id int64) error
}

type routeService struct {
	routeRepository repository.RouteRepository
}

func NewRouteService(repository repository.RouteRepository) *routeService {
	return &routeService{repository}
}

func (service *routeService) GetAllRoute() ([]entity.Route, error) {
	result, err := service.routeRepository.GetAllRoute()
	return result, err
}

func (service *routeService) FindRoute(id int64) (entity.Route, error) {
	result, err := service.routeRepository.FindRoute(id)
	return result, err
}

func (service *routeService) CreateRoute(input entity.Route) error {
	err := service.routeRepository.CreateRoute(input)
	return err
}

func (service *routeService) EditRoute(id int64, input entity.Route) error {
	_, err := service.routeRepository.FindRoute(id)
	if err != nil {
		return err
	}
	input.Id = id
	err = service.routeRepository.EditRoute(input)
	return err
}

func (service *routeService) DeleteRoute(id int64) error {
	err := service.routeRepository.DeleteRoute(id)
	return err
}
