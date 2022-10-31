package riderRepository

import (
	"sls/internal/entity/riderEntity"
)

type RiderRepo interface {
	Persist(req *riderEntity.CreateRiderReq) (*riderEntity.CreateRiderReq, error)
	GetByEmail(email string) (*riderEntity.CreateRiderRes, error)
	GetByPhone(phone string) (*riderEntity.CreateRiderRes, error)
	GetById(id string) (*riderEntity.CreateRiderRes, error)
	UpdateProfile(id string, req *riderEntity.UpdateRiderReq) (*riderEntity.UpdateRiderReq, error)
	ChangePassword(id string, password string) error
	DeleteAccount(req *riderEntity.DeleteUser) error
}
