package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const fileName = "config.yaml"

func New() *Config {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
	}

	p := &Config{}
	err = yaml.Unmarshal(buf, p)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return p
}
