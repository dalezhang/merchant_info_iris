package model

import (
	//"fmt"
	//"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name        string
	Email       string
	Bucket_name string
	Bucket_url  string
}

type AllUser struct {
	Users []User
}

//const URL = "localhost:27017"

//func GetUsers() []User {
//session, err := mgo.Dial(URL) //连接数据库
//if err != nil {
//panic(err)
//}
//defer session.Close()
//session.SetMode(mgo.Monotonic, true)
//db := session.DB("merchant_info_development") //数据库名称
//collection := db.C("application_records")     //如果该集合已经存在的话，则直接返回

//var all_user AllUser
//result := User{}
//iter := collection.Find(bson.M{"_type": "User"}).Iter()
//for iter.Next(&result) {
//all_user.Users = append(all_user.Users, result)
//}
//return all_user.Users
//}

func GetUsers() []User {

	var all_user AllUser
	result := User{}
	iter := db.Find(bson.M{"_type": "User"}).Iter()
	for iter.Next(&result) {
		all_user.Users = append(all_user.Users, result)
	}
	return all_user.Users

}
