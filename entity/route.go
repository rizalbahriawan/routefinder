package entity

import (
	"fmt"

	"github.com/go-playground/validator"
)

func (Route) TableName() string {
	return "route"
}

type Route struct {
	Id                       int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	IdMountain               int64   `json:"id_mountain" validate:"required"`
	Name                     string  `json:"name" validate:"required"`
	Distance                 float64 `json:"distance"`
	DistanceMeasurement      string  `json:"distance_measurement"`
	IsLoop                   bool    `json:"is_loop"`
	IsOfficial               bool    `json:"is_official"`
	PermitRequired           bool    `json:"permit_required"`
	PermitContact            string  `json:"permit_contact"`
	ElevationGain            float64 `json:"elevation_gain"`
	ElevationGainMeasurement string  `json:"elevation_gain_measurement"`
	EstimatedTimeMinutes     int64   `json:"estimated_time_minutes"`
	Difficulty               string  `json:"difficulty"`
}

func (input Route) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

func (a Route) ToString() string {
	return "Route: " + fmt.Sprintf("%+v", a)
}
