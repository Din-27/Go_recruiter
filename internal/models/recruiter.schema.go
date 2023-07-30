package models

import "time"

type DetailPerusahaan struct {
	Id             int    `gorm:"column:id_perusahaan" json:"id_perusahaan"`
	Nama           string `gorm:"type:varchar(255)" json:"nama"`
	Alamat         string `gorm:"type:text" json:"alamat"`
	Deskripsi      string `gorm:"type:text" json:"deskripsi"`
	Bidang         string `gorm:"type:text" json:"bidang"`
	Pencapaian     string `gorm:"type:text" json:"pencapaian"`
	JumlahKaryawan int    `gorm:"type:int" json:"jumlah_karyawan"`
	Website        string `gorm:"type:varchar(255)" json:"website"`
	Logo           string `gorm:"type:varchar(255)" json:"logo"`
	Background     string `gorm:"type:varchar(255)" json:"background"`
}

type LowonganPerusahaan struct {
	Id             int       `gorm:"column:id_perusahaan" json:"id_perusahaan"`
	Title          string    `gorm:"type:varchar(255)" json:"title"`
	Deskripsi      string    `gorm:"type:text" json:"deskripsi"`
	MinGaji        int       `gorm:"type:int" json:"min_gaji"`
	MaxGaji        int       `gorm:"type:int" json:"max_gaji"`
	Poster         string    `gorm:"type:varchar(255)" json:"poster"`
	DurasiLowongan time.Time `gorm:"type:date" json:"durasi_lowongan"`
}

type RequirementLowonganPerusahaan struct {
	Id        int    `gorm:"column:id_perusahaan" json:"id_perusahaan"`
	Nama      string `gorm:"type:varchar(255)" json:"nama"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}

type BenefitLowonganPerusahaan struct {
	Id   int    `gorm:"column:id_perusahaan" json:"id_perusahaan"`
	Nama string `gorm:"type:varchar(255)" json:"nama"`
}

type AddLowongan struct {
	Id             int       `gorm:"column:id_perusahaan" json:"id_perusahaan"`
	Title          string    `gorm:"type:varchar(255)" json:"title"`
	Deskripsi      string    `gorm:"type:text" json:"deskripsi"`
	MinGaji        int       `gorm:"type:int" json:"min_gaji"`
	MaxGaji        int       `gorm:"type:int" json:"max_gaji"`
	Poster         string    `gorm:"type:varchar(255)" json:"poster"`
	DurasiLowongan time.Time `gorm:"type:date" json:"durasi_lowongan"`
	Requirement    []RequirementLowonganPerusahaan
	Benefit        []BenefitLowonganPerusahaan
}
