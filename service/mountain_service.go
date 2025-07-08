package service

import (
	"routefinder_golang/entity"
	"routefinder_golang/repository"
	"routefinder_golang/request"
)

type MountainService interface {
	GetAllMountain() ([]entity.Mountain, error)
	FindMountain(id int64) (entity.Mountain, error)
	CreateMountain(input request.MountainRequest) error
	EditMountain(id int64, input request.MountainRequest) error
	DeleteMountain(id int64) error
}

type mountainService struct {
	mountainRepository repository.MountainRepository
}

func NewMountainService(repository repository.MountainRepository) *mountainService {
	return &mountainService{repository}
}

func (service *mountainService) GetAllMountain() ([]entity.Mountain, error) {
	result, err := service.mountainRepository.GetAllMountain()
	return result, err
}

func (service *mountainService) FindMountain(id int64) (entity.Mountain, error) {
	result, err := service.mountainRepository.FindMountain(id)
	return result, err
}

func (service *mountainService) CreateMountain(input request.MountainRequest) error {
	input.HeightMeasurement = "meter"
	err := service.mountainRepository.CreateMountain(input)
	return err
}

func (service *mountainService) EditMountain(id int64, input request.MountainRequest) error {
	_, err := service.mountainRepository.FindMountain(id)
	if err != nil {
		return err
	}
	input.Id = id
	err = service.mountainRepository.EditMountain(input)
	return err
}

func (service *mountainService) DeleteMountain(id int64) error {
	err := service.mountainRepository.DeleteMountain(id)
	return err
}
