package repository

import (
	"routefinder_golang/entity"
)

type RouteRepository interface {
	GetAllRoute() ([]entity.Route, error)
	FindRoute(id int64) (entity.Route, error)
	CreateRoute(input entity.Route) error
	EditRoute(input entity.Route) error
	DeleteRoute(id int64) error
}
