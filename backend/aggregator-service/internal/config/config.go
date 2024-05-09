package config

import (
	"encoding/json"
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
		Redis    Redis    `json:"redis"`
	}

	Postgres struct {
		LocalAddr  string `json:"local_addr"`
		DockerAddr string `json:"docker_addr"`
	}

	GRPC struct {
		Addr string `json:"addr"`
	}

	Redis struct {
		Password    string `json:"password"`
		Host        string `json:"host"`
		Db          int    `json:"db"`
		MinIdleCons int    `json:"min_idle_cons"`
	}
)

func New() (*Config, error) {
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
			Redis: Redis{
				Password:    jsonFile.Redis.Password,
				Host:        jsonFile.Redis.Host,
				Db:          jsonFile.Redis.Db,
				MinIdleCons: jsonFile.Redis.MinIdleCons,
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
