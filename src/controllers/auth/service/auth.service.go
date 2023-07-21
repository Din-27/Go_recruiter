package service

import (
	"log"
	"net/http"
	"time"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/src/config"
	"github.com/Din-27/Go_job/src/controllers/auth/schema"
	"github.com/gin-gonic/gin"
)

var db = config.DBinit()

func Register(c *gin.Context) {

	var user schema.Register

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}
	result := db.Create(user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ErrorResponse(result.Error))
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}


func Login(c *gin.Context) {
	var (
		login schema.Login
		user  schema.User
	)
	tokenMaker, err := helpers.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	result := db.First(&user).Where("email = ?", user.Email)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "email tidak ditemukan !"})
		return
	}
	token, payload, err := tokenMaker.CreateToken(user.Id, user.Username, time.Minute)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": token, "payload": payload})
}

func Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"msg": "sukses"})
}
