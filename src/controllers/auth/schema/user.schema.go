package schema

type User struct {
	Id         int    `gorm:"type:int" json:"id" validate:"required"`
	FirstName  string `gorm:"type:varchar(255)" json:"first_name" validate:"required"`
	LastName   string `gorm:"type:varchar(255)" json:"last_name" validate:"required"`
	Username   string `gorm:"type:varchar(255)" json:"username" validate:"required"`
	Email      string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Password   string `gorm:"type:varchar(255)" json:"password" validate:"required"`
	Specialist string `gorm:"type:varchar(255)" json:"specialist" validate:"required"`
}

type Register struct {
	FirstName  string `gorm:"type:varchar(255)" json:"first_name" validate:"required"`
	LastName   string `gorm:"type:varchar(255)" json:"last_name" validate:"required"`
	Username   string `gorm:"type:varchar(255)" json:"username" validate:"required"`
	Email      string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Password   string `gorm:"type:varchar(255)" json:"password" validate:"required"`
	Specialist string `gorm:"type:varchar(255)" json:"specialist" validate:"required"`
}

type Login struct {
	Username string `gorm:"type:varchar(255)" json:"username" validate:"required"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(255)" json:"password" validate:"required"`
}
