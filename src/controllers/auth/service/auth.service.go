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
	db     = config.DBinit()
	resErr = helpers.ResponseError
	resSuc = helpers.ResponseSukses
)

func Register(c *gin.Context) {

	var user schema.Register

	if err := c.ShouldBindJSON(&user); err != nil {
		resErr(c, "error", err)
		return
	}

	p := &models.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	encodedHash, err := helpers.GenerateFromPassword("test", p)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = encodedHash
	result := db.Create(user)
	if result.Error != nil {
		resErr(c, "status internal error", result.Error)
		return
	}
	resSuc(c, user)
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
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.
			ErrorResponse(errors.
				New("Email tidak ditemukan !")))
		return
	}

	match, _err := helpers.ComparePasswordAndHash(login.Password, user.Password)
	if _err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}
	if match != true {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse(errors.New("Email atau Password salah !")))
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
