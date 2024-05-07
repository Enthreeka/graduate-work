package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

type (
	Config struct {
		App App `json:"app"`
	}
)

type (
	App struct {
		GRPC     GRPC     `json:"GRPC"`
		Postgres Postgres `json:"postgres"`
	}

	Postgres struct {
		LocalAddr  string `json:"local_addr"`
		DockerAddr string `json:"docker_addr"`
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
			GRPC: GRPC{
				Addr: jsonFile.GRPC.Addr,
			},
			Postgres: Postgres{
				LocalAddr:  jsonFile.Postgres.LocalAddr,
				DockerAddr: jsonFile.Postgres.DockerAddr,
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
