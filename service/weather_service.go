package service

// import (
// 	"routefinder_golang/entity"
// 	"routefinder_golang/repository"
// )

// type WeatherService interface {
// 	GetAllWeather() ([]entity.Weather, error)
// 	FindWeather(id int64) (entity.Weather, error)
// 	CreateWeather(input entity.Weather) error
// 	EditWeather(id int64, input entity.Weather) error
// 	DeleteWeather(id int64) error
// }

// type weatherService struct {
// 	weatherRepository repository.WeatherRepository
// }

// func NewWeatherService(repository repository.WeatherRepository) *weatherService {
// 	return &weatherService{repository}
// }

// func (service *weatherService) GetAllWeather() ([]entity.Weather, error) {
// 	result, err := service.weatherRepository.GetAllWeather()
// 	return result, err
// }

// func (service *weatherService) FindWeather(id int64) (entity.Weather, error) {
// 	result, err := service.weatherRepository.FindWeather(id)
// 	return result, err
// }

// func (service *weatherService) CreateWeather(input entity.Weather) error {
// 	err := service.weatherRepository.CreateWeather(input)
// 	return err
// }

// func (service *weatherService) EditWeather(id int64, input entity.Weather) error {
// 	_, err := service.weatherRepository.FindWeather(id)
// 	if err != nil {
// 		return err
// 	}
// 	input.Id = id
// 	err = service.weatherRepository.EditWeather(input)
// 	return err
// }

// func (service *weatherService) DeleteWeather(id int64) error {
// 	err := service.weatherRepository.DeleteWeather(id)
// 	return err
// }
