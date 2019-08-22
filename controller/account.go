package controller

import (
	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"fmt"
)

func AccountList(c *gin.Context) {
	var params module.ReqAccountList
	var accounts []module.ResAccountList

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

	start := params.Start -4
	length := params.Length
	order := params.Order
	by := params.By
	if(order == ""){
		order = "total_balance"

	}
	if(by == ""){
		by = ""
	}
	fmt.Println(start)
	fmt.Println(length)
	accountModule := mongoIns.AccountCollection()
	totalCount, _ := accountModule.Count()
	start, length = paging(start, length, totalCount)
	//fmt.Println(totalCount)
	//fmt.Println(start)
	//fmt.Println(length)
	if length <= 0 {
		common.ResponseSuccess(c, "accounts find success", module.ResPageList{TotalRecordes: totalCount})
		return
	}

	fmt.Println(by+order)
	if err := accountModule.Find(bson.M{}).Skip(start).Limit(length).Sort(by+order).All(&accounts); err != nil {
		common.ResponseErr(c, "accounts fdfind failed", err)
		return
	}

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

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	accountModule := mongoIns.AccountCollection()
	if err := accountModule.Find(bson.M{"account_name": params.AccountName}).One(&account); err != nil {
		common.ResponseErr(c, "account find failed", err)
		return
	}

	// transactions
	trxModule := mongoIns.TransactionCollection()
	// transaction out count
	out := bson.M{"method": "transfer", "param.from": params.AccountName}
	send, _ := trxModule.Find(out).Count()
	// transaction in
	in := bson.M{"method": "transfer", "param.to": params.AccountName}
	receiveCount, _ := trxModule.Find(in).Count()

	sender := bson.M{"sender": params.AccountName}
	trxCount, _ := trxModule.Find(sender).Count()
	// set account info
	account.SendCount = send
	account.ReceiveCount = receiveCount
	account.TradeCount = trxCount
	res := module.ResDataStruct{
		Data: account,
	}
	common.ResponseSuccess(c, "account find success", res)
}

// 查询账户实时总数
func AccountTotal(c *gin.Context) {
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)
	accountModule := mongoIns.AccountCollection()
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

func Msignaccount(c *gin.Context) {
	var params module.DBMsignAccountModule
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}
	// res := []bson.M{}
	var res []module.AuthorAccountModule
	mongoIns, _ := db.NewDBCollection()
	defer db.CloseSession(mongoIns)
	msignAccountModule := mongoIns.MsignaccountCollection()
	if error := msignAccountModule.Find(bson.M{"param.authority.author_account": params.AuthorAccount}).All(&res); error != nil {
		common.ResponseErr(c, "account find failed", error)
		return
	}

	common.ResponseSuccess(c, "success", res)
}




func MsignProposal(c *gin.Context) {
	var params module.DBMsignAccountModule
	//fmt.Println(params.AuthorAccount)
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	mongoIns, _ := db.NewDBCollection()
	defer db.CloseSession(mongoIns)
	//fmt.Println(bson.M{"param.authority.author_account": params.AuthorAccount})
	var authorRes []module.AuthorAccountModule
	msignAccountModule := mongoIns.MsignaccountCollection()
	if error := msignAccountModule.Find(bson.M{"param.authority.author_account": params.AuthorAccount}).All(&authorRes); error != nil {
		common.ResponseErr(c, "account find failed", error)
		return
	}
	//fmt.Println(authorRes)
	var authAccounts []string
	for index := 0; index < len(authorRes); index++ {
		a := authorRes[index]
		authAccounts = append(authAccounts, a.Param.Account)
	}

	//var res []module.DBTransactions
	res := []bson.M{};
	transactionModule := mongoIns.TransactionCollection()
	if error := transactionModule.Find(bson.M{"method": "pushmsignproposal", "param.account": bson.M{"$in": authAccounts}}).All(&res); error != nil {
		// if error := transactionModule.Find(bson.M{"method": "pushmsignproposal"}).All(&res); error != nil {
		common.ResponseErr(c, "account find failed", error)
		return
	}
	//transfer结构
	for _,v := range res{
		for key,val := range v{
			if key =="param"{
				for kk,vv := range val.(bson.M){
					if kk == "transfer"{
						data := make(map[string]string)
						if err:=json.Unmarshal([]byte(vv.(string)),&data);err !=nil{
							//fmt.Println(err)
							//return
						}
						 val.(bson.M)[kk] = data
					}
				}
			}
		}
	}

	common.ResponseSuccess(c, "success", res)
}

//func MsignProposalList(c *gin.Context)  {
//	var params module.DBMsignAccountModule
//	if err := c.BindJSON(&params); err != nil{
//		common.ResponseErr(c, "params error", err)
//		return
//	}
//	mongoIns,_ := db.NewDBCollection()
//	defer db.CloseSession(mongoIns)
//	var res []bson.M
//	transaction := mongoIns.TransactionCollection()
//	error := transaction.Find(bson.M{
//		"$or":[]bson.M{
//			bson.M{"method": "execmsignproposal","param.from":params.AuthorAccount},
//			bson.M{"method": "pushmsignproposal","param.account":params.AuthorAccount},
//		},
//
//	}).Sort("_id").All(&res)
//	if error != nil{
//		common.ResponseErr(c, "account find failed", error)
//		return
//	}
//	common.ResponseSuccess(c,"find success",res)
//}

func MsignProposalList(c *gin.Context)  {
	var params module.DBMsignAccountModule
	if err := c.BindJSON(&params);err != nil{
		common.ResponseErr(c,"params error",err)
		return
	}
	mongoIns,_ := db.NewDBCollection()
	defer db.CloseSession(mongoIns)
	var Pushres []module.PushProposal
	TransferCollection := mongoIns.TransactionCollection()
	err := TransferCollection.Find(bson.M{"method":"pushmsignproposal","param.account":params.AuthorAccount}).All(&Pushres)
	if err != nil{
		common.ResponseErr(c,"no proposal",err)
		return
	}

	var Execres module.ExecProposal
	for k,v := range Pushres{
		err := TransferCollection.Find(bson.M{
			"$or":[]bson.M{
				bson.M{"method":"execmsignproposal"},
				bson.M{"method":"cancelmsignproposal"},
			},
			"param.proposal":v.Param.Proposal,
			}).One(&Execres)
		if err == nil{

			Pushres[k].Param.From = Execres.Param.From
			Pushres[k].Param.To = Execres.Param.To
			Pushres[k].Param.Value = Execres.Param.Value
			Pushres[k].Param.Memo = Execres.Param.Memo
			Pushres[k].Res = Execres.Method
		}
	}
	common.ResponseSuccess(c,"find success",Pushres)
}




func GetPubAccount(c *gin.Context)  {
	var params module.PubKey
	if err := c.BindJSON(&params); err != nil{
		common.ResponseErr(c, "params error", err)
		return
	}
	mongoIns,_ := db.NewDBCollection()
	defer db.CloseSession(mongoIns)
	var res bson.M
	AccountCollection := mongoIns.AccountCollection()
	error := AccountCollection.Find(bson.M{
		"pubkey":params.Pkey,
	}).One(&res)
	if error != nil{
		common.ResponseErr(c, "account find failed", error)
		return
	}
	common.ResponseSuccess(c,"find success",res)
}