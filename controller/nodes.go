package controller

import (
	"log"

	"github.com/bottos-project/BlockExplorer/db"

	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

//NodeProductList Node product list
func NodeProductList(c *gin.Context) {
	var nodes []module.DBNodeSuper

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeSuperCollection()
	if err := nodeModule.Find(bson.M{"delegate": bson.M{"$exists": true}}).Sort("-transit_votes").All(&nodes); err != nil {
		log.Fatalf("product node list search failed: %v", err)
		common.ResponseErr(c, "product nodes search failed", err)
		return
	}

	common.ResponseSuccess(c, "product node list search success", nodes)
}

func NodeServiceList(c *gin.Context) {
	var nodes []module.DBNodeService
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeServiceCollection()
	if err := nodeModule.Find(bson.M{}).All(&nodes); err != nil {
		common.ResponseErr(c, "service node list search failed", err)
		return
	}
	common.ResponseSuccess(c, "service node list search success", nodes)
}

// NodeServiceSummary get nodes summary
func NodeServiceSummary(c *gin.Context) {
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeServiceCollection()
	superNodeModule := mongoIns.NodeSuperCollection()
	serviceNodeCount, err := nodeModule.Find(bson.M{}).Count()
	if err != nil {
		common.ResponseErr(c, "service node list search failed", err)
		return
	}
	superNodeCount, err := superNodeModule.Find(bson.M{}).Count()
	if err != nil {
		common.ResponseErr(c, "super node list search failed", err)
		return
	}
	res := module.ResServiceNodeSummary{
		DelegateCount: superNodeCount,
		ServiceCount:  serviceNodeCount,
	}

	common.ResponseSuccess(c, "service node list search success", res)
}

func NodeSuperDetail(c *gin.Context) {
	var params module.ReqNodeSuperDetail
	var node module.DBNodeSuper
	var nodes []module.DBNodeSuper
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeSuperCollection()
	if err := nodeModule.Find(bson.M{"delegate": params.NodeName}).One(&node); err != nil {
		log.Fatalf("super node detail search failed: %v", err)
		common.ResponseErr(c, "super node detail search failed", err)
		return
	}

	if err := nodeModule.Find(bson.M{"delegate": bson.M{"$exists": true}}).Sort("transit_votes").All(&nodes); err != nil {
		log.Fatalf("product node list search failed: %v", err)
		common.ResponseErr(c, "product nodes search failed", err)
		return
	}

	for index, value := range nodes {
		if node.Delegate == value.Delegate {
			node.VoteRank = index + 1
		}
	}

	common.ResponseSuccess(c, "super node detail search success", node)
}

func NodeSummaryAuto(c *gin.Context) {
	var nodes module.ResNodeSuperSummary

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeSuperCollection()
	count, err := nodeModule.Find(bson.M{"node_name": bson.M{"$exists": true}}).Count()
	if err != nil {
		log.Fatalf("product node list search failed: %v", err)
		common.ResponseErr(c, "product nodes search failed", err)
		return
	}
	nodes.DelegateCount = int32(count)

	accountModule := mongoIns.AccountCollection()

	voteAccountCount, err := accountModule.Find(bson.M{"vote.votes": bson.M{"$exists": true}}).Count()
	if err != nil {
		common.ResponseErr(c, "accounts count failed", err)
		return
	}
	nodes.VoteCustCount = int32(voteAccountCount)

	var node module.DBNodeSuper
	err = nodeModule.Find(bson.M{}).One(&node)
	if err != nil {
		log.Fatalf("product node search failed: %v", err)
		common.ResponseErr(c, "product node search failed", err)
		return
	}
	nodes.AllTransitVotes = node.AllTransitVotes
	common.ResponseSuccess(c, "product node summary search success", nodes)
}
