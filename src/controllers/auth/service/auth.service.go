package service

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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
	_isErr    = helpers.ErrorReturn
	p         = &models.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
)

func Register(c *gin.Context) {

	var user schema.User

	if err := c.ShouldBindJSON(&user); err != nil {
		_resError(c, "error", err)
		return
	}

	encodedHash, err := helpers.GenerateFromPassword(user.Password, p)
	if err != nil {
		log.Fatal(err)
	}
	encodedApi, err := helpers.GenerateFromPassword(user.Email+user.Password, p)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = encodedHash
	user.ApiToken = encodedApi
	fmt.Println(user)

	result := db.Create(&user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": user})
}

func Login(c *gin.Context) {
	var (
		login schema.Login
		user  schema.User
	)

	if err := c.ShouldBindJSON(&login); err != nil {
		_resError(c, "error", err)
		return
	}

	result := db.Where("email = ?", login.Email).Take(&user)
	if result.Error != nil {
		_resError(c, "error", _isErr("Email tidak ditemukan !"))
		return
	}

	// email := "test1@gmail.com"
	// if email != login.Email {
	// 	_resError(c, "error", _isErr("Email tidak ditemukan !"))
	// 	return
	// }
	encodedHash, err := helpers.GenerateFromPassword(login.Password, p)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = encodedHash
	match, _err := helpers.ComparePasswordAndHash(login.Password, user.Password)
	if _err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if !match {
		_resError(c, "error", _isErr("email atau Password salah !"))
		return
	}
	tokenMaker, err := helpers.NewPasetoMaker()
	if err != nil {
		log.Fatal(err)
	}
	token, _, err := tokenMaker.CreateToken(user.Id, user.Username, login.Email, user.ApiToken, time.Minute)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	results := schema.ResponseLogin{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Email:        user.Email,
		Specialist:   user.Specialist,
		RefreshToken: token,
	}
	c.JSON(http.StatusOK, gin.H{"value": results})
}

func RefreshToken(c *gin.Context) {
	var (
		token string
		user  schema.User
	)
	getToken := c.GetHeader("authorization")
	fields := strings.Fields(getToken)
	token = fields[1]

	value, _ := c.Get("authorization_payload")
	body, _ := value.(*helpers.Payload)
	err := helpers.Valid(body)
	if err != nil {
		tokenMaker, err := helpers.NewPasetoMaker()
		if err != nil {
			log.Fatal(err)
		}
		_token, _, _err := tokenMaker.CreateToken(body.ID, body.Username, body.Email, body.ApiToken, time.Minute)
		token = _token
		if _err != nil {
			_resError(c, "unauthorized", _err)
			return
		}
	}
	result := db.Where("api_token = ?", body.ApiToken).Take(&user)
	if result.Error != nil {
		_resError(c, "unauthorized", _isErr("session telah habis !"))
		return
	}
	fmt.Println(user)
	results := schema.ResponseLogin{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Email:        user.Email,
		Specialist:   user.Specialist,
		RefreshToken: token,
	}
	c.JSON(http.StatusOK, gin.H{"value": results})
}
