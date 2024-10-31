package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherUrlApi string
	WeatherApiKey string
	SporifyClientId string
	SporifyClientSecret string
}

func init() {
	err := godotenv.Load(".env"); if err != nil {
		log.Println("Error on load env file")
	}
}

func LoadConfig() Config {
	return Config{
		WeatherUrlApi: os.Getenv("WEATHER_API_URL"),
		WeatherApiKey: os.Getenv("WEATHER_API_KEY"),
		SporifyClientId: os.Getenv("SPOTIFY_CLIENT_ID"),
		SporifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}
}