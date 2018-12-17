package controller

import (
	"github.com/BlockExplorer/common"
	module "github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func BlockList(c *gin.Context) {
	var blockList []module.DBBlocks
	var params module.ReqBlockList

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	start := params.Start
	length := params.Length
	if length == 0 {
		length = 20
	}

	dbBlock := module.BlockCollection()
	if err := dbBlock.Find(bson.M{}).Skip(start * length).Limit(length).All(&blockList); err != nil {
		common.ResponseErr(c, "failed to select block list", err)
		return
	}

	totalCount, _ := dbBlock.Count()

	type resStruct struct {
		Data         interface{} `json:"data"`
		TotalRecords int         `json:"iTotalDisplayRecords"`
	}

	response := resStruct{
		Data:         blockList,
		TotalRecords: totalCount,
	}
	common.ResponseSuccess(c, "selected block list success", response)
}

// block detail
func BlockDetail(c *gin.Context) {
	var params module.ReqBlockDetail
	var block module.DBBlocks
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	blockModule := module.BlockCollection()
	if err := blockModule.Find(bson.M{"block_number": params.BlockNum}).One(&block); err != nil {
		common.ResponseErr(c, "block detail find failed ", err)
		return
	}

	common.ResponseSuccess(c, "block detail find success", block)
}
