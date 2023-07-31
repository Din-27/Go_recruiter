package handler

import (
	"net/http"

	"github.com/Din-27/Go_job/internal/models"
	"github.com/Din-27/Go_job/internal/utils"
	"github.com/gin-gonic/gin"
)

func AddProfileCompany(c *gin.Context) {
	var company models.DetailPerusahaan
	if err := c.ShouldBindJSON(&company); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "company" {
		_resError(c, "server internal error", _isErr("url ini untuk perusahaan !"))
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
	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "company" {
		_resError(c, "server internal error", _isErr("url ini untuk perusahaan !"))
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
	var (
		company     models.Perusahaan
		detail      models.DetailPerusahaan
		lowongan    []models.LowonganPerusahaan
		// benefit     []models.BenefitLowonganPerusahaan
		// requirement []models.RequirementLowonganPerusahaan
	)

	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "company" {
		_resError(c, "server internal error", _isErr("url ini untuk perusahaan !"))
		return
	}
	db.Where("email = ?", data.Email).Take(&company)
	db.Where("id_company = ?", company.Id).Take(&detail)
	// arr
	if err := db.Table("detail_perusahaans dp").Select("*").Joins("JOIN lowongan_perusahaans lp ON dp.id_company = lp.id_company").Scan(&lowongan).Error; err != nil {
		_resError(c, "server internal error", err)
		return
	}

	// db.Where("id_user = ?", company.Id).Find(&keahlianUser)
	// result := models.GetCompanyByIdResponse{
	// 	Id:             company.Id,
	// 	Nama:           company.Nama,
	// 	Alamat:         detail.Alamat,
	// 	Deskripsi:      detail.Deskripsi,
	// 	Bidang:         detail.Bidang,
	// 	Pencapaian:     detail.Pencapaian,
	// 	JumlahKaryawan: detail.JumlahKaryawan,
	// 	Website:        detail.Website,
	// 	Logo:           detail.Logo,
	// 	Background:     detail.Background,
	// }
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": lowongan})
}
