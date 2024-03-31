package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

type (
	Config struct {
		HTTPServer HTTPServer `json:"http_server"`
		JSON       JSON       `json:"elasticsearch"`
	}

	HTTPServer struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}
)

type (
	JSON struct {
		Elasticsearch Elasticsearch `json:"elasticsearch"`
		Kibana        Kibana        `json:"kibana"`
	}

	Elasticsearch struct {
		Addr []string `json:"addr"`
	}

	Kibana struct {
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
		HTTPServer: HTTPServer{
			Hostname:   os.Getenv("HTTP_SERVER_HOSTNAME"),
			Port:       os.Getenv("HTTP_SERVER_PORT"),
			TypeServer: os.Getenv("HTTP_SERVER_TYPE_SERVER"),
		},
		JSON: JSON{
			Elasticsearch: Elasticsearch{
				Addr: jsonFile.Elasticsearch.Addr,
			},
			Kibana: Kibana{
				Addr: jsonFile.Kibana.Addr,
			},
		},
	}

	return config, nil
}

func NewDecoderJSON(path string) (*JSON, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &JSON{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
