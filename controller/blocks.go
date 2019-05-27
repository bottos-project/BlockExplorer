package controller

import (
	"fmt"
	"log"

	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

//BlockList for delegate or anyone
func BlockList(c *gin.Context) {
	var blockList []module.ResBlockList
	var params module.ReqBlockList

	type resStruct struct {
		Data         interface{} `json:"data"`
		TotalRecords int         `json:"iTotalDisplayRecords"`
	}

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

	selectInfo := bson.M{}
	if params.DelegateName != "" {
		selectInfo["delegate"] = params.DelegateName
	}
	dbBlock := mongoIns.BlockCollection()
	fmt.Println(selectInfo)
	blockCount, err := dbBlock.Find(selectInfo).Count()
	log.Print(blockCount)
	if err != nil {
		common.ResponseErr(c, "failed to select block list than get block count", err)
		return
	}

	start := params.Start
	length := params.Length
	// start, length = paging(start, length, blockCount)
	if length <= 0 {
		common.ResponseSuccess(c, "selected block list success", resStruct{TotalRecords: blockCount})
		return
	}

	if err := dbBlock.Find(selectInfo).Sort("-_id").Skip(start).Limit(length).All(&blockList); err != nil {
		common.ResponseErr(c, "failed to select block list", err)
		return
	}

	// if err := dbBlock.Find(selectInfo).Skip(start).Limit(length).All(&blockList); err != nil {
	// 	common.ResponseErr(c, "failed to select block list", err)
	// 	return
	// }

	response := resStruct{
		Data:         blockList,
		TotalRecords: blockCount,
	}
	common.ResponseSuccess(c, "selected block list success", response)
}

//BlockDetail block detail
func BlockDetail(c *gin.Context) {
	var params module.ReqBlockDetail
	var block module.ResBlockDetail
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
	// string to int
	blockNumInt, _ := common.String2Int(params.BlockNum)
	blockModule := mongoIns.BlockCollection()
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
