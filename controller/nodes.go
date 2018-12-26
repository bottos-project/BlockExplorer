package controller

import (
	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Node product list
func NodeProductList(c *gin.Context) {
	var nodes []module.DBNodeSuper
	nodeModule := module.NodeSuperCollection()
	if err := nodeModule.Find(bson.M{}).All(&nodes); err != nil {
		common.ResponseErr(c, "product nodes search failed", err)
		return
	}
	common.ResponseSuccess(c, "product node list search success", nodes)
}

func NodeServiceList(c *gin.Context) {
	var nodes []module.DBNodeService
	nodeModule := module.NodeServiceCollection()
	if err := nodeModule.Find(bson.M{}).All(&nodes); err != nil {
		common.ResponseErr(c, "service node list search failed", err)
		return
	}
	common.ResponseSuccess(c, "service node list search success", nodes)
}

func NodeSuperDetail(c *gin.Context) {
	var params module.ReqNodeSuperDetail
	var node module.DBNodeSuper
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	nodeModule := module.NodeSuperCollection()
	if err := nodeModule.Find(bson.M{"node_name": params.NodeName}).One(&node); err != nil {
		common.ResponseErr(c, "super node detail search failed", err)
		return
	}
	common.ResponseSuccess(c, "super node detail search success", node)
}
