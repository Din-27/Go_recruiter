package models

type DetailPerusahaan struct {
	Id             int    `gorm:"column:id_company" json:"id_company"`
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
	Id             int    `gorm:"column:id_company" json:"id_company"`
	IdLowongan     int    `gorm:"primaryKey" json:"id_lowongan"`
	Title          string `gorm:"type:varchar(255)" json:"title"`
	Deskripsi      string `gorm:"type:text" json:"deskripsi"`
	MinGaji        int    `gorm:"type:int" json:"min_gaji"`
	MaxGaji        int    `gorm:"type:int" json:"max_gaji"`
	Poster         string `gorm:"type:varchar(255)" json:"poster"`
	DurasiLowongan string `gorm:"type:date" json:"durasi_lowongan"`
}

type RequirementLowonganPerusahaan struct {
	Id         int    `json:"id"`
	IdLowongan int    `json:"id_lowongan"`
	Nama       string `gorm:"type:varchar(255)" json:"nama"`
	Deskripsi  string `gorm:"type:text" json:"deskripsi"`
}

type BenefitLowonganPerusahaan struct {
	Id         int    `json:"id"`
	IdLowongan int    `json:"id_lowongan"`
	Nama       string `gorm:"type:varchar(255)" json:"nama"`
}

type DetailLowongan struct {
	IdLowongan                    int                             `json:"id_lowongan"`
	Title                         string                          `json:"title"`
	Deskripsi                     string                          `json:"deskripsi"`
	MinGaji                       int                             `json:"min_gaji"`
	MaxGaji                       int                             `json:"max_gaji"`
	Poster                        string                          `json:"poster"`
	DurasiLowongan                string                          `json:"durasi_lowongan"`
	RequirementLowonganPerusahaan []RequirementLowonganPerusahaan `json:"requirement"`
	BenefitLowonganPerusahaan     []BenefitLowonganPerusahaan     `json:"benefit"`
}

type AddLowongan struct {
	Id             int                             `gorm:"column:id_lowongan" json:"id_lowongan"`
	Title          string                          `gorm:"type:varchar(255)" json:"title"`
	Deskripsi      string                          `gorm:"type:text" json:"deskripsi"`
	MinGaji        int                             `gorm:"type:int" json:"min_gaji"`
	MaxGaji        int                             `gorm:"type:int" json:"max_gaji"`
	Poster         string                          `gorm:"type:varchar(255)" json:"poster"`
	DurasiLowongan string                          `gorm:"type:date" json:"durasi_lowongan"`
	Requirement    []RequirementLowonganPerusahaan `json:"requirement"`
	Benefit        []BenefitLowonganPerusahaan     `json:"benefit"`
}

type GetCompanyByIdResponse struct {
	Id             int              `json:"id_company"`
	Nama           string           `json:"nama"`
	Alamat         string           `json:"alamat"`
	Deskripsi      string           `json:"deskripsi"`
	Bidang         string           `json:"bidang"`
	Pencapaian     string           `json:"pencapaian"`
	JumlahKaryawan int              `json:"jumlah_karyawan"`
	Website        string           `json:"website"`
	Logo           string           `json:"logo"`
	Background     string           `json:"background"`
	DetailLowongan []DetailLowongan `json:"detail_lowongan"`
}

type ResLowonganPerusahaan struct {
	Id             int    `json:"id_company"`
	Logo           string `json:"logo_perusahaan"`
	Nama           string `json:"nama_perusahaan"`
	IdLowongan     int    `json:"id_lowongan"`
	Title          string `json:"title"`
	Deskripsi      string `json:"deskripsi"`
	MinGaji        int    `json:"min_gaji"`
	MaxGaji        int    `json:"max_gaji"`
	Poster         string `json:"poster"`
	DurasiLowongan string `json:"durasi_lowongan"`
}
