package config

import (
	"encoding/json"
	"log"
)

type BrokerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewBrokerConfig() *BrokerConfig {
	return &BrokerConfig{}
}

func (c *BrokerConfig) ToJson() string {
	ret, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Printf("Error marshalling config: %s", err)
	}

	return string(ret)
}
