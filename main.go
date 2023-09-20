package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// User/v1
	r.POST("/user", PostUser)
	//r.PUT("/user/:id/personal-data", PutUserPersonalData)

	dbConnection()

	r.Run()

}
