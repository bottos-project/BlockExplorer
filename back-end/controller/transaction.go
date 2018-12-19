package controller

import (
	"fmt"

	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// transaction list
func TrxTransactionList(c *gin.Context) {
	var transactions []module.ResTransactions
	var params module.ReqTransactionList
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	start := params.Start
	length := params.Length
	blockNum, _ := common.String2Int(params.BlockNum)

	if length == 0 {
		length = 20
	}
	findInfo := bson.M{}
	if blockNum != 0 {
		findInfo["block_number"] = blockNum
	}

	trxModule := module.TransactionCollection()
	selectInfo := bson.M{"block_number": 1, "transaction_id": 1, "sender": 1, "contract": 1, "method": 1, "timestamp": 1}
	if err := trxModule.Find(findInfo).Select(selectInfo).Sort("-block_number").Skip(start).Limit(length).All(&transactions); err != nil {
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
	var resTransferList []module.ResTransfers
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	blockNum, _ := common.String2Int(params.BlockNum)
	start := params.Start
	length := params.Length

	if length == 0 {
		length = 20
	}

	findInfo := bson.M{"method": "transfer"}
	if blockNum != 0 {
		findInfo["block_number"] = blockNum
	}

	transferModule := module.TransactionCollection()
	if err := transferModule.Find(findInfo).Sort("-block_number").Skip(start).Limit(length).All(&transferList); err != nil {
		common.ResponseErr(c, "transfer list find failed", err)
		return
	}

	for i := 0; i < len(transferList); i++ {
		t := transferList[i]
		tempT := module.ResTransfers{
			BlockNumber:   t.BlockNumber,
			From:          t.Param["from"].(string),
			To:            t.Param["to"].(string),
			Value:         t.Param["value"].(float64),
			CoinType:      "BTO",
			TransactionID: t.TransactionID,
			TimeStamp:     t.TimeStamp,
		}
		resTransferList = append(resTransferList, tempT)
	}

	totalCount, _ := transferModule.Find(bson.M{"method": "transfer"}).Count()
	res := module.ResPageList{
		Data:          resTransferList,
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

// personal transfer list
func TrxPersonalTransferList(c *gin.Context) {
	var params module.ReqTransactionList
	var transferList []module.DBTransaction
	var resTransferList []module.ResTransfers

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	start := params.Start
	length := params.Length

	if length == 0 {
		length = 20
	}
	fmt.Println(params)

	trxModule := module.TransactionCollection()
	findInfo := bson.M{"method": "transfer", "$or": []bson.M{bson.M{"param.from": params.AccountName}, bson.M{"param.to": params.AccountName}}}
	if err := trxModule.Find(findInfo).Sort("-block_number").Skip(start).Limit(length).All(&transferList); err != nil {
		common.ResponseErr(c, "find personal transfers failed", err)
		return
	}

	for i := 0; i < len(transferList); i++ {
		t := transferList[i]
		tempT := module.ResTransfers{
			BlockNumber:   t.BlockNumber,
			From:          t.Param["from"].(string),
			To:            t.Param["to"].(string),
			Value:         t.Param["value"].(float64),
			CoinType:      "BTO",
			TransactionID: t.TransactionID,
			TimeStamp:     t.TimeStamp,
		}
		resTransferList = append(resTransferList, tempT)
	}

	totalCount, _ := trxModule.Find(findInfo).Count()
	res := module.ResPageList{
		Data:          resTransferList,
		TotalRecordes: totalCount,
	}

	common.ResponseSuccess(c, "find personal transfer success", res)
}
