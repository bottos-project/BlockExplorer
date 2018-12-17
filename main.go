package main

import (
	"net/http"
	"time"

	"github.com/BlockExplorer/db"
	"github.com/BlockExplorer/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	db.InitMongoDB("test")
	routers.Routes(router)
	s.ListenAndServe()
}
