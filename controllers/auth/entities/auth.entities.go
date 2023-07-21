package entities

import "github.com/Din-27/Go_job/controllers/auth/dto"

type UserRepository interface {
	Register(user dto.UserDto) (dto.UserDto, error)
}
