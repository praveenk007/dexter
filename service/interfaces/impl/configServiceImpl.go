package impl

import (
	"fmt"

	"github.com/praveenk007/dexter/dao"

	"github.com/praveenk007/dexter/dao/interfaces"
	"github.com/praveenk007/dexter/dao/interfaces/impl"
)

type ConfigServiceImpl struct {
	session *dao.Session
}

func NewConfigService(session *dao.Session) *ConfigServiceImpl {
	return &ConfigServiceImpl{session}
}

func (cs *ConfigServiceImpl) GetConfig(id string, ctype string) {
	fmt.Println("[getBrokerConfig]")
	a := getDao(ctype)
	getConfig(a, cs.session)

}

func getDao(ctype string) interfaces.IConfigDao {
	if ctype == "broker_config" {
		return &impl.BrokerConfigDao{}
	}
	return nil
}

func getConfig(configDao interfaces.IConfigDao, session *dao.Session) {
	configDao.GetConfig(session, "broker_id")
}
