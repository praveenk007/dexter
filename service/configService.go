package service

import (
	"dexter/dao/interfaces"
	"dexter/dao/interfaces/impl"
	"fmt"
)

type ConfigService struct {
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (cs *ConfigService) GetBrokerConfig() {
	fmt.Println("[getBrokerConfig]")
	a := impl.BrokerConfigDao{}
	getConfig(&a)

}

func getConfig(configDao interfaces.IConfigDao) {
	configDao.GetConfig("broker_id")
}
