package controller

import (
	"fmt"

	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func AccountList(c *gin.Context) {
	var params module.ReqAccountList
	var accounts []module.ResAccountList

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	start := params.Start
	length := params.Length

	accountModule := module.AccountCollection()
	fmt.Println(accountModule)

	if err := accountModule.Find(bson.M{}).Skip(start).Limit(length).All(&accounts); err != nil {
		common.ResponseErr(c, "accounts find failed", err)
		return
	}
	totalCount, _ := accountModule.Count()

	res := module.ResPageList{
		Data:          accounts,
		TotalRecordes: totalCount,
	}

	common.ResponseSuccess(c, "accounts find success", res)
}

func AccountDetail(c *gin.Context) {
	var params module.ReqAccountInfo
	var account module.ResAccountDetail

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	accountModule := module.AccountCollection()
	if err := accountModule.Find(bson.M{"account_name": params.AccountName}).One(&account); err != nil {
		common.ResponseErr(c, "account find failed", err)
		return
	}

	// transactions
	trxModule := module.TransactionCollection()
	// transaction out count
	out := bson.M{"method": "transfer", "param.from": params.AccountName}
	send, _ := trxModule.Find(out).Count()
	// transaction in
	in := bson.M{"method": "transfer", "param.to": params.AccountName}
	receiveCount, _ := trxModule.Find(in).Count()

	// set account info
	account.SendCount = send
	account.ReceiveCount = receiveCount
	account.TradeCount = send
	res := module.ResDataStruct{
		Data: account,
	}
	common.ResponseSuccess(c, "account find success", res)
}

// 查询账户实时总数
func AccountTotal(c *gin.Context) {
	accountModule := module.AccountCollection()
	totalCount, err := accountModule.Count()
	if err != nil {
		common.ResponseErr(c, "account total search failed", err)
		return
	}

	type total struct {
		Count int `json:"rtCustCount"`
	}

	res := total{
		Count: totalCount,
	}
	common.ResponseSuccess(c, "success", res)
}
