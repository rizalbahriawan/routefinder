package repository

import (
	"routefinder_golang/entity"
	"routefinder_golang/request"
)

type MountainRepository interface {
	GetAllMountain() ([]entity.Mountain, error)
	FindMountain(id int64) (entity.Mountain, error)
	CreateMountain(input request.MountainRequest) error
	EditMountain(input request.MountainRequest) error
	DeleteMountain(id int64) error
}
