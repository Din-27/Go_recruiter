package models

import "time"

type DetailUser struct {
	Id           int       `gorm:"column:id_user"`
	Gender       string    `gorm:"type:varchar(255)" json:"gender"`
	Usia         int       `gorm:"type:int" json:"usia"`
	NoHp         string    `gorm:"type:varchar(255)" json:"no_hp"`
	Alamat       string    `gorm:"type:text" json:"alamat"`
	TanggalLahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	Cv           string    `gorm:"type:varchar(255)" json:"cv"`
}

type PendidikanFormalUser struct {
	Id                int       `gorm:"column:id_user"`
	NamaSekolah       string    `gorm:"type:varchar(255)" json:"nama_sekolah"`
	TanggalMasuk      time.Time `gorm:"type:date" json:"tanggal_masuk"`
	TanggalLulus      time.Time `gorm:"type:date" json:"tanggal_lulus"`
	TingkatPendidikan string    `gorm:"type:varchar(255)" json:"tingkat_pendidikan"`
	Jurusan           string    `gorm:"type:varchar(255)" json:"jurusan"`
}

type PengalamanUser struct {
	Id             int       `gorm:"column:id_user"`
	NamaPerusahaan string    `gorm:"type:varchar(255)" json:"nama_perusahaan"`
	TanggalMasuk   time.Time `gorm:"type:date" json:"tanggal_masuk"`
	TanggalKeluar  time.Time `gorm:"type:date" json:"tanggal_keluar"`
	PosisiTerakhir string    `gorm:"type:varchar(255)" json:"posisi_terakhir"`
}

type PendidikanNonFormalUser struct {
	Id            int       `gorm:"column:id_user"`
	Penyelenggara string    `gorm:"type:varchar(255)" json:"penyelenggara"`
	TanggalMasuk  time.Time `gorm:"type:date" json:"tanggal_masuk"`
	TanggalLulus  time.Time `gorm:"type:date" json:"tanggal_lulus"`
	Jurusan       string    `gorm:"type:varchar(255)" json:"jurusan"`
}

type KeahlianUsers struct {
	Id           int    `gorm:"column:id_user"`
	NamaKeahlian string `gorm:"type:varchar(255)" json:"nama_keahlian"`
	Level        string `gorm:"type:varchar(255)" json:"level"`
}

type GetUserByIdResponse struct {
	Fullname string
	Email    string
	DetailUser
	KeahlianUsers           []KeahlianUsers
	PendidikanFormalUser    []PendidikanFormalUser
	PengalamanUser          []PengalamanUser
	PendidikanNonFormalUser []PendidikanNonFormalUser
}

type LamaranUser struct {
	Fullname string
	DetailUser
	KeahlianUsers           []KeahlianUsers
	PendidikanFormalUser    []PendidikanFormalUser
	PengalamanUser          []PengalamanUser
	PendidikanNonFormalUser []PendidikanNonFormalUser
}
