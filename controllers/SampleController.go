package controllers

import (
	"GoBP/model/user"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"strconv"

)



func Ping(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PingWithAuth(c *gin.Context){
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"text":     "Hello World.",
	})
}

func Example(c *gin.Context){
	var id,_ = strconv.Atoi(c.Query("id"))
	var userData = user.GetUser(id)
	c.JSON(200, userData)
}