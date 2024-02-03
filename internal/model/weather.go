package model

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	Localtime string `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	TempC      float64   `json:"temp_c"`
	IsDay      int       `json:"is_day"`
	Condition  Condition `json:"condition"`
	WindKPH    float64   `json:"wind_kph"`
	Humidity   int       `json:"humidity"`
	Cloud      int       `json:"cloud"`
	FeelsLikeC float64   `json:"feelslike_c"`
	VisKM      float64   `json:"vis_km"`
}

type WeatherData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
