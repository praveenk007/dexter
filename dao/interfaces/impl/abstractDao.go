package impl

import (
	"fmt"

	"github.com/praveenk007/dexter/dao"
)

type AbstractDao struct {
	session *dao.Session
}

func NewAbstractDao(session *dao.Session) *AbstractDao {
	return &AbstractDao{session}
}

func (abstractDao *AbstractDao) FetchById(id string) {
	type result struct {
		//Id     bson.ObjectId `bson:"_id,omitempty"`
		Broker        string //`bson:"broker,omitempty"`
		SessionDomain string `bson:"sessionDomain,omitempty"`
	}
	var results []result
	collection := abstractDao.session.GetCollection("mintpro", "BrokerConfig")
	count, _ := collection.Count()
	fmt.Printf("coll %d", count)
	collection.Find(nil).All(&results)
	fmt.Println(results[0].SessionDomain)
}
