package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type (
	Config struct {
		Redis   Redis   `json:"redis"`
		GRPC    GRPC    `json:"grpc"`
		Gateway Gateway `json:"gateway"`
	}

	GRPC struct {
		ElasticsearchService string `json:"elasticsearch_service"`
	}

	Gateway struct {
		Endpoint   string `json:"endpoint"`
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}

	Redis struct {
		Password    string `json:"password"`
		Host        string `json:"host"`
		Db          int    `json:"db"`
		MinIdleCons int    `json:"min_idle_cons"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load("configs/proxy.env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		Gateway: Gateway{
			Port: os.Getenv("port"),
		},
		GRPC: GRPC{
			ElasticsearchService: os.Getenv("elasticsearchService"),
		},
	}

	return config, nil
}

func parseEnvInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}
