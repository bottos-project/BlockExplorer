package main

import (
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/bottos-project/BlockExplorer/routers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	router.Use(cors.Default())
	db.InitMongoDB("blockchainbowser")
	routers.Routes(router)
	s.ListenAndServe()
}
