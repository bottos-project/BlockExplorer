package controller

import (
	"time"

	"github.com/bottos-project/BlockExplorer/db"

	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/module"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// HomeGetTotalCount home page detail info
func HomeGetTotalCount(c *gin.Context) {
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	dbBlock := mongoIns.BlockCollection()
	accountModule := mongoIns.AccountCollection()
	transactionModule := mongoIns.TransactionCollection()
	superNodeModule := mongoIns.NodeSuperCollection()
	ServiceModule := mongoIns.NodeServiceCollection()
	blockCount, err := dbBlock.Count()
	if err != nil {
		common.ResponseErr(c, "failed to find lastest block number", err)
		return
	}

	superNodeCount, err := superNodeModule.Count()
	if err != nil {
		common.ResponseErr(c, "failed to find super nodes", err)
		return
	}
	serviceNodeCount, err := ServiceModule.Count()
	if err != nil {
		common.ResponseErr(c, "failed to find service nodes", err)
		return
	}
	nodeCount := superNodeCount + serviceNodeCount

	accounts, _ := accountModule.Count()
	timeDate := time.Now()
	timeDate = timeDate.AddDate(0, 0, -1)
	timestamp := timeDate.Unix()

	transactionCount, _ := transactionModule.Find(bson.M{"timestamp": bson.M{"$gte": timestamp}}).Count()

	totalCount := module.ResTotalCount{
		BlockNumber:      uint64(blockCount),
		NodeCount:        nodeCount,
		Accounts:         accounts,
		TransactionCount: transactionCount,
	}

	res := module.ResDataStruct{
		Data: totalCount,
	}

	common.ResponseSuccess(c, "find success", res)
}

// HomeSearch Home search
func HomeSearch(c *gin.Context) {
	var params module.ReqHomeSearch
	var blockDetail module.ResBlockDetail
	var transactionDetail module.ResTransferDetail
	var accountDetail module.ResAccountDetail

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

	condition := params.Condition
	length := len(condition)

	if length == 64 {
		// condition is hash
		blockModule := mongoIns.BlockCollection()
		transactionModule := mongoIns.TransactionCollection()
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
		blockModule := mongoIns.BlockCollection()
		accountModule := mongoIns.AccountCollection()
		blockNum, _ := common.String2Int(condition)

		if blockNum == 0 { // accountName
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
		} else { // block number
			if err := blockModule.Find(bson.M{"block_number": blockNum}).One(&blockDetail); err != nil {
				common.ResponseErr(c, "search failed", err)
				return
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

}
