package controller

import (
	"github.com/BlockExplorer/common"
	"github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
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
