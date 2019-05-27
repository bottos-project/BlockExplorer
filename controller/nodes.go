package controller

import (
	"log"

	"github.com/bottos-project/BlockExplorer/db"

	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// //NodeProductList Node product list
// func NodeProductList(c *gin.Context) {
// 	var nodes []module.DBNodeSuper

// 	mongoIns, err := db.NewDBCollection()
// 	if err != nil {
// 		common.ResponseErr(c, "connect mongoDB error", err)
// 		return
// 	}

// 	defer db.CloseSession(mongoIns)

// 	nodeModule := mongoIns.NodeSuperCollection()
// 	if err := nodeModule.Find(bson.M{"delegate": bson.M{"$exists": true}}).Sort("-votes").All(&nodes); err != nil {
// 		log.Fatalf("product node list search failed: %v", err)
// 		common.ResponseErr(c, "product nodes search failed", err)
// 		return
// 	}

// 	common.ResponseSuccess(c, "product node list search success", nodes)
// }

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
	accountModule := mongoIns.AccountCollection()
	trxModule := mongoIns.TransactionCollection()
	blockModule := mongoIns.BlockCollection()
	res := bson.M{}
	account := bson.M{}
	// node info
	if err := nodeModule.Find(bson.M{"delegate": params.NodeName}).One(&res); err != nil {
		common.ResponseErr(c, "super node detail search failed", err)
		return
	}

	// account info
	accountModule.Find(bson.M{"account_name": params.NodeName}).One(&account)
	res["detail"] = account

	// blocks
	blockCount, _ := blockModule.Find(bson.M{"delegate": params.NodeName}).Count()
	res["produce_block_count"] = blockCount

	// voters
	voteCount, _ := trxModule.Find(bson.M{"method": "votedelegate", "param.delegate": params.NodeName}).Count()
	res["voters"] = voteCount

	common.ResponseSuccess(c, "super node detail search success", res)
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

//NodeProductList Node product list
func NodeProductList(c *gin.Context) {
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	selector := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "Accounts",
				"localField":   "delegate",
				"foreignField": "account_name",
				"as":           "account_info",
			},
		},
		{
			"$replaceRoot": bson.M{ // 文档替换
				"newRoot": bson.M{
					"$mergeObjects": []interface{}{ // 文档合并
						bson.M{"$arrayElemAt": []interface{}{"$account_info", 0}},
						"$$ROOT",
					},
				},
			},
		},
		{ // 过滤有效字段
			"$project": bson.M{"account_info": 0},
		},
	}
	nodeModule := mongoIns.NodeSuperCollection()

	var res []bson.M

	if err := nodeModule.Pipe(selector).All(&res); err != nil {
		common.ResponseErr(c, "product nodes search failed", err)
		return
	}
	common.ResponseSuccess(c, "product node list search success", res)
}

func NodeVoterList(c *gin.Context) {
	var params module.ReqNodeSuperDetail
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
	data := []bson.M{}

	start := params.Start
	length := params.Length

	trxModule := mongoIns.TransactionCollection()
	if err := trxModule.Find(bson.M{"method": "votedelegate", "param.delegate": params.NodeName}).Sort("-_id").Skip(start).Limit(length).All(&data); err != nil {
		common.ResponseErr(c, "super node voters search failed", err)
		return
	}

	count, _ := trxModule.Find(bson.M{"method": "votedelegate", "param.delegate": params.NodeName}).Count()

	res := bson.M{
		"data":                 data,
		"iTotalDisplayRecords": count,
	}

	common.ResponseSuccess(c, "super node voters search success", res)
}
