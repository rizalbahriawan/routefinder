package entity

import (
	"fmt"

	"github.com/go-playground/validator"
)

func (Mountain) TableName() string {
	return "ref_mountain"
}

type Mountain struct {
	Id                uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string  `json:"name" validate:"required"`
	Height            int64   `json:"height" validate:"required"`
	HeightMeasurement string  `json:"height_measurement"`
	LocationName      string  `json:"location_name"`
	Province          string  `json:"province"`
	Country           string  `json:"country"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	PhotoUrl          string  `json:"photo_url"`
	CoverImageUrl     string  `json:"cover_image_url"`
	IsVolcano         bool    `json:"is_volcano"`
	Route             []Route `json:"route" gorm:"foreignKey:IdMountain;references:Id"`
}

func (input Mountain) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

func (a Mountain) ToString() string {
	return "Mountain: " + fmt.Sprintf("%+v", a)
}
