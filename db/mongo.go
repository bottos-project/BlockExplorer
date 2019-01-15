package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

const URL string = "mongodb://127.0.0.1:27017"

var c *mgo.Collection
var session *mgo.Session
var mongo *mgo.Database

func GetMongoDB() *mgo.Database {
	return InitMongoDB("blockchainbowser")
}

func InitMongoDB(dbname string) *mgo.Database {
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Fatalf("InintMongoDb: mongodb connect err: %v", err)
		return nil
	}
	//切换到数据库
	db := session.DB(dbname)
	mongo = db
	return db
}

// 切换collection
func ChackCollection(c string) {
	mongo.C(c)
}
