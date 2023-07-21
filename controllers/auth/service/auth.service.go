package service

import (
	"fmt"
	"github.com/Din-27/Go_job/controllers/auth/dto"
	"github.com/Din-27/Go_job/middlewares/tokenpaseto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *Repository {
	return &Repository{db}
}

func errorResponse(err error) *gin.H {
	return &gin.H{"message": err.Error()}
}

func (d *Repository) Register(c *gin.Context) {

	var (
		user dto.UserDto
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result := d.db.Create(user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(result.Error))
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (d *Repository) Login(c *gin.Context) {

	var (
		user dto.LoginDto
	)
	tokenMaker, err := tokenpaseto.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result := d.db.First(&user).Where("email = ?", user.Email)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "email tidak ditemukan !"})
		return
	}
	fmt.Println(user)
	token, payload, err := tokenMaker.CreateToken(1, user.Username, time.Minute)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": token, "payload": payload})
}

func (d *Repository) Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"msg": "sukses"})
}
