package impl

import (
	"fmt"

	"github.com/praveenk007/dexter/models"

	"github.com/praveenk007/dexter/dao"
)

type AbstractDao struct {
	session *dao.Session
}

func NewAbstractDao(session *dao.Session) *AbstractDao {
	return &AbstractDao{session}
}

func (abstractDao *AbstractDao) FetchById(id string) {
	var results []models.BrokerConfig
	collection := abstractDao.session.GetCollection("mintpro", "BrokerConfig")
	collection.Find(nil).All(&results)
	//fmt.Println(results[0].SessionDomain)
	fmt.Println("=====")
	fmt.Println(results[0])
	fmt.Println("=====")

}
