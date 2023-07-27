package schema

type User struct {
	Id         int    `gorm:"column:id_user" json:"id_user"`
	FirstName  string `gorm:"type:varchar(255)" json:"first_name" validate:"required"`
	LastName   string `gorm:"type:varchar(255)" json:"last_name" validate:"required"`
	Username   string `gorm:"type:varchar(255)" json:"username" validate:"required"`
	Email      string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Password   string `gorm:"type:varchar(255)" json:"password" validate:"required"`
	Specialist string `gorm:"type:varchar(255)" json:"specialist" validate:"required"`
}

type Login struct {
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(255)" json:"password" validate:"required"`
}

type ResponseLogin struct {
	Id           int    `gorm:"column:id_user, primaryKey" json:"id_user" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Specialist   string `json:"specialist" validate:"required"`
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ResponseRefresh struct {
	Id          int    `gorm:"column:id_user, primaryKey" json:"id_user" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Specialist  string `json:"specialist" validate:"required"`
	AccessToken string `json:"access_token" validate:"required"`
}