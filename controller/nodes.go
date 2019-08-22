package controller

import (
	"log"

	"github.com/bottos-project/BlockExplorer/db"
	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"os"
	"io"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"image/png"
	"fmt"
	"crypto/md5"
	"time"
	"strconv"
	"encoding/hex"
	"path/filepath"
)

// //NodeProductList Node product list
// func NodeProductList(c *gin.Context) {
// 	var nodes []module.DBNodeSuper

// 	mongoIns, err := db.NewDBCollection()
// 	if err != nil {
// 		common.ResponseErr(c, "connect mongoDB error", err)
// 		return
// 	}

// 	defer db.CloseSession(mongoIns)

// 	nodeModule := mongoIns.NodeSuperCollection()
// 	if err := nodeModule.Find(bson.M{"delegate": bson.M{"$exists": true}}).Sort("-votes").All(&nodes); err != nil {
// 		log.Fatalf("product node list search failed: %v", err)
// 		common.ResponseErr(c, "product nodes search failed", err)
// 		return
// 	}

// 	common.ResponseSuccess(c, "product node list search success", nodes)
// }

func NodeServiceList(c *gin.Context) {
	var nodes []module.DBNodeService
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeServiceCollection()
	if err := nodeModule.Find(bson.M{}).All(&nodes); err != nil {
		common.ResponseErr(c, "service node list search failed", err)
		return
	}
	common.ResponseSuccess(c, "service node list search success", nodes)
}

// NodeServiceSummary get nodes summary
func NodeServiceSummary(c *gin.Context) {
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeServiceCollection()
	superNodeModule := mongoIns.NodeSuperCollection()
	serviceNodeCount, err := nodeModule.Find(bson.M{}).Count()
	if err != nil {
		common.ResponseErr(c, "service node list search failed", err)
		return
	}
	superNodeCount, err := superNodeModule.Find(bson.M{}).Count()
	if err != nil {
		common.ResponseErr(c, "super node list search failed", err)
		return
	}
	res := module.ResServiceNodeSummary{
		DelegateCount: superNodeCount,
		ServiceCount:  serviceNodeCount,
	}

	common.ResponseSuccess(c, "service node list search success", res)
}

func NodeSuperDetail(c *gin.Context) {
	var params module.ReqNodeSuperDetail
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

	nodeModule := mongoIns.NodeSuperCollection()
	accountModule := mongoIns.AccountCollection()
	//trxModule := mongoIns.TransactionCollection()
	blockModule := mongoIns.BlockCollection()
	res := bson.M{}
	account := bson.M{}
	// node info
	if err := nodeModule.Find(bson.M{"delegate": params.NodeName}).One(&res); err != nil {
		common.ResponseErr(c, "super node detail search failed", err)
		return
	}

	// account info
	accountModule.Find(bson.M{"account_name": params.NodeName}).One(&account)
	res["detail"] = account

	// blocks
	blockCount, _ := blockModule.Find(bson.M{"delegate": params.NodeName}).Count()
	res["produce_block_count"] = blockCount

	// voters
	/*voteCount, _ := trxModule.Find(bson.M{"method": "votedelegate", "param.delegate": params.NodeName}).Count()
	res["voters"] = voteCount*/

	voteCount, _ := accountModule.Find(bson.M{"vote.delegate": params.NodeName,"vote.votes":bson.M{"$ne":0}}).Count()
	res["voters"] = voteCount


	common.ResponseSuccess(c, "super node detail search success", res)
}

func NodeSummaryAuto(c *gin.Context) {
	var nodes module.ResNodeSuperSummary

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	nodeModule := mongoIns.NodeSuperCollection()
	count, err := nodeModule.Find(bson.M{"node_name": bson.M{"$exists": true}}).Count()
	if err != nil {
		log.Fatalf("product node list search failed: %v", err)
		common.ResponseErr(c, "product nodes search failed", err)
		return
	}
	nodes.DelegateCount = int32(count)

	accountModule := mongoIns.AccountCollection()

	voteAccountCount, err := accountModule.Find(bson.M{"vote.votes": bson.M{"$exists": true}}).Count()
	if err != nil {
		common.ResponseErr(c, "accounts count failed", err)
		return
	}
	nodes.VoteCustCount = int32(voteAccountCount)

	var node module.DBNodeSuper
	err = nodeModule.Find(bson.M{}).One(&node)
	if err != nil {
		log.Fatalf("product node search failed: %v", err)
		common.ResponseErr(c, "product node search failed", err)
		return
	}
	nodes.AllTransitVotes = node.AllTransitVotes
	common.ResponseSuccess(c, "product node summary search success", nodes)
}

//NodeProductList Node product list
func NodeProductList(c *gin.Context) {
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	selector := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "Accounts",//同一个数据库下等待被Join的集合。
				"localField":   "delegate",
				"foreignField": "account_name",
				"as":           "account_info",
			},
		},
		{
			"$replaceRoot": bson.M{ // 文档替换
				"newRoot": bson.M{
					"$mergeObjects": []interface{}{ // 文档合并
						bson.M{"$arrayElemAt": []interface{}{"$account_info", 0}},
						"$$ROOT",
					},
				},
			},
		},
		{ // 过滤有效字段
			"$project": bson.M{"account_info": 0},
		},
	}
	nodeModule := mongoIns.NodeSuperCollection()

	var res []bson.M

	if err := nodeModule.Pipe(selector).All(&res); err != nil {
		common.ResponseErr(c, "product nodes search failed", err)
		return
	}
	common.ResponseSuccess(c, "product node list search success", res)
}

