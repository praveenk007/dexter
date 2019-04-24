package impl

import (
	"fmt"

	"github.com/praveenk007/dexter/dao"
)

type BrokerConfigDao struct {
	//session *dao.Session
}

func (s *BrokerConfigDao) GetConfig(sess *dao.Session, id string) *[]byte {
	fmt.Println("[GetConfig]")
	return NewAbstractDao(sess).FetchById("broker", id, "mintpro", "BrokerConfig")
}
