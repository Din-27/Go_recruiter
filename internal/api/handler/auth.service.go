package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Din-27/Go_job/internal/config"
	"github.com/Din-27/Go_job/internal/models"
	"github.com/Din-27/Go_job/internal/utils"
	"github.com/gin-gonic/gin"
)

var (
	oneWeek     = 7 * 24 * time.Hour
	fiveMinutes = 15 * time.Minute
	db          = config.DBinit()
	_resError   = utils.ResponseError
	_isErr      = utils.ErrorReturn
	p           = &models.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
)

type _role struct {
	Name string
}

func RoleHandle(c *gin.Context) {
	result := []_role{
		{Name: "user"},
		{Name: "company"},
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": result})
}

func Register(c *gin.Context) {

	role := c.Param("role")
	if role == "user" {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			_resError(c, "error", err)
			return
		}
		user.Role = role
		err := utils.ValidateBody(user)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}
		encodedHash, err := utils.GenerateFromPassword(user.Password, p)
		if err != nil {
			_resError(c, "server internal error", err)
		}
		user.Password = encodedHash
		fmt.Println(user)

		// result := db.Create(&user)
		// if result.Error != nil {
		// 	_resError(c, "server internal error", result.Error)
		// 	return
		// }
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": user})
		return
	}
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		_resError(c, "error", err)
		return
	}
	company.Role = role
	err := utils.ValidateBody(company)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	encodedHash, err := utils.GenerateFromPassword(company.Password, p)
	if err != nil {
		_resError(c, "server internal error", err)
	}
	company.Password = encodedHash
	fmt.Println(company)

	// result := db.Create(&company)
	// if result.Error != nil {
	// 	_resError(c, "server internal error", result.Error)
	// 	return
	// }
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": company})
	return

}

func Login(c *gin.Context) {
	var (
		login   models.Login
		user    models.User
		company models.Company
	)
	role := c.Param("role")
	if err := c.ShouldBindJSON(&login); err != nil {
		_resError(c, "error", err)
		return
	}

	// if role == "user" {
	// 	result := db.Where("email = ?", login.Email).Take(&user)
	// 	if result.Error != nil {
	// 		_resError(c, "error", _isErr("Email tidak ditemukan !"))
	// 		return
	// 	}
	// }

	email := "test1@gmail.com"
	if email != login.Email {
		_resError(c, "error", _isErr("Email tidak ditemukan !"))
		return
	}
	encodedHash, err := utils.GenerateFromPassword(login.Password, p)
	if err != nil {
		_resError(c, "server internal error", err)
	}
	user.Password = encodedHash
	match, _err := utils.ComparePasswordAndHash(login.Password, user.Password)
	if _err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if !match {
		_resError(c, "error", _isErr("email atau Password salah !"))
		return
	}

	refresh_token, err := utils.GenerateRefreshToken(user.Username, login.Email, role, oneWeek)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	access_token, err := utils.GenerateAccessToken(user.Username, login.Email, role, fiveMinutes)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if role == "user" {
		results := models.ResponseLogin{
			Id:           user.Id,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Username:     user.Username,
			Email:        user.Email,
			Specialist:   user.Specialist,
			AccessToken:  access_token,
			RefreshToken: refresh_token,
		}
		c.JSON(http.StatusOK, gin.H{"value": results})
		return
	}
	results := models.ResponseLoginCompany{
		Id:           company.Id,
		Name:         company.Name,
		Email:        company.Email,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	c.JSON(http.StatusOK, gin.H{"value": results})
	return
}

func RefreshToken(c *gin.Context) {

	var requestData struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	result, err := utils.DecodedToken(requestData.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	// If the token is valid, generate a new access token forthe same user
	accessToken, err := utils.GenerateAccessToken(result.Username, result.Email, result.Role, fiveMinutes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}
	refresh_token, err := utils.GenerateRefreshToken(result.Username, result.Email, result.Role, oneWeek)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refresh_token})
}
