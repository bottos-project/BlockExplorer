package module

import (
	db "github.com/BlockExplorer/db"
	mgo "gopkg.in/mgo.v2"
)

func BlockCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("Blocks")
}

func AccountCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("Accounts")
}

func TransactionCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("Transactions")
}

func NodeServiceCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("NodeService")
}

func NodeSuperCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("NodeSuper")
}
