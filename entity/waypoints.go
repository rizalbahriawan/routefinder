package entity

import (
	"fmt"

	"github.com/go-playground/validator"
)

func (Waypoints) TableName() string {
	return "waypoints"
}

type Waypoints struct {
	Id                   int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	IdRoute              string  `json:"id_route" validate:"required"`
	Name                 string  `json:"name" validate:"required"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	Elevation            int64   `json:"elevation"`
	ElevationMeasurement string  `json:"elevation_measurement"`
	Description          string  `json:"description"`
	OrderIndex           int64   `json:"order_index"`
}

func (input Waypoints) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

func (a Waypoints) ToString() string {
	return "Waypoints: " + fmt.Sprintf("%+v", a)
}
