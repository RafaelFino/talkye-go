package config

import (
	"encoding/json"
	"log"
)

type DBConfig struct {
	Host     string `json: host`
	Port     int    `json: port`
	Username string `json: username`
	Password string `json: password`
	Database string `json: database`
}

func NewDatabaseConfig() *DBConfig {
	return &DBConfig{}
}

func (c *DBConfig) ToJson() string {
	ret, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Printf("Error marshalling config: %s", err)
	}

	return string(ret)
}

func LoadDBConfigFromJson(json string) (*DBConfig, error) {
	ret := NewDatabaseConfig()
	err := json.Unmarshal([]byte(json), ret)

	if err != nil {
		log.Printf("Error unmarshalling config: %s", err)
	}

	return ret, err
}
