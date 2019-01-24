package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bottos-project/BlockExplorer/routers"

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
	routers.Routes(router)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Linsten and Serve err: %v", err)
	}
}
