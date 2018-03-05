package model

import "gopkg.in/mgo.v2"

var GlobalMgoSession *mgo.Session

func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}
