package userService

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"sls/internal/Repository/userRepository"
	"sls/internal/entity/errorEntity"
	"sls/internal/entity/userEntity"
	"sls/internal/service/cryptoService"
	"sls/internal/service/timeService"
	"sls/internal/service/validationService"
	"strings"
	"time"
)

type UserSrv interface {
	SignUp(user *userEntity.CreateUser) (*userEntity.UserAccess, *errorEntity.ErrorRes)
	Login(req *userEntity.LoginReq) (*userEntity.UserAccess, *errorEntity.ErrorRes)
	UpdateProfile(req *userEntity.UpdateUserReq) (*userEntity.UpdateUserRes, *errorEntity.ErrorRes)
	ChangePassword(password *userEntity.ChangePassword) *errorEntity.ErrorRes
	ResetPassword(req *userEntity.ResetPassword) *errorEntity.ErrorRes
	DeleteAccount(req *userEntity.DeleteAccountReq) *errorEntity.ErrorRes
}

type userSrv struct {
	log       *logrus.Logger
	cryptoSrv cryptoService.CryptoSrv
	timeSrv   timeService.TimeSrv
	vldSrv    validationService.ValidationService
	repoSrv   userRepository.UserRepo
}

func (u *userSrv) SignUp(user *userEntity.CreateUser) (*userEntity.UserAccess, *errorEntity.ErrorRes) {
	// validate user input
	err := u.vldSrv.Validate(user)
	if err != nil {
		u.log.Errorf("Error when validating request body: %v", err)
		return nil, errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !")
	}

	// check if email exists
	_, err = u.repoSrv.GetByEmail(user.Email)
	if err == nil {
		u.log.Errorf("Error because email already exists: %v", err)
		return nil, errorEntity.NewErrorRes(http.StatusConflict, "User with that email exists already")
	}
	// hash the user password and update password field
	password, cryptoErr := u.cryptoSrv.HashPassword(user.Password)
	if cryptoErr != nil {
		u.log.Errorf("Error hashing password: %v", cryptoErr)
		return nil, errorEntity.NewErrorRes(http.StatusInternalServerError, "Error Hashing Password; something went wrong")
	}
	user.UserId = uuid.New()
	user.Password = password
	user.CreatedAt = u.timeSrv.Now().Local().Format(time.RFC3339)
	user.UserStatus = "active"

	// save user to db finally
	userData, err := u.repoSrv.Persist(user)
	if err != nil {
		fmt.Println(2)
		u.log.Errorf("Error saving user to db: %v", err)
		return nil, errorEntity.NewErrorRes(http.StatusInternalServerError, "Error Saving User, try again later")
	}

	finalData := &userEntity.UserAccess{
		UserId:         user.UserId,
		Email:          userData.Email,
		Phone:          userData.Phone,
		FirstName:      userData.FirstName,
		LastName:       userData.LastName,
		ProfilePicture: userData.ProfilePicture,
		Status:         user.UserStatus,
	}
	return finalData, nil
}

func (u *userSrv) Login(req *userEntity.LoginReq) (*userEntity.UserAccess, *errorEntity.ErrorRes) {

	//validate request body
	err := u.vldSrv.Validate(req)
	if err != nil {
		u.log.Errorf("Error when validating request body: %v", err)
		return nil, errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !")
	}

	//check if  credentials is number or email
	contains := strings.Contains(req.Credential, "@")
	switch contains {

	case false:
		// if it doesn't contain '@' it's a phone number
		userDetails, err := u.repoSrv.GetByPhone(req.Credential)
		if err != nil {
			u.log.Errorf("Error finding by number: %v", err)
			return nil, errorEntity.NewErrorRes(http.StatusNotFound, "Invalid Credentials")
		}

		//compare passwords
		err = u.cryptoSrv.ComparePassword(userDetails.Password, req.Password)
		if err != nil {
			u.log.Errorf("Error While comparing passwords: %v", err)
			return nil, errorEntity.NewErrorRes(http.StatusUnauthorized, "Invalid Credentials")
		}

		return userDetails, nil

	case true:
		// if it contains '@' it's a phone number
		userDetails, err := u.repoSrv.GetByEmail(req.Credential)
		if err != nil {
			u.log.Errorf("Error finding by email: %v", err)
			return nil, errorEntity.NewErrorRes(http.StatusNotFound, "Invalid Credentials: Check Email ")
		}

		//compare passwords
		err = u.cryptoSrv.ComparePassword(userDetails.Password, req.Password)
		if err != nil {
			u.log.Errorf("Error While comparing passwords: %v", err)
			return nil, errorEntity.NewErrorRes(http.StatusNotFound, "Invalid Credentials")
		}

		return userDetails, nil
	}
	return nil, nil
}

