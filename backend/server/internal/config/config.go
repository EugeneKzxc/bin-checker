package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const CfgPath = "./server/internal/config/config.yaml"

type Config struct {
	Port   uint   `yaml:"port"`
	DbPath string `yaml:"dbPath"`
}

func (c *Config) Read() error {
	f, err := os.Open(CfgPath)
	if err = yaml.NewDecoder(f).Decode(c); err != nil {
		return err
	}

	if DbPath := os.Getenv("DbPath"); DbPath != "" {
		c.DbPath = DbPath
	}

	c.Print()
	return nil
}

func (c *Config) Print() {
	log.Println("=========== SERVICE STARTED ===========")
	log.Println("PORT...........................", c.Port)
	log.Println("DbPath.........................", c.DbPath)
}
