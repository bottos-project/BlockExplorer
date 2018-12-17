package controller

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	module "github.com/BlockExplorer/module"
	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	var user module.User
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
	}

	UserModule := module.UserCollection()
	err := UserModule.Insert(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":     1,
			"response": "insert failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"response": "ok",
	})
}

func Find(c *gin.Context) {
	var users []module.User
	UserModule := module.UserCollection()
	err := UserModule.Find(bson.M{}).All(&users)
	if err != nil {
		log.Println("find err")
		c.JSON(http.StatusOK, gin.H{
			"code":     "0",
			"response": "failed to find data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     "0",
		"response": users,
	})
}

func Update(c *gin.Context) {
	var user module.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":     "1",
			"response": "param invalid",
		})
		return
	}

	UserModule := module.UserCollection()
	err := UserModule.Update(bson.M{"username": user.UserName}, user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":     "1",
			"response": "update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     "0",
		"response": user,
	})
}

func Delete(c *gin.Context) {
	var user module.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":     "1",
			"response": "param invalid",
		})
		return
	}
	UserModule := module.UserCollection()
	err := UserModule.Remove(bson.M{"username": user.UserName})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":     "1",
			"response": "delete failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     "0",
		"response": "delete success",
	})
}
