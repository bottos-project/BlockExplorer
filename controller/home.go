package controller

import (
	"fmt"

	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func HomeGetTotalCount(c *gin.Context) {
	var block module.DBBlocks
	dbBlock := module.BlockCollection()
	accountModule := module.AccountCollection()
	transactionModule := module.TransactionCollection()
	if err := dbBlock.Find(bson.M{}).Sort("-block_number").One(&block); err != nil {
		common.ResponseErr(c, "failed to find lastest block number", err)
		return
	}

	accounts, _ := accountModule.Count()
	transactionCount, _ := transactionModule.Count()

	totalCount := module.ResTotalCount{
		BlockNumber:      block.BlockNumber,
		NodeCount:        "29",
		Accounts:         accounts,
		TransactionCount: transactionCount,
	}

	res := module.ResDataStruct{
		Data: totalCount,
	}

	common.ResponseSuccess(c, "find success", res)
}

// Home search
func HomeSearch(c *gin.Context) {
	var params module.ReqHomeSearch
	var blockDetail module.ResBlockDetail
	var transactionDetail module.ResTransferDetail
	var accountDetail module.ResAccountDetail

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	condition := params.Condition
	length := len(condition)

	if length == 64 {
		// condition is hash
		blockModule := module.BlockCollection()
		transactionModule := module.TransactionCollection()
		if err := blockModule.Find(bson.M{"block_hash": condition}).One(&blockDetail); err != nil {
			if err := transactionModule.Find(bson.M{"transaction_id": condition}).One(&transactionDetail); err != nil {
				common.ResponseErr(c, "search failed", err)
				return
			} else {
				// transaction hash
				transactionDetail.ResultType = "trade"
				res := module.ResSuccess{
					Data:    transactionDetail,
					Message: "搜索成功",
					Success: true,
				}
				common.ResponseSuccess(c, "搜索成功", res)
			}
		} else {
			// block hash
			blockDetail.ResultType = "block"
			blockDetail.BlockStatus = "已确认"
			res := module.ResSuccess{
				Data:    blockDetail,
				Message: "搜索成功",
				Success: true,
			}
			common.ResponseSuccess(c, "搜索成功", res)
		}
	} else {
		fmt.Printf("search account name or blockNum")
		blockModule := module.BlockCollection()
		accountModule := module.AccountCollection()
		blockNum, _ := common.String2Int(condition)
		if err := blockModule.Find(bson.M{"block_number": blockNum}).One(&blockDetail); err != nil {
			if err := accountModule.Find(bson.M{"account_name": condition}).One(&accountDetail); err != nil {
				common.ResponseErr(c, "search failed", err)
				return
			} else {
				// account name
				accountDetail.ResultType = "cust"
				res := module.ResSuccess{
					Data:    accountDetail,
					Message: "搜索成功",
					Success: true,
				}
				common.ResponseSuccess(c, "搜索成功", res)
			}
		} else {
			// block number
			blockDetail.ResultType = "block"
			blockDetail.BlockStatus = "已确认"
			res := module.ResSuccess{
				Data:    blockDetail,
				Message: "搜索成功",
				Success: true,
			}
			common.ResponseSuccess(c, "搜索成功", res)
		}
	}

}
