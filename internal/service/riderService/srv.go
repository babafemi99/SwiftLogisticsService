package riderService

import (
	"errors"
	"github.com/google/uuid"
	"sls/internal/Repository/riderRepository"
	"sls/internal/entity/riderEntity"
	"sls/internal/service/cryptoService"
	"sls/internal/service/timeService"
	"sls/internal/service/validationService"
	"time"
)

type RiderSrv interface {
	Create(req *riderEntity.CreateRiderReq) (*riderEntity.CreateRiderRes, error)
	Login(req *riderEntity.LoginReq) (*riderEntity.CreateRiderRes, error)
	ResetPassword(password *riderEntity.ResetPassword) error
	UpdateProfile(id string, req *riderEntity.UpdateRiderReq) (*riderEntity.UpdateRiderRes, error)
	ChangePassword(req *riderEntity.ChangePassword) error
	DeleteAccount(req *riderEntity.DeleteUser) error
}

type riderSrv struct {
	timeSrv   timeService.TimeSrv
	crypto    cryptoService.CryptoSrv
	vldSrv    validationService.ValidationService
	riderRepo riderRepository.RiderRepo
}

func NewRiderSrv(timeSrv timeService.TimeSrv, crypto cryptoService.CryptoSrv, vldSrv validationService.ValidationService, riderRepo riderRepository.RiderRepo) RiderSrv {
	return &riderSrv{timeSrv: timeSrv, crypto: crypto, vldSrv: vldSrv, riderRepo: riderRepo}
}

func (r *riderSrv) Login(req *riderEntity.LoginReq) (*riderEntity.CreateRiderRes, error) {
	err := r.vldSrv.Validate(req)
	if err != nil {
		return nil, err
	}

	rider, err := r.riderRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = r.crypto.ComparePassword(rider.Password, req.Password)
	if err != nil {
		return nil, err
	}

	return rider, nil

}

func (r *riderSrv) ResetPassword(req *riderEntity.ResetPassword) error {
	err := r.vldSrv.Validate(req)
	if err != nil {
		return err
	}

	password, err := r.crypto.HashPassword(req.Password)
	if err != nil {
		return err
	}

	err = r.riderRepo.ChangePassword(req.Id, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *riderSrv) UpdateProfile(id string, req *riderEntity.UpdateRiderReq) (*riderEntity.UpdateRiderRes, error) {
	err := r.vldSrv.Validate(req)
	if err != nil {
		return nil, err
	}

	profile, err := r.riderRepo.UpdateProfile(id, req)
	if err != nil {
		return nil, err
	}

	data := &riderEntity.UpdateRiderRes{
		RiderId:        profile.RiderId,
		FirstName:      profile.FirstName,
		LastName:       profile.LastName,
		Phone:          profile.Phone,
		Email:          profile.Email,
		PhoneNumber:    profile.PhoneNumber,
		ProfilePicture: profile.ProfilePicture,
	}

	return data, nil
}

func (r *riderSrv) ChangePassword(req *riderEntity.ChangePassword) error {
	err := r.vldSrv.Validate(req)
	if err != nil {
		return err
	}

	rider, err := r.riderRepo.GetById(req.RiderId)
	if err != nil {
		return err
	}

	err = r.crypto.ComparePassword(rider.Password, req.OldPassword)
	if err != nil {
		return err
	}

	password, err := r.crypto.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	err = r.riderRepo.ChangePassword(req.RiderId, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *riderSrv) DeleteAccount(req *riderEntity.DeleteUser) error {
	err := r.vldSrv.Validate(req)
	if err != nil {
		return err
	}

	err = r.riderRepo.DeleteAccount(req)
	if err != nil {
		return err
	}

	return nil
}

func (r *riderSrv) Create(req *riderEntity.CreateRiderReq) (*riderEntity.CreateRiderRes, error) {
	err := r.vldSrv.Validate(req)
	if err != nil {
		return nil, err
	}

	_, err = r.riderRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	password, err := r.crypto.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.RiderId = uuid.New().String()
	req.Password = password
	req.DateCreated = r.timeSrv.Now().Format(time.RFC3339)
	req.AccountStatus = "PENDING"

	for _, grt := range req.Guarantor {
		grt.GuarantorId = uuid.New().String()
		grt.RiderId = req.RiderId
	}

	rider, err := r.riderRepo.Persist(req)
	if err != nil {
		return nil, err
	}

	finalData := &riderEntity.CreateRiderRes{
		RiderId:        rider.RiderId,
		FirstName:      rider.FirstName,
		LastName:       rider.LastName,
		Phone:          rider.PhoneNumber,
		Email:          rider.Email,
		Password:       rider.Password,
		ProfilePicture: rider.ProfilePicture,
		AccountStatus:  rider.AccountStatus,
	}

	return finalData, nil
}
