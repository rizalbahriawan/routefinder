package entity

import (
	"fmt"

	"github.com/go-playground/validator"
)

func (Weather) TableName() string {
	return "weather_data"
}

type Weather struct {
	Id                     int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	IdMountain             string  `json:"id_mountain" validate:"required"`
	Temperature            float64 `json:"temperature"`
	TemperatureMeasurement string  `json:temperature_measurement`
	WindSpeed              int64   `json:wind_speed`
	WindSpeedDirection     string  `json:wind_speed_direction`
	Visibility             int64   `json:visibility`
	ForecastDate           string  `forecast_date`
}

func (input Weather) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

func (a Weather) ToString() string {
	return "Weather: " + fmt.Sprintf("%+v", a)
}
