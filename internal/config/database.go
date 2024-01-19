package config

import (
	"encoding/json"
	"fmt"
	"log"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
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

func (c *DBConfig) ToConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Username, c.Password, c.Database)
}
