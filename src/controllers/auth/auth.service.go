package auth

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/helpers/models"
	"github.com/Din-27/Go_job/src/config"
	"github.com/Din-27/Go_job/src/controllers/Auth/schema"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
)

var (
	oneWeek     = 7 * 24 * time.Hour
	fiveMinutes = 15 * time.Minute
	db          = config.DBinit()
	_resError   = helpers.ResponseError
	_isErr      = helpers.ErrorReturn
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
		var user schema.User
		if err := c.ShouldBindJSON(&user); err != nil {
			_resError(c, "error", err)
			return
		}

		encodedHash, err := helpers.GenerateFromPassword(user.Password, p)
		if err != nil {
			_resError(c, "server internal error", err)
		}
		user.Password = encodedHash
		user.Role = role
		fmt.Println(user)

		result := db.Create(&user)
		if result.Error != nil {
			_resError(c, "server internal error", result.Error)
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": user})
		return
	}
	var company schema.User
	if err := c.ShouldBindJSON(&company); err != nil {
		_resError(c, "error", err)
		return
	}

	encodedHash, err := helpers.GenerateFromPassword(company.Password, p)
	if err != nil {
		_resError(c, "server internal error", err)
	}
	company.Password = encodedHash
	company.Role = role
	fmt.Println(company)

	result := db.Create(&company)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": company})
	return

}

func Login(c *gin.Context) {
	var (
		login   schema.Login
		user    schema.User
		company schema.Company
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
	encodedHash, err := helpers.GenerateFromPassword(login.Password, p)
	if err != nil {
		_resError(c, "server internal error", err)
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

	refresh_token, err := helpers.GenerateRefreshToken(user.Username, login.Email, role, oneWeek)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}

	access_token, err := helpers.GenerateAccessToken(user.Username, login.Email, role, fiveMinutes)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if role == "user" {
		results := schema.ResponseLogin{
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
	results := schema.ResponseLoginCompany{
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
	// Handle the refresh token request to issue a new access token.
	var newJsonToken paseto.JSONToken
	var newFooter string
	// Extract the refresh token from the request
	b, _ := hex.DecodeString(os.Getenv("PUBLIC_KEY"))
	publicKey := ed25519.PublicKey(b)
	var requestData struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := paseto.NewV2().Verify(requestData.RefreshToken, publicKey, &newJsonToken, &newFooter); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid refresh token %s", err)})
		return
	}
	username := newJsonToken.Get("username")
	email := newJsonToken.Get("email")
	role := newJsonToken.Get("role")

	// If the token is valid, generate a new access token forthe same user
	accessToken, err := helpers.GenerateAccessToken(username, email, role, fiveMinutes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}
	refresh_token, err := helpers.GenerateRefreshToken(username, email, role, oneWeek)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refresh_token})
}
