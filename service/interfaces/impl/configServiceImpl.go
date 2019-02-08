package impl

import (
	"dexter/dao/interfaces"
	"dexter/dao/interfaces/impl"
	"fmt"
)

type ConfigServiceImpl struct {
}

func NewConfigService() *ConfigServiceImpl {
	return &ConfigServiceImpl{}
}

func (cs *ConfigServiceImpl) GetConfig(id string, ctype string) {
	fmt.Println("[getBrokerConfig]")
	a := getDao(ctype)
	getConfig(a)

}

func getDao(ctype string) interfaces.IConfigDao {
	if ctype == "broker_config" {
		return &impl.BrokerConfigDao{}
	}
	return nil
}

func getConfig(configDao interfaces.IConfigDao) {
	configDao.GetConfig("broker_id")
}
