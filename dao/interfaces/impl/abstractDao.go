package impl

import (
	"fmt"
)

type AbstractDao struct {
}

func NewAbstractDao() *AbstractDao {
	return &AbstractDao{}
}

func (abstractDat AbstractDao) FetchById(id string) {
	fmt.Println("AbstractDao.FetchById called")
}
