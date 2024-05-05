package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

type (
	Config struct {
		App App `json:"elasticsearch"`
	}
)

type (
	App struct {
		IsDebugLevel bool `json:"is_debug_level"`

		Elasticsearch Elasticsearch `json:"elasticsearch"`
		Kibana        Kibana        `json:"kibana"`
		GRPC          GRPC          `json:"GRPC"`
	}

	Elasticsearch struct {
		Addr []string `json:"addr"`

		SearchFields []string `json:"search_fields"`
	}

	Kibana struct {
		Addr string `json:"addr"`
	}

	GRPC struct {
		Addr string `json:"addr"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load("configs/app.env")
	if err != nil {
		return nil, err
	}

	jsonFile, err := NewDecoderJSON("configs/app.json")
	if err != nil {
		return nil, err
	}

	config := &Config{
		App: App{
			Elasticsearch: Elasticsearch{
				Addr:         jsonFile.Elasticsearch.Addr,
				SearchFields: jsonFile.Elasticsearch.SearchFields,
			},
			Kibana: Kibana{
				Addr: jsonFile.Kibana.Addr,
			},
			GRPC: GRPC{
				Addr: jsonFile.GRPC.Addr,
			},
		},
	}

	return config, nil
}

func NewDecoderJSON(path string) (*App, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &App{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