func (u *userSrv) ChangePassword(req *userEntity.ChangePassword) *errorEntity.ErrorRes {

	//validate request body
	err := u.vldSrv.Validate(req)
	if err != nil {
		u.log.Errorf("Error when validating request body: %v", err)
		return errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !")
	}

	// get Details by Email
	user, getErr := u.repoSrv.GetById(req.UserId)
	if getErr != nil {
		u.log.Errorf("Error when getting by email: %v", err)
		return errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !")
	}

	// compare passwords
	err = u.cryptoSrv.ComparePassword(user.Password, req.OldPassword)
	if err != nil {
		u.log.Errorf("Error while Comparing password: %v", err)
		return errorEntity.NewErrorRes(http.StatusUnauthorized, "Old password is not correct")
	}

	// hash new password
	tokemStr, err := u.cryptoSrv.HashPassword(req.NewPassword)
	if err != nil {
		u.log.Errorf("Error while Hashing NEw password: %v", err)
		return errorEntity.NewErrorRes(http.StatusInternalServerError, "Hashing Password Failed")
	}

	// change password
	err = u.repoSrv.ChangePassword(req.UserId, tokemStr)
	if err != nil {
		u.log.Errorf("Error while Comparing password: %v", err)
		return errorEntity.NewErrorRes(http.StatusInternalServerError, "Error Changing Password")
	}

	return nil
}

func (u *userSrv) ResetPassword(req *userEntity.ResetPassword) *errorEntity.ErrorRes {

	//validate request body
	err := u.vldSrv.Validate(req)
	if err != nil {
		u.log.Errorf("Error when validating request body: %v", err)
		return errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !")
	}

	userDetails, err := u.repoSrv.GetByEmail(req.Email)
	if err != nil {
		u.log.Errorf("Error finding by email: %v", err)
		return errorEntity.NewErrorRes(http.StatusNotFound, "Invalid Credentials: Check Email ")
	}

	tokenStr, err := u.cryptoSrv.HashPassword(req.Password)
	if err != nil {
		u.log.Errorf("Error while Comparing password: %v", err)
		return errorEntity.NewErrorRes(http.StatusInternalServerError, "Hashing Password Failed")
	}

	err = u.repoSrv.ChangePassword(userDetails.UserId, tokenStr)
	if err != nil {
		u.log.Errorf("Error when changing password: %v", err)
		return errorEntity.NewErrorRes(http.StatusInternalServerError, "Error Changing Password")
	}
	return nil
}

func (u *userSrv) UpdateProfile(req *userEntity.UpdateUserReq) (*userEntity.UpdateUserRes, *errorEntity.ErrorRes) {

	//validate request body
	err := u.vldSrv.Validate(req)
	if err != nil {

		u.log.WithFields(logrus.Fields{
			"error": err,
			"msg":   "Update Profile > Validation service  Failed to validate",
		}).Error("Error validating User")

		return nil, errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !")
	}

	data, err := u.repoSrv.UpdateProfile(req)
	if err != nil {

		u.log.WithFields(logrus.Fields{
			"error": err,
			"msg":   "Update Profile > Repo SRV failed to update ",
		}).Error("Error Updating User")

		return nil, errorEntity.NewErrorRes(http.StatusInternalServerError, "Error Updating")
	}

	user := &userEntity.UpdateUserRes{
		Email:          data.Email,
		Phone:          data.Phone,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		ProfilePicture: data.LastName,
	}

	return user, nil
}

func (u *userSrv) DeleteAccount(req *userEntity.DeleteAccountReq) *errorEntity.ErrorRes {

	//validate request body
	err := u.vldSrv.Validate(req)
	if err != nil {
		u.log.Errorf("Error when validating request body: %v", err)
		return errorEntity.NewErrorRes(http.StatusBadRequest, "Bad Request !")
	}

	err = u.repoSrv.DeactivateAccount(req.UserId)
	if err != nil {
		return errorEntity.NewErrorRes(http.StatusInternalServerError, "Error completing deleting operation")
	}
	return nil
}

func NewUserSrv(log *logrus.Logger, cryptoSrv cryptoService.CryptoSrv, timeSrv timeService.TimeSrv,
	vldSrv validationService.ValidationService, repoSrv userRepository.UserRepo) UserSrv {
	return &userSrv{log: log, cryptoSrv: cryptoSrv, timeSrv: timeSrv, vldSrv: vldSrv, repoSrv: repoSrv}
}
