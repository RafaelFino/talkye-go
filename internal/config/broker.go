package config

import (
	"encoding/json"
	"log"
)

type BrokerConfig struct {
	Host string `json: host`
	Port int    `json: port`
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

func LoadFromJson(json string) (*BrokerConfig, error) {
	ret := NewBrokerConfig()
	err := json.Unmarshal([]byte(json), ret)

	if err != nil {
		log.Printf("Error unmarshalling config: %s", err)
	}

	return ret, nil
}
