package repository

import (
	"routefinder_golang/entity"
)

type WaypointsRepository interface {
	GetAllWaypoints() ([]entity.Waypoints, error)
	FindWaypoints(id int64) (entity.Waypoints, error)
	CreateWaypoints(input entity.Waypoints) error
	EditWaypoints(input entity.Waypoints) error
	DeleteWaypoints(id int64) error
}
