package riderHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sls/internal/entity/riderEntity"
	"sls/internal/service/riderService"
)

type riderHttpHandler struct {
	riderSrv riderService.RiderSrv
}

func NewRiderHttpHandler(riderSrv riderService.RiderSrv) *riderHttpHandler {
	return &riderHttpHandler{riderSrv: riderSrv}
}

func (rider *riderHttpHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var newBike riderEntity.CreateRiderReq
	err := json.NewDecoder(r.Body).Decode(&newBike)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Decoding into struct: %v", err), http.StatusBadRequest)
		return
	}

	create, err := rider.riderSrv.Create(&newBike)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating new rider: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(create)

}

//func (rider *riderHttpHandler) Login(w http.ResponseWriter, r *http.Request) {
//
//}
//func (rider *riderHttpHandler) EditProfile(w http.ResponseWriter, r *http.Request) {
//
//}
//func (rider *riderHttpHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
//
//}
//func (rider *riderHttpHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
//
//}
//func (rider *riderHttpHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
//
//}
