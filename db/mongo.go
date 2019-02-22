package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// const URL string = "mongodb://118.184.215.41:27018"
const URL string = "mongodb://125.94.34.23:27018"

//const URL string = "mongodb://114.67.80.209:27018"
// const URL string = "mongodb://127.0.0.1:27017"

type MongoDB struct {
	session *mgo.Session
	mongo   *mgo.Database
}

// NewDBCollection create instance of MongoDB
func NewDBCollection() (*MongoDB, error) {
	s, err := mgo.Dial(URL)
	if err != nil {
		log.Printf("New mongo dial fail: %v", err)
		return nil, err
	}
	dbCollection := MongoDB{
		session: s,
	}
	dbCollection.mongo = dbCollection.session.DB("blockchainbowser")
	return &dbCollection, nil
}

//CloseSession close mongodb's session
func CloseSession(d *MongoDB) {
	if d != nil && d.session != nil {
		d.session.Close()
	}
}

//CheckCollection 切换collection
func (d *MongoDB) CheckCollection(c string) *mgo.Collection {
	return d.mongo.C(c)
}

//BlockCollection get block collection
func (d *MongoDB) BlockCollection() *mgo.Collection {
	return d.mongo.C("Blocks")
}

//AccountCollection get account collection
func (d *MongoDB) AccountCollection() *mgo.Collection {
	return d.mongo.C("Accounts")
}

//TransactionCollection get transaction collection
func (d *MongoDB) TransactionCollection() *mgo.Collection {
	return d.mongo.C("Transactions")
}

//StatisticCollection get statistic collection
func (d *MongoDB) StatisticCollection() *mgo.Collection {
	return d.mongo.C("Statistic")
}

//NodeServiceCollection get service node collection
func (d *MongoDB) NodeServiceCollection() *mgo.Collection {
	return d.mongo.C("NodeService")
}

//NodeSuperCollection get super node collection
func (d *MongoDB) NodeSuperCollection() *mgo.Collection {
	return d.mongo.C("NodeSuper")
}

//GetCollectionByName is get collection by name param
func (d *MongoDB) GetCollectionByName(collectionName string) *mgo.Collection {
	return d.mongo.C(collectionName)
}
