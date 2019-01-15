package module

import (
	"github.com/bottos-project/BlockExplorer/db"
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

func StatisticCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("Statistic")
}

func NodeServiceCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("NodeService")
}

func NodeSuperCollection() *mgo.Collection {
	mongo := db.GetMongoDB()
	return mongo.C("NodeSuper")
}
