package controller

import (
	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func BlockList(c *gin.Context) {
	var blockList []module.ResBlockList
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

	selectInfo := bson.M{"block_hash": 1, "block_number": 1, "delegate": 1, "timestamp": 1, "transaction_count": 1}
	if params.DelegateName != "" {
		selectInfo["delegate"] = params.DelegateName
	}
	dbBlock := module.BlockCollection()
	if err := dbBlock.Find(bson.M{}).Select(selectInfo).Sort("-block_number").Skip(start).Limit(length).All(&blockList); err != nil {
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
	var block module.ResBlockDetail
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}
	// string to int
	blockNumInt, _ := common.String2Int(params.BlockNum)
	blockModule := module.BlockCollection()
	if err := blockModule.Find(bson.M{"block_number": blockNumInt}).One(&block); err != nil {
		common.ResponseErr(c, "block detail find failed ", err)
		return
	}

	if block.BlockNumber > 0 {
		block.PrevBlockNum = block.BlockNumber - 1
	} else {
		block.PrevBlockNum = 0
	}

	block.BlockStatus = "已确认"
	block.ResultType = "block"

	common.ResponseSuccess(c, "block detail find success", block)
}
