package controller

import (
	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func HomeGetTotalCount(c *gin.Context) {
	var block module.DBBlocks
	dbBlock := module.BlockCollection()
	if err := dbBlock.Find(bson.M{}).Sort("-block_number").One(&block); err != nil {
		common.ResponseErr(c, "failed to find lastest block number", err)
		return
	}

	res := module.ResTotalCount{
		BlockNumber: block.BlockNumber,
	}

	common.ResponseSuccess(c, "find success", res)
}
