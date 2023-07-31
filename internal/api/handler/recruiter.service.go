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
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "company" {
		_resError(c, "server internal error", _isErr("url ini untuk perusahaan !"))
		return
	}
	company.Id = data.Id
	result := db.Create(&company)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "sukses memperbarui detail profile"})
}

func AddLowongan(c *gin.Context) {
	var bodyLowongan models.AddLowongan

	if err := c.ShouldBindJSON(&bodyLowongan); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "company" {
		_resError(c, "server internal error", _isErr("url ini untuk perusahaan !"))
		return
	}

	lowongan := models.LowonganPerusahaan{
		Id:             data.Id,
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
		BenefitLowonganPerusahaan.IdLowongan = lowongan.IdLowongan
		result := db.Create(&BenefitLowonganPerusahaan) // Insert data ke tabel
		if result.Error != nil {
			_resError(c, "server internal error", result.Error)
			return
		}
	}
	for _, RequirementLowonganPerusahaan := range bodyLowongan.Requirement {
		RequirementLowonganPerusahaan.IdLowongan = lowongan.IdLowongan
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
		company        models.Perusahaan
		detail         models.DetailPerusahaan
		lowongan       []models.LowonganPerusahaan
		benefit        models.BenefitLowonganPerusahaan
		requirement    models.RequirementLowonganPerusahaan
		detailLowongan []models.DetailLowongan
	)

	data, err := utils.DecodedTokenBearer(c, db)
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
	db.Where("id_company = ?", company.Id).Find(&lowongan)
	for i := 0; i < len(lowongan); i++ {
		detailLowongan = append(detailLowongan, models.DetailLowongan{
			IdLowongan:     lowongan[i].IdLowongan,
			Title:          lowongan[i].Title,
			Deskripsi:      lowongan[i].Deskripsi,
			MinGaji:        lowongan[i].MinGaji,
			MaxGaji:        lowongan[i].MaxGaji,
			Poster:         lowongan[i].Poster,
			DurasiLowongan: lowongan[i].DurasiLowongan,
		})
		db.Where("id_lowongan = ?", lowongan[i].IdLowongan).Find(&benefit)
		db.Where("id_lowongan = ?", lowongan[i].IdLowongan).Find(&requirement)
		detailLowongan[i].BenefitLowonganPerusahaan = append(detailLowongan[i].BenefitLowonganPerusahaan, benefit)
		detailLowongan[i].RequirementLowonganPerusahaan = append(detailLowongan[i].RequirementLowonganPerusahaan, requirement)
	}

	result := models.GetCompanyByIdResponse{
		Id:             company.Id,
		Nama:           company.Nama,
		Alamat:         detail.Alamat,
		Deskripsi:      detail.Deskripsi,
		Bidang:         detail.Bidang,
		Pencapaian:     detail.Pencapaian,
		JumlahKaryawan: detail.JumlahKaryawan,
		Website:        detail.Website,
		Logo:           detail.Logo,
		Background:     detail.Background,
		DetailLowongan: detailLowongan,
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": result})
}
