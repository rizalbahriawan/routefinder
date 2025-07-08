package service

import (
	"routefinder_golang/entity"
	"routefinder_golang/repository"
)

type WaypointsService interface {
	GetAllWaypoints() ([]entity.Waypoints, error)
	FindWaypoints(id int64) (entity.Waypoints, error)
	CreateWaypoints(input entity.Waypoints) error
	EditWaypoints(id int64, input entity.Waypoints) error
	DeleteWaypoints(id int64) error
}

type waypointsService struct {
	waypointsRepository repository.WaypointsRepository
}

func NewWaypointsService(repository repository.WaypointsRepository) *waypointsService {
	return &waypointsService{repository}
}

func (service *waypointsService) GetAllWaypoints() ([]entity.Waypoints, error) {
	result, err := service.waypointsRepository.GetAllWaypoints()
	return result, err
}

func (service *waypointsService) FindWaypoints(id int64) (entity.Waypoints, error) {
	result, err := service.waypointsRepository.FindWaypoints(id)
	return result, err
}

func (service *waypointsService) CreateWaypoints(input entity.Waypoints) error {
	err := service.waypointsRepository.CreateWaypoints(input)
	return err
}

func (service *waypointsService) EditWaypoints(id int64, input entity.Waypoints) error {
	_, err := service.waypointsRepository.FindWaypoints(id)
	if err != nil {
		return err
	}
	input.Id = id
	err = service.waypointsRepository.EditWaypoints(input)
	return err
}

func (service *waypointsService) DeleteWaypoints(id int64) error {
	err := service.waypointsRepository.DeleteWaypoints(id)
	return err
}
