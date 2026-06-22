package system

import (
	"encoding/json"
	"os"
)

type Config struct {
	LibraryDatas []LibraryData `json:"library_paths"`
}

type LibraryData struct {
	Name 	string	`json:"name"`
	Path 	string	`json:"path"`
	Enabled bool	`json:"enabled"`
}

func GetConfig() (*Config, error){
	file, err := os.Open("./config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
