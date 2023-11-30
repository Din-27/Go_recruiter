package handler

import (
	"fmt"
	"net/http"

	"github.com/Din-27/Go_recruiter/internal/models"
	"github.com/Din-27/Go_recruiter/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetAllLowongan(c *gin.Context) {

	var _data []models.ResLowonganPerusahaan

	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}

	result := db.Model(&models.LowonganPerusahaan{}).
		Select("dp.logo, p.nama, lowongan_perusahaans.*").
		Joins("join perusahaans p on p.id_company = lowongan_perusahaans.id_company").
		Joins("join detail_perusahaans dp on dp.id_company = lowongan_perusahaans.id_company").
		Scan(&_data)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": _data})
}

func GetDetailLowongan(c *gin.Context) {
	id_lowongan := c.Param("id_lowongan")
	var (
		lowongan       models.LowonganPerusahaan
		benefit        []models.BenefitLowonganPerusahaan
		requirement    []models.RequirementLowonganPerusahaan
		detailLowongan models.DetailLowongan
	)

	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}

	// arr
	getFoundation := db.Where("id_company = ? And id_lowongan=?", data.Id, id_lowongan).Find(&lowongan)
	if getFoundation.Error != nil {
		_resError(c, "server internal error", getFoundation.Error)
		return
	}
	if getFoundation.RowsAffected == 0 {
		_resError(c, "server internal error", _isErr("Data tidak di temukan !"))
		return
	}
	db.Where("id_lowongan = ?", id_lowongan).Find(&benefit)
	db.Where("id_lowongan = ?", id_lowongan).Find(&requirement)

	detailLowongan.IdLowongan = lowongan.IdLowongan
	detailLowongan.Title = lowongan.Title
	detailLowongan.Deskripsi = lowongan.Deskripsi
	detailLowongan.MinGaji = lowongan.MinGaji
	detailLowongan.MaxGaji = lowongan.MaxGaji
	detailLowongan.Poster = lowongan.Poster
	detailLowongan.DurasiLowongan = lowongan.DurasiLowongan
	detailLowongan.BenefitLowonganPerusahaan = benefit
	detailLowongan.RequirementLowonganPerusahaan = requirement

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": detailLowongan})
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
		Category:       bodyLowongan.Category,
		Deskripsi:      bodyLowongan.Deskripsi,
		MinGaji:        bodyLowongan.MinGaji,
		MaxGaji:        bodyLowongan.MaxGaji,
		Poster:         bodyLowongan.Poster,
		TipePekerjaan:  bodyLowongan.TipePekerjaan,
		LevelPekerjaan: bodyLowongan.LevelPekerjaan,
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

func GetAllLowonganCompany(c *gin.Context) {

	var (
		_data          []models.ResLowonganPerusahaan
		company        models.Perusahaan
		detail_company models.DetailPerusahaan
		lowongan       []models.LowonganPerusahaan
	)

	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "company" {
		_resError(c, "server internal error", _isErr("url ini untuk company !"))
		return
	}
	allLowongan := db.Where("id_company=?", data.Id).Find(&lowongan)
	if allLowongan.Error != nil {
		_resError(c, "server internal error", allLowongan.Error)
		return
	}
	for i := 0; i < len(lowongan); i++ {
		db.Where("id_company = ?", lowongan[i].Id).Take(&company)
		db.Where("id_company = ?", lowongan[i].Id).Take(&detail_company)

		_data = append(_data, models.ResLowonganPerusahaan{
			Id:             lowongan[i].Id,
			Logo:           detail_company.Logo,
			Nama:           company.Nama,
			IdLowongan:     lowongan[i].IdLowongan,
			Title:          lowongan[i].Title,
			Deskripsi:      lowongan[i].Deskripsi,
			MinGaji:        lowongan[i].MinGaji,
			MaxGaji:        lowongan[i].MaxGaji,
			Poster:         lowongan[i].Poster,
			DurasiLowongan: lowongan[i].DurasiLowongan,
		})
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": _data})
}

func GetAllFilterLowongan(c *gin.Context) {
	var (
		tipe_pekerjaan     []models.ResFilter
		level_pekerjaan    []models.ResFilter
		kategori_pekerjaan []models.ResFilter
	)
	getCategory := db.Model(&models.LowonganPerusahaan{}).Select("category as nama, COUNT(title) AS jumlah").Group("category").Find(&kategori_pekerjaan)
	if getCategory.Error != nil {
		_resError(c, "server internal error", getCategory.Error)
		return
	}
	getLevel := db.Model(&models.LowonganPerusahaan{}).Select("level_pekerjaan as nama, COUNT(title) AS jumlah").Group("level_pekerjaan").Find(&level_pekerjaan)
	if getLevel.Error != nil {
		_resError(c, "server internal error", getLevel.Error)
		return
	}
	getTipe := db.Model(&models.LowonganPerusahaan{}).Select("tipe_pekerjaan as nama, COUNT(title) AS jumlah").Group("tipe_pekerjaan").Find(&tipe_pekerjaan)
	if getTipe.Error != nil {
		_resError(c, "server internal error", getTipe.Error)
		return
	}

	fmt.Println(kategori_pekerjaan)
	result := map[string]interface{}{
		"tipe_pekerjaan":     tipe_pekerjaan,
		"level_pekerjaan":    level_pekerjaan,
		"kategori_pekerjaan": kategori_pekerjaan,
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": result})
}
