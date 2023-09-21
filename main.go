package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicus101/scaling-giggle/config"
	"github.com/nicus101/scaling-giggle/db"
)

func main() {
	log.Println("Service starting...")

	// Make infrastructure
	err := config.Load()
	if err != nil {
		log.Fatal("Loading configuration failed:", err)
	}

	err = db.InitConnection()
	if err != nil {
		log.Fatal("Connecting to db failed:", err)
	}
	defer db.Close()

	// setup http server
	r := gin.Default()

	// User/v1
	r.POST("/user", PostUser)
	//r.PUT("/user/:id/personal-data", PutUserPersonalData)
	r.GET("/error", func(c *gin.Context) {
		err := fmt.Errorf("zjem Ci Kotka")
		c.Error(err)
		c.JSON(http.StatusTeapot, gin.H{
			"error": err.Error(),
		})
	})

	// execute
	r.Run()
}
