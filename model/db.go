package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const URL = "localhost:27017"

func db() {
	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("merchant_info_development") //数据库名称
	collection := db.C("application_records")     //如果该集合已经存在的话，则直接返回
	return db
}
