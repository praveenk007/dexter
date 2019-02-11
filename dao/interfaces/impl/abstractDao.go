package impl

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

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
	var result models.BrokerConfig
	collection := abstractDao.session.GetCollection("mintpro", "BrokerConfig")
	collection.Find(bson.M{"broker": id}).One(&result)
	fmt.Println(result)
}
