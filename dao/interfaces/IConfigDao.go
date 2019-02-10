package interfaces

import "github.com/praveenk007/dexter/dao"

type IConfigDao interface {
	GetConfig(s *dao.Session, id string)
}
