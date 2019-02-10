package dao

import (
	mgo "gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
}

func NewSession(url string) (*Session, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}
