package controller

import (
	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func StatisticTransaction(c *gin.Context) {
	var trxStatistics []module.ResStatisticTransaction
	sModule := module.StatisticCollection()
	if err := sModule.Find(bson.M{}).Limit(14).All(&trxStatistics); err != nil {
		common.ResponseErr(c, "transaction statistic failed", err)
		return
	}
	common.ResponseSuccess(c, "transaction statistic success", trxStatistics)
}

func StatisticAccount(c *gin.Context) {
	var accountStatistic []module.ResStatisticAccount
	sModule := module.StatisticCollection()
	if err := sModule.Find(bson.M{}).Limit(7).All(&accountStatistic); err != nil {
		common.ResponseErr(c, "account statistic failed", err)
		return
	}

	common.ResponseSuccess(c, "account statistic success", accountStatistic)
}
