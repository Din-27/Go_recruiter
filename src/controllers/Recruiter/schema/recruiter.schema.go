package schema

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

type Perusahaan struct {
	Id       int    `gorm:"column:id_perusahaan" json:"id_perusahaan"`
	Nama     string `gorm:"type:varchar(255)" json:"nama" validate:"required"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(255)" json:"password" validate:"required"`
	Role     string `gorm:"type:varchar(255)" json:"role"`
}