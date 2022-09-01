package userRepository

import (
	"github.com/google/uuid"
	"sls/internal/entity/userEntity"
)

type UserRepo interface {
	Persist(user *userEntity.CreateUser) (*userEntity.CreateUser, error)
	GetByEmail(email string) (*userEntity.UserAccess, error)
	GetByPhone(phone string) (*userEntity.UserAccess, error)
	GetById(id uuid.UUID) (*userEntity.UserAccess, error)
	ChangePassword(id uuid.UUID, password string) error
	UpdateProfile(req *userEntity.UpdateUserReq) (*userEntity.UpdateUserReq, error)
	DeactivateAccount(id uuid.UUID) error
}
