package handler

import (
	"fmt"
	"net/http"

	"github.com/Din-27/Go_job/internal/models"
	"github.com/Din-27/Go_job/internal/utils"
	"github.com/gin-gonic/gin"
)

func AddUserDetail(c *gin.Context) {
	var (
		user        models.User
		user_detail models.DetailUser
	)

	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	getIdUser := db.Where("email = ?", data.Email).Take(&user)
	if getIdUser.Error != nil {
		_resError(c, "server internal error", getIdUser.Error)
		return
	}
	user_detail.Id = user.Id
	result := db.Create(&user_detail)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan detail diri anda"})
}

func AddUserPendidikanFormal(c *gin.Context) {
	var (
		user        models.User
		formal_user models.PendidikanFormalUser
	)
	if err := c.ShouldBindJSON(&formal_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	getIdUser := db.Where("email = ?", data.Email).Take(&user)
	if getIdUser.Error != nil {
		_resError(c, "server internal error", getIdUser.Error)
		return
	}
	formal_user.Id = user.Id
	result := db.Create(&formal_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan pendidikan formal anda"})
}

func AddUserPendidikanNonFormal(c *gin.Context) {
	var (
		user            models.User
		non_formal_user models.PendidikanNonFormalUser
	)
	if err := c.ShouldBindJSON(&non_formal_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	getIdUser := db.Where("email = ?", data.Email).Take(&user)
	if getIdUser.Error != nil {
		_resError(c, "server internal error", getIdUser.Error)
		return
	}
	non_formal_user.Id = user.Id
	result := db.Create(&non_formal_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan pendidikan non formal anda"})
}

func AddUserPengalaman(c *gin.Context) {
	var (
		user            models.User
		pengalaman_user models.PengalamanUser
	)
	if err := c.ShouldBindJSON(&pengalaman_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	getIdUser := db.Where("email = ?", data.Email).Take(&user)
	if getIdUser.Error != nil {
		_resError(c, "server internal error", getIdUser.Error)
		return
	}
	pengalaman_user.Id = user.Id
	result := db.Create(&pengalaman_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan pengalaman anda"})
}

func AddUserKeahlian(c *gin.Context) {
	var (
		user          models.User
		keahlian_user models.KeahlianUsers
	)
	if err := c.ShouldBindJSON(&keahlian_user); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	getIdUser := db.Where("email = ?", data.Email).Take(&user)
	if getIdUser.Error != nil {
		_resError(c, "server internal error", getIdUser.Error)
		return
	}
	keahlian_user.Id = user.Id
	result := db.Create(&keahlian_user)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": "sukses menambahkan keahlian anda"})
}

func ApplyLamaranUser(c *gin.Context) {
	var (
		user  models.User
		apply models.ApplyLamaranUser
	)
	if err := c.ShouldBindJSON(&apply); err != nil {
		_resError(c, "error", err)
		return
	}
	data, err := utils.DecodedTokenBearer(c)
	if err != nil {
		_resError(c, "server internal error", err)
		return
	}
	if data.Role != "user" {
		_resError(c, "server internal error", _isErr("url ini untuk user !"))
		return
	}
	getIdUser := db.Where("email = ?", data.Email).Take(&user)
	if getIdUser.Error != nil {
		_resError(c, "error", getIdUser.Error)
		return
	}
	apply.IdUser = user.Id
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
	data, err := utils.DecodedTokenBearer(c)
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
	var (
		user  models.User
		apply []models.ApplyLamaranUser
	)

	data, err := utils.DecodedTokenBearer(c)
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

	db.Where("id_user = ?", user.Id).Find(&apply)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": apply})
}
