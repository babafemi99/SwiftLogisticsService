package handlerEntity

import "sls/internal/entity/userEntity"

type ClientAccess struct {
	Token          string `json:"token"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Password       string `json:"-"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
}

type SuccessRes struct {
	Message string `json:"message"`
}

func NewClientAccess(token string, access *userEntity.UserAccess) *ClientAccess {
	return &ClientAccess{
		Token:          token,
		Email:          access.Email,
		Phone:          access.Phone,
		Password:       access.Password,
		FirstName:      access.FirstName,
		LastName:       access.LastName,
		ProfilePicture: access.ProfilePicture,
	}
}

func NewSuccessMessage(msg string) *SuccessRes {
	return &SuccessRes{
		Message: msg,
	}
}
