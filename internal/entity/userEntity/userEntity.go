package userEntity

import "github.com/google/uuid"

type CreateUser struct {
	UserId          uuid.UUID `json:"user_id"`
	Email           string    `json:"email" validate:"required,email"`
	Phone           string    `json:"phone" validate:"required,e164"`
	FirstName       string    `json:"first_name" validate:"required"`
	LastName        string    `json:"last_name" validate:"required"`
	Password        string    `json:"password" validate:"required,alphanum"`
	ProfilePicture  string    `json:"profile_picture"`
	DefaultLocation string    `json:"default_location" validate:"required"`
	UserStatus      string    `json:"user_status"`
	UpdateAt        string    `json:"update_at"`
	CreatedAt       string    `json:"created_at"`
}

type UserAccess struct {
	UserId         uuid.UUID `json:"user_id"`
	Email          string    `json:"email" validate:"required,email"`
	Phone          string    `json:"phone" validate:"required,e164"`
	Password       string    `json:"-"`
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	ProfilePicture string    `json:"profile_picture"`
	Status         string    `json:"status"`
}

type UpdateUserReq struct {
	UserId         uuid.UUID `json:"user_id" validate:"required"`
	Email          string    `json:"email" validate:"required,email"`
	Phone          string    `json:"phone" validate:"required,e164"`
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	ProfilePicture string    `json:"profile_picture"`
}

type UpdateUserRes struct {
	Email          string `json:"email" validate:"required,email"`
	Phone          string `json:"phone" validate:"required,e164"`
	Password       string `json:"-"`
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	ProfilePicture string `json:"profile_picture"`
}

type LoginReq struct {
	Credential string `json:"credential" validate:"required,email|e164"`
	Password   string `json:"password" validate:"required"`
}

type ResetPassword struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type DeleteAccountReq struct {
	UserId uuid.UUID `json:"user_id" validate:"required"`
}

type ChangePassword struct {
	UserId      uuid.UUID `json:"-"`
	OldPassword string    `json:"old_password" validate:"required"`
	NewPassword string    `json:"new_password" validate:"required"`
}
