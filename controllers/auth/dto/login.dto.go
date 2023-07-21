package dto

type LoginDto struct {
	Username string `gorm:"type:varchar(255)" json:"username" validate:"required"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(255)" json:"password" validate:"required"`
}