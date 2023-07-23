package schema

type Keahlian struct {
	Id           int    `gorm:"column:id_keahlian" json:"id_keahlian"`
	NamaKeahlian string `gorm:"varchar(255)" json:"nama_keahlian" validate:"required"`
	Image        string `gorm:"varchar(255)" json:"image" validate:"required"`
}

type BodyKeahlian struct {
	NamaKeahlian string `gorm:"varchar(255)" json:"nama_keahlian" validate:"required"`
	Image        string `gorm:"varchar(255)" json:"image" validate:"required"`
}
