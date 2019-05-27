package controller

import (
	"net/http"

	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// http://explorer.testnet.botfans.org:8080/api/historys/stake?account=testnameuser&page_num=0&

func StakeHistory(c *gin.Context) {
	account := c.Query("account")
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)
	trxModule := mongoIns.TransactionCollection()
	res := []bson.M{}
	if err := trxModule.Find(bson.M{"sender": account, "method": "stake"}).All(&res); err != nil {
		// common.ResponseErr(c, "stake history search error", err)
		c.JSON(http.StatusOK, gin.H{
			"data":    err,
			"message": "stake history search error",
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  res,
		"message": "find stake history success",
		"success": true,
	})
	// common.ResponseSuccess(c, "find stake history success", res)
}

func VoteHistory(c *gin.Context) {
	account := c.Query("account")
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)
	trxModule := mongoIns.TransactionCollection()
	res := []bson.M{}
	if err := trxModule.Find(bson.M{"sender": account, "method": "votedelegate"}).All(&res); err != nil {
		// common.ResponseErr(c, "stake history search error", err)
		c.JSON(http.StatusOK, gin.H{
			"data":    err,
			"message": "stake history search error",
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  res,
		"message": "find stake history success",
		"success": true,
	})
}
