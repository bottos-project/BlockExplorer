package controller

import (
	"fmt"

	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// transaction statistics
func StatisticTransaction(c *gin.Context) {
	var statics []module.ResStatisticTransaction
	staticItem := module.ResStatisticTransaction{
		DailyTransCount: 5,
		TradeDate:       "2018-12-20",
	}

	statics = append(statics, staticItem)
	type summaryTradeListStrct struct {
		SummaryTradeList interface{} `json:"summaryTradeList"`
	}

	sumary := summaryTradeListStrct{
		SummaryTradeList: statics,
	}

	res := module.ResSuccess{
		Data:    sumary,
		Message: "account statistic search success",
		Success: true,
	}
	common.ResponseSuccess(c, "transaction statistic search success", res)
}

// account statistics
func StatisticAccount(c *gin.Context) {
	var statics []module.ResStatisticAccount
	staticItem := module.ResStatisticAccount{
		DailyRegCount: 36,
		RegDate:       "2018-09-22",
		TotalRegCount: 0,
	}
	statics = append(statics, staticItem)
	type regSummaryListStruct struct {
		RegSummaryList interface{} `json:"regSummaryList"`
	}

	regList := regSummaryListStruct{
		RegSummaryList: statics,
	}

	res := module.ResSuccess{
		Data:    regList,
		Message: "account statistic search success",
		Success: true,
	}

	common.ResponseSuccess(c, "account statistic search success", res)
}

func TestStatisticAccount(c *gin.Context) {
	var accounts []module.DBAccounts
	accountModule := module.AccountCollection()
	m := []bson.M{
		{"$match": bson.M{"timestamp": bson.M{"$gte": 1545389205, "$lte": 1545389205}}},
	}

	if err := accountModule.Pipe(m).All(&accounts); err != nil {
		fmt.Println("TestStatisticAccount failed")
		fmt.Println(err)
		common.ResponseErr(c, "statistic account failed", err)
		return
	}
	common.ResponseSuccess(c, "statistic account success", accounts)
}
