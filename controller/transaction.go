package controller

import (
	"fmt"

	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// 交易查询
func TrxTransactionList(c *gin.Context) {
	var transactions []module.DBTransaction
	var params module.ReqTransactionList
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	start := params.Start
	length := params.Length
	blockNum := params.BlockNum

	if length == 0 {
		length = 20
	}
	fmt.Println(params)
	fmt.Println(blockNum)

	trxModule := module.TransactionCollection()
	if err := trxModule.Find(bson.M{}).Sort("-block_number").Skip(start * length).Limit(length).All(&transactions); err != nil {
		common.ResponseErr(c, "transactions search error", err)
		return
	}
	totalCount, _ := trxModule.Count()

	res := module.ResPageList{
		Data:          transactions,
		TotalRecordes: totalCount,
	}

	common.ResponseSuccess(c, "transaction list find success", res)
}

// 转账
func TrxTransferList(c *gin.Context) {
	var params module.ReqTransactionList
	var transferList []module.DBTransaction
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	start := params.Start
	length := params.Length

	if length == 0 {
		length = 20
	}

	transferModule := module.TransactionCollection()
	if err := transferModule.Find(bson.M{}).Sort("-block_number").Skip(start * length).Limit(length).All(&transferList); err != nil {
		common.ResponseErr(c, "transfer list find failed", err)
		return
	}

	totalCount, _ := transferModule.Count()
	res := module.ResPageList{
		Data:          transferList,
		TotalRecordes: totalCount,
	}
	common.ResponseSuccess(c, "transfer list find success", res)
}

// 交易详情
func TrxTransactionDetail(c *gin.Context) {
	var params module.ReqGetTransaction
	var trx module.DBTransaction

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	trxModule := module.TransactionCollection()
	if err := trxModule.Find(bson.M{"transaction_hash": params.TrxHash}).One(&trx); err != nil {
		common.ResponseErr(c, "transaction detail find failed", err)
		return
	}

	common.ResponseSuccess(c, "transaction detail find success", trx)
}
