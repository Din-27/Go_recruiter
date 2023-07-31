package handler

import (
	"fmt"
	"net/http"

	"github.com/Din-27/Go_job/internal/models"
	"github.com/Din-27/Go_job/internal/utils"
	"github.com/gin-gonic/gin"
)

func AddUserDetail(c *gin.Context) {
	var user_detail models.DetailUser

	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}

	user_detail.Id = data.Id
	result := db.Create(&user_detail)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan detail diri anda"})
}

func AddUserPendidikanFormal(c *gin.Context) {
	var formal_user models.PendidikanFormalUser

	if err := c.ShouldBindJSON(&formal_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	formal_user.Id = data.Id
	result := db.Create(&formal_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan pendidikan formal anda"})
}

func AddUserPendidikanNonFormal(c *gin.Context) {
	var non_formal_user models.PendidikanNonFormalUser

	if err := c.ShouldBindJSON(&non_formal_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	non_formal_user.Id = data.Id
	result := db.Create(&non_formal_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan pendidikan non formal anda"})
}

func AddUserPengalaman(c *gin.Context) {
	var pengalaman_user models.PengalamanUser

	if err := c.ShouldBindJSON(&pengalaman_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	pengalaman_user.Id = data.Id
	result := db.Create(&pengalaman_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan pengalaman anda"})
}

func AddUserKeahlian(c *gin.Context) {
	var keahlian_user models.KeahlianUsers

	if err := c.ShouldBindJSON(&keahlian_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}

	keahlian_user.Id = data.Id
	result := db.Create(&keahlian_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan keahlian anda"})
}

func ApplyLamaranUser(c *gin.Context) {
	var apply models.ApplyLamaranUser

	if err := c.ShouldBindJSON(&apply); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}

	apply.IdUser = data.Id
	checkLamaran := db.Where("id_user = ? and id_company = ?", apply.IdUser, apply.IdCompany).Find(&apply)
	if checkLamaran.RowsAffected == 1 {
		_resError(c, "error", _isErr("Anda sudah melamar ke perusahaan ini !"))
		return
	}
	result := db.Create(&apply)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses apply lamaran"})
}

func GetUserById(c *gin.Context) {
	// id_user := c.Param("id_user")
	var (
		user                models.User
		detailUser          models.DetailUser
		keahlianUser        []models.KeahlianUsers
		pengalamanUser      []models.PengalamanUser
		pendidikanFormal    []models.PendidikanFormalUser
		pendidikanNonFormal []models.PendidikanNonFormalUser
	)
	// _user := db.Where("id_user = ?", id_user).Take(&user)
	// if _user.Error != nil {
	// 	_resError(c, "error", _user.Error)
	// 	return
	// }
	// detail := db.Where("id_user = ?", id_user).Take(&detailUser)
	// if detail.Error != nil {
	// 	_resError(c, "error", detail.Error)
	// 	return
	// }
	// pendidikan_formal := db.Where("id_user = ?", id_user).Take(&pendidikanFormal)
	// if pendidikan_formal.Error != nil {
	// 	_resError(c, "error", pendidikan_formal.Error)
	// 	return
	// }
	// pendidikan_non_formal := db.Where("id_user = ?", id_user).Take(&pendidikanNonFormal)
	// if pendidikan_non_formal.Error != nil {
	// 	_resError(c, "error", pendidikan_non_formal.Error)
	// 	return
	// }
	// pengalaman_user := db.Where("id_user = ?", id_user).Take(&pengalamanUser)
	// if pengalaman_user.Error != nil {
	// 	_resError(c, "error", pengalaman_user.Error)
	// 	return
	// }
	// keahlian_user := db.Where("id_user = ?", id_user).Take(&keahlianUser)
	// if keahlian_user.Error != nil {
	// 	_resError(c, "error", keahlian_user.Error)
	// 	return
	// }
	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	// obj
	db.Where("email = ?", data.Email).Take(&user)
	db.Where("id_user = ?", user.Id).Take(&detailUser)
	// arr
	db.Where("id_user = ?", user.Id).Find(&pendidikanFormal)
	db.Where("id_user = ?", user.Id).Find(&pendidikanNonFormal)
	db.Where("id_user = ?", user.Id).Find(&pengalamanUser)
	db.Where("id_user = ?", user.Id).Find(&keahlianUser)

	fullname := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	result := models.GetUserByIdResponse{
		Id:                      user.Id,
		Fullname:                fullname,
		Email:                   user.Email,
		Gender:                  detailUser.Gender,
		Usia:                    detailUser.Usia,
		NoHp:                    detailUser.NoHp,
		Alamat:                  detailUser.Alamat,
		TanggalLahir:            detailUser.TanggalLahir,
		Cv:                      detailUser.Cv,
		KeahlianUsers:           keahlianUser,
		PendidikanNonFormalUser: pendidikanNonFormal,
		PengalamanUser:          pengalamanUser,
		PendidikanFormalUser:    pendidikanFormal,
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": result})
}

func GetUserHistoryLamaranById(c *gin.Context) {
	// id_user := c.Param("id_user")
	var apply []models.ApplyLamaranUser

	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	// obj
	db.Where("id_user = ?", data.Id).Find(&apply)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": apply})
}

func GetAllLowonganCompany(c *gin.Context) {

	var (
		_data          []models.ResLowonganPerusahaan
		company        models.Perusahaan
		detail_company models.DetailPerusahaan
		lowongan       []models.LowonganPerusahaan
		detailUser     models.DetailUser
	)

	data, err := utils.DecodedTokenBearer(c, db)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}

	getDetailUser := db.Where("id_user = ?", data.Id).Take(&detailUser)
	if getDetailUser.Error != nil {
		_resError(c, "server internal error", getDetailUser.Error)
		return
	}

	allLowongan := db.Find(&lowongan)
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
