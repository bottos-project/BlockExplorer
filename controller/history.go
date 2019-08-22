package controller

import (
	"net/http"

	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"github.com/bottos-project/BlockExplorer/module"
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
	nodeModule := mongoIns.NodeSuperCollection()

	res := []bson.M{}
	nodeInfo := module.NodeSuper{}
	//value := bson.M{}
	if err := trxModule.Find(
		bson.M{
		"sender":account,
		"$or": []bson.M{
			bson.M{"method": "votedelegate"},
			bson.M{"method": "stake","param.target":"vote"},
			bson.M{"method": "unstake","param.source":"vote"}}}).Sort("_id").All(&res); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    err,
			"message": "stake history search error",
			"success": false,
		})
		return
	}
	for k,v :=range res{
		//data[k] =v


			vv := v["param"].(bson.M)
			nodeModule.Find(bson.M{"delegate":vv["delegate"]}).One(&nodeInfo)

		res[k]["node_info"] = nodeInfo


	}
	//fmt.Printf("%+v",data)

	//os.Exit(801)
	c.JSON(http.StatusOK, gin.H{
		"result":  res,
		"message": "find stake history success",
		"success": true,
	})
}










