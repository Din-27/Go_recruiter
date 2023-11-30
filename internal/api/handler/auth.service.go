package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Din-27/Go_recruiter/internal/config"
	"github.com/Din-27/Go_recruiter/internal/models"
	"github.com/Din-27/Go_recruiter/internal/utils"
	"github.com/gin-gonic/gin"
)

var (
	oneWeek   = 7 * 24 * time.Hour
	oneDay    = 1 * 24 * time.Hour
	db        = config.DBinit()
	_resError = utils.ResponseError
	_isErr    = utils.ErrorReturn
	p         = &models.Params{
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
	result := []map[string]interface{}{
		{"nama": "user"},
		{"nama": "company"},
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
		err := utils.ValidateBody(user)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}
		checkEmail := db.Where("email = ?", user.Email).Take(&user)
		if checkEmail.Error == nil {
			_resError(c, "error", _isErr("Email sudah terdaftar !"))
			return
		}
		encodedHash, err := utils.GenerateFromPassword(user.Password, p)
		if err != nil {
			_resError(c, "server internal error", err)
		}
		user.Password = encodedHash
		fmt.Println(user)

		result := db.Create(&user)
		if result.Error != nil {
			_resError(c, "server internal error", result.Error)
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "sukses register"})
		return
	} else {
		var company models.Perusahaan
		if err := c.ShouldBindJSON(&company); err != nil {
			_resError(c, "error", err)
			return
		}
		err := utils.ValidateBody(company)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}
		checkEmail := db.Where("email = ?", company.Email).Take(&company)
		if checkEmail.Error == nil {
			_resError(c, "error", _isErr("Email sudah terdaftar !"))
			return
		}
		encodedHash, err := utils.GenerateFromPassword(company.Password, p)
		if err != nil {
			_resError(c, "server internal error", err)
		}
		company.Password = encodedHash
		fmt.Println(company)

		result := db.Create(&company)
		if result.Error != nil {
			_resError(c, "server internal error", result.Error)
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "sukses register"})
		return
	}
}

func Login(c *gin.Context) {
	var (
		login   models.Login
		user    models.User
		company models.Perusahaan
	)
	role := c.Param("role")
	if err := c.ShouldBindJSON(&login); err != nil {
		_resError(c, "error", err)
		return
	}

	if role == "user" {
		result := db.Where("email = ?", login.Email).Take(&user)
		if result.Error != nil {
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
		access_token, err := utils.GenerateAccessToken(user.Username, login.Email, role, oneDay)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}
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
	} else {
		result := db.Where("email = ?", login.Email).Take(&company)
		if result.Error != nil {
			_resError(c, "error", _isErr("Email tidak ditemukan !"))
			return
		}
		encodedHash, err := utils.GenerateFromPassword(login.Password, p)
		if err != nil {
			_resError(c, "server internal error", err)
		}
		company.Password = encodedHash
		match, _err := utils.ComparePasswordAndHash(login.Password, company.Password)
		if _err != nil {
			_resError(c, "server internal error", err)
			return
		}
		if !match {
			_resError(c, "error", _isErr("email atau Password salah !"))
			return
		}

		refresh_token, err := utils.GenerateRefreshToken(company.Nama, login.Email, role, oneWeek)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}
		access_token, err := utils.GenerateAccessToken(company.Nama, login.Email, role, oneDay)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}
		results := models.ResponseLoginCompany{
			Id:           company.Id,
			Name:         company.Nama,
			Email:        company.Email,
			AccessToken:  access_token,
			RefreshToken: refresh_token,
		}
		c.JSON(http.StatusOK, gin.H{"value": results})
		return
	}

	// email := "test1@gmail.com"
	// if email != login.Email {
	// 	_resError(c, "error", _isErr("Email tidak ditemukan !"))
	// 	return
	// }

}

func RefreshToken(c *gin.Context) {

	var (
		user    models.User
		company models.Perusahaan
	)

	authorizationHeader := c.GetHeader("authorization")

	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != "bearer" {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	refreshToken := fields[1]
	result, err := utils.RefreshDecodedToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	if result.Role == "user" {
		checkUser := db.Where("email = ?", result.Email).Take(&user)
		if checkUser.Error != nil {
			_resError(c, "error", _isErr("Failed token invalid !"))
			return
		}
		// If the token is valid, generate a new access token forthe same user
		accessToken, err := utils.GenerateAccessToken(result.Username, result.Email, result.Role, oneDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
			return
		}
		refresh_token, err := utils.GenerateRefreshToken(result.Username, result.Email, result.Role, oneWeek)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}

		results := models.ResponseLogin{
			Id:           user.Id,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Username:     user.Username,
			Email:        user.Email,
			Specialist:   user.Specialist,
			AccessToken:  accessToken,
			RefreshToken: refresh_token,
		}
		c.JSON(http.StatusOK, gin.H{"value": results})
	} else {
		checkUser := db.Where("email = ?", result.Email).Take(&company)
		if checkUser.Error != nil {
			_resError(c, "error", _isErr("Failed token invalid !"))
			return
		}
		// If the token is valid, generate a new access token forthe same user
		accessToken, err := utils.GenerateAccessToken(result.Username, result.Email, result.Role, oneDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
			return
		}
		refresh_token, err := utils.GenerateRefreshToken(result.Username, result.Email, result.Role, oneWeek)
		if err != nil {
			_resError(c, "server internal error", err)
			return
		}

		results := models.ResponseLoginCompany{
			Id:           company.Id,
			Name:         company.Nama,
			Email:        company.Email,
			AccessToken:  accessToken,
			RefreshToken: refresh_token,
		}

		c.JSON(http.StatusOK, gin.H{"value": results})
	}

}
