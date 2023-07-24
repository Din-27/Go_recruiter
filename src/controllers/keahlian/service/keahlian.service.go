package keahlian

import (
	"context"
	"net/http"
	"time"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/src/config"
	"github.com/Din-27/Go_job/src/controllers/keahlian/schema"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

var (
	db        = config.DBinit()
	_resError = helpers.ResponseError
	_isErr    = helpers.ErrorReturn
)

func ListKeahlian(c *gin.Context) {
	var keahlian []schema.Keahlian
	result := db.Find(&keahlian)
	if result.Error != nil {
		_resError(c, "error", _isErr("Email tidak ditemukan !"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": keahlian})
}

func AddKeahlian(c *gin.Context) {
	var keahlian schema.Keahlian

	nama := c.PostForm("nama_keahlian")
	fileHeader, _ := c.FormFile("image")
	file, _ := fileHeader.Open()

	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL("cloudinary://394378188537969:MadSDoFe38KVudmltbRKecmjy1U@dce5mf135")
	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:   "image_keahlian",
		PublicID: fileHeader.Filename + time.Now().String(),
	})

	keahlian.NamaKeahlian = nama
	keahlian.Image = resp.SecureURL

	result := db.Create(&keahlian)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": nama, "image": fileHeader})
}
