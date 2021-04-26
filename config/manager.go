package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	HttpAddress string `json:"httpAddress"`
	Port        string `json:"port"`
}

func NewConfig(path string) *Configuration {

	var config Configuration

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error in Open Config file: %v", err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("error in Decode file: %v", err.Error())
	}

	return &config
}
