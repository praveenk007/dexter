package impl

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/praveenk007/dexter/dao"
)

type AbstractDao struct {
	session *dao.Session
}

func NewAbstractDao(session *dao.Session) *AbstractDao {
	return &AbstractDao{session}
}

func (abstractDao *AbstractDao) FetchById(key string, value string, db string, collection string) *[]byte {
	var result bson.M
	collectionObject := abstractDao.session.GetCollection(db, collection)
	collectionObject.Find(bson.M{key: value}).One(&result)
	j, _ := json.Marshal(result)
	return &j
}
