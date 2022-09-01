package riderRepository

import (
	"github.com/google/uuid"
	"sls/internal/entity/riderEntity"
)

type RiderRepo interface {
	Persist(req *riderEntity.CreateRiderReq) (*riderEntity.CreateRiderReq, error)
	GetByEmail(email string) (*riderEntity.CreateRiderRes, error)
	GetByPhone(phone string) (*riderEntity.CreateRiderRes, error)
	GetById(id uuid.UUID) (*riderEntity.CreateRiderRes, error)
	UpdateProfile(req *riderEntity.UpdateRiderReq) (*riderEntity.UpdateRiderReq, error)
	ChangePassword(id uuid.UUID, password string) error
	DeleteAccount(req *riderEntity.DeleteUser) error
}
