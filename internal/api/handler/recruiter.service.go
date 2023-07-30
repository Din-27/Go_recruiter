package handler

import (
	"net/http"

	"github.com/Din-27/Go_job/internal/models"
	"github.com/gin-gonic/gin"
)

func AddProfileCompany(c *gin.Context) {
	var company models.DetailPerusahaan
	if err := c.ShouldBindJSON(&company); err != nil {
		_resError(c, "error", err)
		return
	}
	result := db.Create(&company)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": company})
}

func AddLowongan(c *gin.Context) {
	var (
		bodyLowongan models.AddLowongan
	)
	if err := c.ShouldBindJSON(&bodyLowongan); err != nil {
		_resError(c, "error", err)
		return
	}
	lowongan := models.LowonganPerusahaan{
		Id:             bodyLowongan.Id,
		Title:          bodyLowongan.Title,
		Deskripsi:      bodyLowongan.Deskripsi,
		MinGaji:        bodyLowongan.MinGaji,
		MaxGaji:        bodyLowongan.MaxGaji,
		Poster:         bodyLowongan.Poster,
		DurasiLowongan: bodyLowongan.DurasiLowongan,
	}
	result := db.Create(&lowongan)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	for _, BenefitLowonganPerusahaan := range bodyLowongan.Benefit {
		result := db.Create(&BenefitLowonganPerusahaan) // Insert data ke tabel
		if result.Error != nil {
			_resError(c, "server internal error", result.Error)
			return
		}
	}
	for _, RequirementLowonganPerusahaan := range bodyLowongan.Requirement {
		result := db.Create(&RequirementLowonganPerusahaan) // Insert data ke tabel
		if result.Error != nil {
			_resError(c, "server internal error", result.Error)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "sukses membuat lowongan"})
}

func GetProfileCompany(c *gin.Context) {
	var company models.DetailPerusahaan
	if err := c.ShouldBindJSON(&company); err != nil {
		_resError(c, "error", err)
		return
	}
	result := db.Create(&company)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": company})
}
