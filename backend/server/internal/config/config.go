package config

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

const CfgPath = "./server/internal/config/config.yaml"

type Config struct {
	Port        uint   `yaml:"port"`
	DbPath      string `yaml:"dbPath"`
	ExternalApi bool   `yaml:"externalApi"`
}

func (c *Config) Read() error {
	f, err := os.Open(CfgPath)
	if err = yaml.NewDecoder(f).Decode(c); err != nil {
		return err
	}

	if DbPath := os.Getenv("DbPath"); DbPath != "" {
		c.DbPath = DbPath
	}

	if ExternalApi := os.Getenv("ExternalApi"); ExternalApi != "" {
		if ExternalApiVal, err := strconv.ParseBool(ExternalApi); err == nil {
			c.ExternalApi = ExternalApiVal
		}
	}

	c.Print()
	return nil
}

func (c *Config) Print() {
	log.Println("=========== SERVICE STARTED ===========")
	log.Println("PORT...........................", c.Port)
	log.Println("DbPath.........................", c.DbPath)
	log.Println("ExternalApi....................", c.ExternalApi)
}
