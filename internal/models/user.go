package models

type DetailUser struct {
	Id           int    `gorm:"column:id"`
	Gender       string `gorm:"type:varchar(255)" json:"gender"`
	Usia         int    `gorm:"type:int" json:"usia"`
	NoHp         string `gorm:"type:varchar(255)" json:"no_hp"`
	Alamat       string `gorm:"type:text" json:"alamat"`
	TanggalLahir string `gorm:"type:date" json:"tanggal_lahir"`
	Cv           string `gorm:"type:varchar(255)" json:"cv"`
}

type PendidikanFormalUser struct {
	Id                int    `gorm:"column:id"`
	NamaSekolah       string `gorm:"type:varchar(255)" json:"nama_sekolah"`
	TanggalMasuk      string `gorm:"type:date" json:"tanggal_masuk"`
	TanggalLulus      string `gorm:"type:date" json:"tanggal_lulus"`
	TingkatPendidikan string `gorm:"type:varchar(255)" json:"tingkat_pendidikan"`
	Jurusan           string `gorm:"type:varchar(255)" json:"jurusan"`
}

type PengalamanUser struct {
	Id             int    `gorm:"column:id"`
	NamaPerusahaan string `gorm:"type:varchar(255)" json:"nama_perusahaan"`
	TanggalMasuk   string `gorm:"type:date" json:"tanggal_masuk"`
	TanggalKeluar  string `gorm:"type:date" json:"tanggal_keluar"`
	PosisiTerakhir string `gorm:"type:varchar(255)" json:"posisi_terakhir"`
}

type PendidikanNonFormalUser struct {
	Id           int    `gorm:"column:id"`
	NamaSekolah  string `gorm:"type:varchar(255)" json:"nama_sekolah"`
	TanggalMasuk string `gorm:"type:date" json:"tanggal_masuk"`
	TanggalLulus string `gorm:"type:date" json:"tanggal_lulus"`
	Jurusan      string `gorm:"type:varchar(255)" json:"jurusan"`
}

type KeahlianUsers struct {
	Id           int    `gorm:"column:id"`
	NamaKeahlian string `gorm:"type:varchar(255)" json:"nama_keahlian"`
	Level        string `gorm:"type:varchar(255)" json:"level"`
}

type GetUserByIdResponse struct {
	Id                      int
	Fullname                string
	Email                   string
	Gender                  string                    `json:"gender"`
	Usia                    int                       `json:"usia"`
	NoHp                    string                    `json:"no_hp"`
	Alamat                  string                    `json:"alamat"`
	TanggalLahir            string                    `json:"tanggal_lahir"`
	Cv                      string                    `json:"cv"`
	KeahlianUsers           []KeahlianUsers           `json:"keahlian"`
	PendidikanFormalUser    []PendidikanFormalUser    `json:"pendidikan_formal"`
	PengalamanUser          []PengalamanUser          `json:"pengalaman"`
	PendidikanNonFormalUser []PendidikanNonFormalUser `json:"pendidikan_non_formal"`
}

type ApplyLamaranUser struct {
	Id        int    `gorm:"column:id" json:"id"`
	IdUser    int    `gorm:"column:id_user" json:"id_user"`
	Pesan     string `gorm:"type:text" json:"pesan"`
	IdCompany int    `gorm:"column:id_company" json:"id_company"`
	CreatedAt string `gorm:"column:created_at" json:"tanggal_lamar"`
}

type ResApplyLamaranUser struct {
	Id        int    `gorm:"column:id" json:"id"`
	IdUser    int    `gorm:"column:id_user" json:"id_user"`
	Pesan     string `gorm:"type:text" json:"pesan"`
	IdCompany int    `gorm:"column:id_company" json:"id_company"`
	CreatedAt string `gorm:"column:created_at" json:"tanggal_lamar"`
}