func NodeVoterList(c *gin.Context) {
	var params module.ReqNodeSuperDetail
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
	data := []bson.M{}

	start := params.Start
	length := params.Length

	trxModule := mongoIns.TransactionCollection()
	if err := trxModule.Find(bson.M{"method": "votedelegate", "param.delegate": params.NodeName}).Sort("-_id").Skip(start).Limit(length).All(&data); err != nil {
		common.ResponseErr(c, "super node voters search failed", err)
		return
	}

	count, _ := trxModule.Find(bson.M{"method": "votedelegate", "param.delegate": params.NodeName}).Count()

	res := bson.M{
		"data":                 data,
		"iTotalDisplayRecords": count,
	}

	common.ResponseSuccess(c, "super node voters search success", res)
}

func NodeVoterHistory(c *gin.Context)  {
	var params module.ReqNodeSuperDetail
	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}
	coll,err :=db.NewDBCollection()
	if err !=nil{
		common.ResponseErr(c,"connect mongoDB error", err)
	}
	defer  db.CloseSession(coll)
	data := []bson.M{}

	start := params.Start
	length := params.Length

	collAccount := coll.AccountCollection()

	if err := collAccount.Find(bson.M{"vote.delegate": params.NodeName,"vote.votes":bson.M{"$ne":0}}).Sort("-_id").Skip(start).Limit(length).All(&data);err != nil{
		common.ResponseErr(c,"not found",err)
	}
	common.ResponseSuccess(c, "product node list search success", data)
}






func UpdateNodeInfo(c *gin.Context)  {
	var url string = "explorerpic.botfans.org"
	var path string  = "/var/www/explorerpic/"
	var avatar_width uint = 150
	var avatar_height uint
	var superNodeInfo bson.M

	node_name := c.Request.FormValue("node_name")
	if(node_name ==""){
		common.ResponseErr(c,"node_name miss",nil)
		return
	}

	_,e :=c.FormFile("avatar")

	if(e ==nil){
		fmt.Println("filefilefilefilefilefile")
		file, header, err := c.Request.FormFile("avatar") //image这个是uplaodify参数定义中的   'fileObjName':'image'
		if err != nil {
			common.ResponseErr(c,"params error",err)
			return
		}

		filename :=   header.Filename
		ext := filepath.Ext(header.Filename)
		//fmt.Println(file, err, filename)
		//创建文件

		hashmd5 := md5.New()
		timestamp :=  strconv.FormatInt(time.Now().Unix(),10)
		hashmd5.Write([]byte(filename+timestamp))
		mdname := hex.EncodeToString(hashmd5.Sum(nil))

		//os.Exit(222)
		out, err := os.Create(path+mdname+ext)
		//注意此处的 static/uploadfile/ 不是/static/uploadfile/
		if err != nil {
			log.Println(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Println(err)
		}
		open_file,_ := os.Open(path+mdname+ext)
		superNodeInfo = bson.M{
			"avatar":url+"/"+mdname+ext,
		}

		//缩略图

		//image.Decode(open_file)
		//png.Decode(open_file)
		temp_img,_,err := image.Decode(open_file)
		if err != nil{
			log.Println(err)

		}
		open_file.Close()

		thumb_avatar := resize.Resize(avatar_width,avatar_height,temp_img,resize.NearestNeighbor)

		out_thumb,err := os.Create(path+"thumb_"+mdname+ext)
		defer out_thumb.Close()
		png.Encode(out_thumb,thumb_avatar)
		if err !=nil{
			log.Println(err)

		}
		superNodeInfo = bson.M{
			"avatar":url+"/"+mdname+ext,
			"avatar_thumb":url+"/"+"thumb_"+mdname+ext,
		}
	}

	description := c.Request.FormValue("description")
	if(description !=""){
		if(superNodeInfo == nil){
			superNodeInfo =bson.M{"description":description}
		}else {
			superNodeInfo["description"] = description
		}

	}

	fmt.Printf("%+v",superNodeInfo)
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	NodeSuperModule := mongoIns.NodeSuperCollection()
	where := bson.M{"delegate":node_name}
	update := bson.M{"$set": superNodeInfo}


	errUpdate := NodeSuperModule.Update(where,update);if errUpdate != nil{
		fmt.Println(errUpdate)
	}
	//c.String(http.StatusCreated, "upload successful")
	common.ResponseSuccess(c, "upload successful", superNodeInfo)

}


//11000000000
//true
//a0123456789
//Beijing,China
//Bottos up up
//04ce4936e88d161951023eba9eedacd4277ccb83490aff50b56343baa5b212ec8836a5a097ea73fea6f9bc4e0b3926b60611ecb43af21130ea349f01b1d562f6cc