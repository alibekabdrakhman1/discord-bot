package transport

import (
	"encoding/json"
	"fmt"
	"github.com/alibekabdrakhman1/discord-bot/internal/config"
	"github.com/alibekabdrakhman1/discord-bot/internal/model"
	"net/http"
)

type WeatherClient struct {
	APIKey string
}

func NewWeatherClient(config *config.Config) *WeatherClient {
	return &WeatherClient{
		APIKey: config.Weather.APIKey,
	}
}

func (client *WeatherClient) GetWeather(city string) (*model.WeatherData, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", client.APIKey, city)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return &model.WeatherData{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return &model.WeatherData{}, fmt.Errorf("invalid city: %s", city)
	}
	fmt.Println(url)

	var weatherResponse model.WeatherData
	err = json.NewDecoder(response.Body).Decode(&weatherResponse)
	fmt.Println(weatherResponse)
	if err != nil {
		return &model.WeatherData{}, err
	}

	return &weatherResponse, nil
}
