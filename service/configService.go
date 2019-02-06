package service

import (
	"fmt"
)

type ConfigService struct {
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (cs *ConfigService) GetBrokerConfig() {
	fmt.Println("[getBrokerConfig]")
}
