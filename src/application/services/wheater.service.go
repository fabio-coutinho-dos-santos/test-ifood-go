package services

import (
	"fmt"
	http_errors "ifood-backend-test/src/application/http-errors"
	"io"
	"net/http"
)

type IWeatherService interface {
	GetTempByLatLong(lat float32, long float32) (string, error)
	GetTempByCity(lat float32, long float32) (string, error)
}

type WeatherService struct {
	urlApi string
	apiKey string
}

func NewWeatherService(urlApi string, apiKey string) (*WeatherService) {
	return &WeatherService{
		urlApi: urlApi,
		apiKey: apiKey,
	}
}

func (w WeatherService) FindByLocation(lat string, long string) (string, *http_errors.HttpError, error) {
	url := w.urlApi+"?lat="+lat+"&lon="+long+"&appid="+w.apiKey+"&units=metric"
	
	resp, err := http.Get(url)

	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", nil, err
	}

	status := resp.StatusCode

	if status != http.StatusOK {
		fmt.Println(string(body))
		return "", http_errors.NewHttpError(resp.StatusCode, string(body)), nil
	}

	return string(body), nil, nil
}

func (w WeatherService) FindByCity(city string) (string, *http_errors.HttpError, error) {
	url := w.urlApi+"?q="+city+"&appid="+w.apiKey+"&units=metric"

	resp, err := http.Get(url)	

	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", nil, err
	}

	status := resp.StatusCode

	if status != http.StatusOK {
		fmt.Println(string(body))
		return "", http_errors.NewHttpError(resp.StatusCode, string(body)), nil
	}

	return string(body), nil, nil
}