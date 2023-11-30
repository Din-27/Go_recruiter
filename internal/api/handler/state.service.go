package handler

import (
	"net/http"

	"github.com/Din-27/Go_recruiter/internal/utils"
	"github.com/gin-gonic/gin"
)

func ListProvince(c *gin.Context) {

	value, err := utils.FetchGetProvinsi()
	if err != nil {
		_resError(c, "error", err)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": value})
}

func ListKabupaten(c *gin.Context) {
	id_provinsi := c.Param("id_provinsi")
	value, err := utils.FetchGetKabupaten(id_provinsi)
	if err != nil {
		_resError(c, "error", err)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": value})
}

func ListKecamatan(c *gin.Context) {
	id_kabupaten := c.Param("id_kabupaten")
	value, err := utils.FetchGetKecamatan(id_kabupaten)
	if err != nil {
		_resError(c, "error", err)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": value})
}

func ListKelurahan(c *gin.Context) {
	id_kecamatan := c.Param("id_kecamatan")
	value, err := utils.FetchGetKecamatan(id_kecamatan)
	if err != nil {
		_resError(c, "error", err)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": value})
}
