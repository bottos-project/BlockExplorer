package controller

import (
	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

//TrxTransactionList transaction list
func TrxTransactionList(c *gin.Context) {
	var transactions []module.ResTransactions
	var params module.ReqTransactionList
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

	res := module.ResPageList{
		Data:          &transactions,
		TotalRecordes: 0,
	}

	start := params.Start
	length := params.Length
	blockNum, _ := common.String2Int(params.BlockNum)

	findInfo := bson.M{}
	if blockNum != 0 {
		findInfo["block_number"] = int64(blockNum)
	}
	if params.AccountName != "" {
		findInfo["sender"] = params.AccountName
	}

	trxModule := mongoIns.TransactionCollection()

	trxCount, err := trxModule.Find(findInfo).Count()
	res.TotalRecordes = trxCount
	if err != nil {
		common.ResponseErr(c, "transactions search error", err)
		return
	}

	start, length = paging(start, length, trxCount)
	if length <= 0 {
		common.ResponseSuccess(c, "transfer list find success", res)
		return
	}
	if err := trxModule.Find(findInfo).Skip(start).Limit(length).All(&transactions); err != nil {
		common.ResponseErr(c, "transactions search error", err)
		return
	}

	common.ResponseSuccess(c, "transaction list find success", res)
}

//TrxTransferList 转账
func TrxTransferList(c *gin.Context) {
	var params module.ReqTransactionList
	var transferList []module.DBTransaction
	var resTransferList []module.ResTransfers

	res := module.ResPageList{
		Data:          &resTransferList,
		TotalRecordes: 0,
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

	blockNum, _ := common.String2Int(params.BlockNum)
	start := params.Start
	length := params.Length

	findInfo := bson.M{"method": "transfer"}
	if blockNum != 0 {
		findInfo["block_number"] = blockNum
	}

	transferModule := mongoIns.TransactionCollection()

	transferCount, err := transferModule.Find(findInfo).Count()
	if err != nil {
		common.ResponseErr(c, "transactions search error", err)
		return
	}
	res.TotalRecordes = transferCount

	start, length = paging(start, length, transferCount)
	if length <= 0 {
		common.ResponseSuccess(c, "transfer list find success", res)
		return
	}
	if err := transferModule.Find(findInfo).Skip(start).Limit(length).All(&transferList); err != nil {
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

	common.ResponseSuccess(c, "transfer list find success", res)
}

//TrxTransactionDetail 交易详情
func TrxTransactionDetail(c *gin.Context) {
	var params module.ReqGetTransaction
	var trx module.ResTransferDetail

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

	trxModule := mongoIns.TransactionCollection()
	if err := trxModule.Find(bson.M{"transaction_id": params.TransactionID}).One(&trx); err != nil {
		common.ResponseErr(c, "transaction detail find failed", err)
		return
	}

	trx.TradeStatus = "已确认"
	trx.ResultType = "trade"

	res := module.ResDataStruct{
		Data: trx,
	}

	common.ResponseSuccess(c, "transaction detail find success", res)
}

func TransferAccountHistory(c *gin.Context)  {
	var params module.ReqTransactionList
	var transferList []module.TransferHistory
	//var transferList []string
	//var resTransferList []module.ResTransfers
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}
	//start := params.Start
	//length := params.Length
	//res := module.ResPageList{
	//	Data:          &resTransferList,
	//	TotalRecordes: 0,
	//}
	//findInfo := bson.M{
	//	"method": "transfer",
	//	"param.from":params.AccountName,
	//	//"$or": []bson.M{
	//	//	bson.M{"param.from": params.AccountName},
	//	//	//bson.M{"param.to": params.AccountName},
	//	//},
	//}

	p := []bson.M{
		{"$match":bson.M{
			//"method":"transfer",
			"$or":[]bson.M{
				bson.M{"param.from": params.AccountName},
				//bson.M{"param.to": params.AccountName},
			},
		}},

		{
			"$group":bson.M{"_id":"$param.to"},
		},




		{
			"$sort":bson.M{"_id":1},
		},
	}



	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}
	defer db.CloseSession(mongoIns)
	trxModule := mongoIns.TransactionCollection()
	//trxCount, err := trxModule.Find(findInfo).Count()
	//if err != nil {
	//	common.ResponseErr(c, "transactions search error", err)
	//	return
	//}
	//res.TotalRecordes = trxCount
	//start, length = paging(start, length, trxCount)
	//if length <= 0 {
	//	common.ResponseSuccess(c, "transfer list find success", res)
	//	return
	//}
	//if err := trxModule.Find(findInfo).Skip(start).Limit(length).All(&transferList); err != nil {
	//	common.ResponseErr(c, "find personal transfers failed", err)
	//	return
	//}
	//trxModule.Find(findInfo).Sort("block_number").Distinct("param.to",&transferList)

	trxModule.Pipe(p).All(&transferList)
	common.ResponseSuccess(c, "find personal transfer success", transferList)
	fmt.Printf("%+v",transferList)



}




//TrxPersonalTransferList personal transfer list
func TrxPersonalTransferList(c *gin.Context) {
	var params module.ReqTransactionList
	var transferList []module.DBTransaction
	var resTransferList []module.ResTransfers

	res := module.ResPageList{
		Data:          &resTransferList,
		TotalRecordes: 0,
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

	start := params.Start
	length := params.Length

	trxModule := mongoIns.TransactionCollection()
	findInfo := bson.M{
		//
		//"$or":[]bson.M{
		//	bson.M{"method":"transfer"},
		//	bson.M{"method":"execmsignproposal"},
		//},
		"$or": []bson.M{
			bson.M{"param.from": params.AccountName},
			bson.M{"param.to": params.AccountName},
		},
	}

	trxCount, err := trxModule.Find(findInfo).Count()
	if err != nil {
		common.ResponseErr(c, "transactions search error", err)
		return
	}
	res.TotalRecordes = trxCount

	start, length = paging(start, length, trxCount)
	if length <= 0 {
		common.ResponseSuccess(c, "transfer list find success", res)
		return
	}
	if err := trxModule.Find(findInfo).Skip(start).Limit(length).All(&transferList); err != nil {
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
			Method:        t.Method,
		}
		resTransferList = append(resTransferList, tempT)
	}

	common.ResponseSuccess(c, "find personal transfer success", res)
}

//TrxPersonTransactionByMethod  personal transfer list by method
func TrxPersonTransactionByMethod(c *gin.Context) {
	var params module.ReqPersonTransactionByMethod
	var transferList []module.ResPersonalTransactionByMethod

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

	res := module.ResPageList{
		Data:          &transferList,
		TotalRecordes: 0,
	}

	start := params.Start
	length := params.Length

	if params.Method == "" {
		params.Method = "transfer"
	}

	trxModule := mongoIns.TransactionCollection()
	findInfo := bson.M{
		"method": params.Method,
		"$or": []bson.M{
			bson.M{"param.from": params.AccountName},
			bson.M{"param.to": params.AccountName},
		},
	}
	trxCount, err := trxModule.Find(findInfo).Count()
	if err != nil {
		common.ResponseErr(c, "transactions search error", err)
		return
	}

	res.TotalRecordes = trxCount

	start, length = paging(start, length, trxCount)
	if length <= 0 {
		// common.ResponseSuccess(c, "transfer list find success", res)
		common.ResponseSuccess(c, "transfer list find success", transferList)
		return
	}
	if err := trxModule.Find(findInfo).Skip(start).Limit(length).All(&transferList); err != nil {
		common.ResponseErr(c, "find personal transfers failed", err)
		return
	}

	// common.ResponseSuccess(c, "find personal transfers success", res)
	common.ResponseSuccess(c, "find personal transfers success", transferList)

}
