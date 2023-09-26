package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicus101/scaling-giggle/pkg/user"
)

// POST /user - create user - returns 200 with json {user_id: 13}
func PostUser(c *gin.Context) {
	log.Println("Creating User")

	var createUserCommand user.CreateCommand
	if err := c.BindJSON(&createUserCommand); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId, err := user.CreateUser(c, createUserCommand)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
	})
}

func PutPersonalData(c *gin.Context) {
	log.Println("Adding personal data")

	var AddPersonalDataCommand user.AddPersonalDataCommand
	if err := c.BindJSON(&AddPersonalDataCommand); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := user.AddPersonalData(c, AddPersonalDataCommand)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// PUT /user/:id/personal-data - akceptuje json - zwraca 204 czyli no-data
//func PutUserPersonalData

// c.BindJson :p https://pkg.go.dev/github.com/gin-gonic/gin#Context.BindJSON

func LogInTest(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")

	response, err := user.Auth(c.Request.Context(), login, password)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if response.Id == 0 {
		err := fmt.Errorf("invalid user or password")
		c.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
		// wyślij użytkownikowi 401 z {"error":"invalid user or password"}
	}

	c.JSON(http.StatusOK, gin.H{
		"user": response,
	})
}
