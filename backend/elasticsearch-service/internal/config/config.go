package config

import (
	"github.com/joho/godotenv"
	"os"
)

type (
	Config struct {
		HTTPServer HTTPServer `json:"http_server"`
	}

	HTTPServer struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load("configs/app.env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		HTTPServer: HTTPServer{
			Hostname:   os.Getenv("HTTP_SERVER_HOSTNAME"),
			Port:       os.Getenv("HTTP_SERVER_PORT"),
			TypeServer: os.Getenv("HTTP_SERVER_TYPE_SERVER"),
		},
	}

	return config, nil
}
