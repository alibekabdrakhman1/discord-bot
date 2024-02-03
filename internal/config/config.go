package config

type Config struct {
	Discord   Discord   `yaml:"Discord"`
	Weather   Weather   `yaml:"Weather"`
	Translate Translate `yaml:"Translate"`
}

type Discord struct {
	Token string `yaml:"Token"`
}

type Weather struct {
	APIKey string `yaml:"APIKey"`
}

type Translate struct {
	APIKey string `yaml:"APIKey"`
}
