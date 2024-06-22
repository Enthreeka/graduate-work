package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type (
	Config struct {
		Redis      Redis      `json:"redis"`
		GRPC       GRPC       `json:"grpc"`
		Gateway    Gateway    `json:"gateway"`
		Prometheus Prometheus `json:"prometheus"`
	}

	GRPC struct {
		ElasticsearchService string `json:"elasticsearch_service"`
		AggregatorService    string `json:"aggregator_service"`
	}

	Gateway struct {
		Endpoint   string `json:"endpoint"`
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
		ServerName string `json:"server_name"`
	}

	Redis struct {
		Password    string `json:"password"`
		Host        string `json:"host"`
		Db          int    `json:"db"`
		MinIdleCons int    `json:"min_idle_cons"`
	}

	Prometheus struct {
		Address string `json:"address"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load("configs/proxy.env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		Gateway: Gateway{
			Port:       os.Getenv("PORT_GATEWAY"),
			ServerName: os.Getenv("SERVER_NAME"),
		},
		GRPC: GRPC{
			ElasticsearchService: os.Getenv("ELASTICSEARCH_SERVICE"),
			AggregatorService:    os.Getenv("AGGREGATOR_SERVICE"),
		},
		Redis: Redis{
			Password:    os.Getenv("PASSWORD_REDIS"),
			Host:        os.Getenv("HOST_REDIS"),
			Db:          parseEnvInt(os.Getenv("DB_REDIS")),
			MinIdleCons: parseEnvInt(os.Getenv("MIN_IDLE_CONS_REDIS")),
		},
		Prometheus: Prometheus{
			Address: os.Getenv("PROMETHEUS_ADDRESS"),
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
