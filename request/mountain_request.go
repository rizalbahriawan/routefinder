package request

import "github.com/go-playground/validator"

func (MountainRequest) TableName() string {
	return "ref_mountain"
}

type MountainRequest struct {
	Id                int64   `json:"id"`
	Name              string  `json:"name" validate:"required"`
	Height            int64   `json:"height" validate:"required"`
	HeightMeasurement string  `json:"height_measurement"`
	LocationName      string  `json:"location_name"`
	Province          string  `json:"province"`
	Country           string  `json:"country"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	PhotoUrl          string  `json:"photourl"`
	CoverImageUrl     string  `json:"coverimageurl"`
	IsVolcano         bool    `json:"is_volcano"`
}

func (input MountainRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
