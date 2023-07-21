package service

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/helpers/models"
	"github.com/Din-27/Go_job/src/config"
	"github.com/Din-27/Go_job/src/controllers/auth/schema"
	"github.com/gin-gonic/gin"
)

var (
	db        = config.DBinit()
	_resError = helpers.ResponseError
	p         = &models.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
)

func Register(c *gin.Context) {

	var user schema.Register

	if err := c.ShouldBindJSON(&user); err != nil {
		_resError(c, "error", err)
		return
	}

	encodedHash, err := helpers.GenerateFromPassword("test", p)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = encodedHash
	result := db.Create(user)
	if result.Error != nil {
		_resError(c, "status internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": user})
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
		_resError(c, "error", err)
		return
	}

	// result := db.First(&user).Where("email = ? or username", login.Email, login.Username)
	// if result.Error != nil {
	// 	_resError(c, "error", errors.New("Email tidak ditemukan !"))
	// 	return
	// }

	email := "test1@gmail.com"
	if email != login.Email {
		_resError(c, "error", errors.New("Email tidak ditemukan !"))
		return
	}
	encodedHash, err := helpers.GenerateFromPassword("test123", p)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = encodedHash
	match, _err := helpers.ComparePasswordAndHash(login.Password, user.Password)
	if _err != nil {
		_resError(c, "status internal error", err)
		return
	}
	if match != true {
		_resError(c, "error", errors.New("Email atau Password salah !"))
		return
	}

	token, payload, err := tokenMaker.CreateToken(user.Id, login.Username, time.Minute)
	if err != nil {
		_resError(c, "status internal error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": token, "payload": payload})
}

func Test(c *gin.Context) {
	value, _ := c.Get("authorization_payload")
	c.JSON(http.StatusOK, gin.H{"msg": value})
}
