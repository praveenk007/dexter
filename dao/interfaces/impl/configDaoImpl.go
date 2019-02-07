package impl

import (
	"fmt"
)

type BrokerConfigDao struct {
	//interfaces.IConfigDao
}

func (s *BrokerConfigDao) GetConfig(id string) {
	fmt.Println("[GetConfig]")
	NewAbstractDao().FetchById(id)
}
