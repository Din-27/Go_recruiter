package models

type DetailPerusahaan struct {
	Id             int    `gorm:"column:id_company" json:"id_company"`
	Didirikan      string `gorm:"type:date" json:"didirikan"`
	Deskripsi      string `gorm:"type:text" json:"deskripsi"`
	Industri       string `gorm:"type:varchar(255)" json:"industri"`
	IdTeknologi    int    `gorm:"type:int" json:"id_teknologi"`
	IdSosmed       int    `gorm:"type:int" json:"id_sosmed"`
	JumlahKaryawan int    `gorm:"type:int" json:"jumlah_karyawan"`
	Lokasi         string `gorm:"type:varchar(255)" json:"lokasi"`
	Website        string `gorm:"type:varchar(255)" json:"website"`
	Logo           string `gorm:"type:varchar(255)" json:"logo"`
	Background     string `gorm:"type:varchar(255)" json:"background"`
}

type TeknologiPerusahaan struct {
	Id   int    `gorm:"column:id_company" json:"id_company"`
	Nama string `gorm:"type:varchar(255)" json:"nama"`
}

type SosmedPerusahaan struct {
	Id   int    `gorm:"column:id_company" json:"id_company"`
	Link string `gorm:"type:varchar(255)" json:"link"`
}

type LowonganPerusahaan struct {
	Id             int    `gorm:"column:id_company" json:"id_company"`
	IdLowongan     int    `gorm:"primaryKey" json:"id_lowongan"`
	Title          string `gorm:"type:varchar(255)" json:"title"`
	Category       string `gorm:"type:varchar(255)" json:"category"`
	Deskripsi      string `gorm:"type:text" json:"deskripsi"`
	MinGaji        int    `gorm:"type:int" json:"min_gaji"`
	MaxGaji        int    `gorm:"type:int" json:"max_gaji"`
	Poster         string `gorm:"type:varchar(255)" json:"poster"`
	TipePekerjaan  string `gorm:"type:varchar(255)" json:"tipe_pekerjaan"`
	LevelPekerjaan string `gorm:"type:varchar(255)" json:"level_pekerjaan"`
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
	Deskripsi  string `gorm:"type:text" json:"deskripsi"`
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
	Category       string                          `gorm:"type:varchar(255)" json:"category"`
	Deskripsi      string                          `gorm:"type:text" json:"deskripsi"`
	MinGaji        int                             `gorm:"type:int" json:"min_gaji"`
	MaxGaji        int                             `gorm:"type:int" json:"max_gaji"`
	Poster         string                          `gorm:"type:varchar(255)" json:"poster"`
	DurasiLowongan string                          `gorm:"type:date" json:"durasi_lowongan"`
	TipePekerjaan  string                          `gorm:"type:varchar(255)" json:"tipe_pekerjaan"`
	LevelPekerjaan string                          `gorm:"type:varchar(255)" json:"level_pekerjaan"`
	Requirement    []RequirementLowonganPerusahaan `json:"requirement"`
	Benefit        []BenefitLowonganPerusahaan     `json:"benefit"`
}

type GetCompanyByIdResponse struct {
	Id             int              `json:"id_company"`
	Nama           string           `json:"nama"`
	Didirikan      string           `json:"didirikan"`
	Deskripsi      string           `json:"deskripsi"`
	Industri       string           `gorm:"type:varchar(255)" json:"industri"`
	IdTeknologi    int              `gorm:"type:int" json:"id_teknologi"`
	JumlahKaryawan int              `json:"jumlah_karyawan"`
	Website        string           `json:"website"`
	Logo           string           `json:"logo"`
	Background     string           `json:"background"`
	DetailLowongan []DetailLowongan `json:"detail_lowongan"`
}

type ResLowonganPerusahaan struct {
	Id             int    `gorm:"column:id_company" json:"id_company"`
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

type ResFilter struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
}
