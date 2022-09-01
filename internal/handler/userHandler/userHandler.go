package userHandler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sls/internal/entity/errorEntity"
	"sls/internal/entity/handlerEntity"
	"sls/internal/entity/userEntity"
	"sls/internal/service/tokenService"
	"sls/internal/service/userService"
	"sls/internal/service/validationService"
	"strings"
	"time"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
	PING(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userSrv  userService.UserSrv
	tokenSrv tokenService.TokenSrv
	vldSrv   validationService.ValidationService
}

func NewUserController(userSrv userService.UserSrv, tokenSrv tokenService.TokenSrv,
	vldSrv validationService.ValidationService) UserHandler {
	return &userController{userSrv: userSrv, tokenSrv: tokenSrv, vldSrv: vldSrv}
}

func (u *userController) PING(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("PONG !!")
}

func (u *userController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var signUp userEntity.CreateUser

	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Bad Request, "+
			"Error validating Data"))
		return
	}

	err = u.vldSrv.Validate(&signUp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError,
			"Bad Request"+err.Error()))
		return
	}

	up, res := u.userSrv.SignUp(&signUp)
	if res != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	token, err := u.tokenSrv.CreateToken(up.UserId, time.Minute*20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Error Creating Token"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handlerEntity.NewClientAccess(token, up))
}

func (u *userController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var data userEntity.LoginReq

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Bad Request"))
		return
	}

	login, loginErr := u.userSrv.Login(&data)
	if loginErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(loginErr)
		return
	}
	fmt.Println(login.UserId)
	token, err := u.tokenSrv.CreateToken(login.UserId, time.Minute*20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Error Creating Token"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handlerEntity.NewClientAccess(token, login))
}

func (u *userController) ChangePassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var req userEntity.ChangePassword

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Bad Request"))
		return
	}
	id := strings.Split(r.URL.Path, "/")[4]
	parse, err2 := uuid.Parse(id)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Invalid Id Format"))
		return
	}
	req.UserId = parse

	err = u.vldSrv.Validate(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError,
			"Bad Request: Error Validating Request"))
		return
	}

	changePasswordErr := u.userSrv.ChangePassword(&req)
	if changePasswordErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(changePasswordErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handlerEntity.NewSuccessMessage("password changed successfully"))
}

func (u *userController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var req userEntity.ResetPassword

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Bad Request"))
		return
	}

	err = u.vldSrv.Validate(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError,
			"Bad Request: Error Validating Request"))
		return
	}

	resetErr := u.userSrv.ResetPassword(&req)
	if resetErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resetErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handlerEntity.NewSuccessMessage("password changed successfully"))
}

func (u *userController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var req userEntity.UpdateUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Bad Request"))
		return
	}

	id := strings.Split(r.URL.Path, "/")[3]
	parse, err2 := uuid.Parse(id)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Invalid Id Format"))
		return
	}
	req.UserId = parse

	err = u.vldSrv.Validate(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError,
			"Bad Request: Error Validating Request"))
		return
	}

	update, updateErr := u.userSrv.UpdateProfile(&req)
	if updateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(updateErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&update)

}

func (u *userController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	id := strings.Split(r.URL.Path, "/")[3]
	var req userEntity.DeleteAccountReq
	parse, err2 := uuid.Parse(id)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorEntity.NewErrorRes(http.StatusInternalServerError, "Invalid Id Format"))
		return
	}
	req.UserId = parse

	err := u.userSrv.DeleteAccount(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handlerEntity.NewSuccessMessage("Delete Operation Successfully"))

}
