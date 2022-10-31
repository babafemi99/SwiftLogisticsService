package adminHandler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"sls/internal/entity/adminEntity"
	"sls/internal/entity/bikesEntity"
	"sls/internal/entity/riderEntity"
	"sls/internal/service/adminService"
	"strings"
)

type AdminHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	FindByEmail(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	FndAll(w http.ResponseWriter, r *http.Request)

	// to-do for admin-panel

	CreateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
	FavoriteTodo(w http.ResponseWriter, r *http.Request)
	MarkAsDone(w http.ResponseWriter, r *http.Request)
	FindTodoByTitle(w http.ResponseWriter, r *http.Request)
	FindAllTodo(w http.ResponseWriter, r *http.Request)

	// biker - rider relationships

	CreateBike(w http.ResponseWriter, r *http.Request)
	CreateRider(w http.ResponseWriter, r *http.Request)
	ViewRiderById(w http.ResponseWriter, r *http.Request)
	ViewAllRiders(w http.ResponseWriter, r *http.Request)
	ModifyRider(w http.ResponseWriter, r *http.Request)
	ModifyBike(w http.ResponseWriter, r *http.Request)
	AssignBikeToRider(w http.ResponseWriter, r *http.Request)
	UpdateBikeHistory(w http.ResponseWriter, r *http.Request)
	GetBikeById(w http.ResponseWriter, r *http.Request)
	DeleteBike(w http.ResponseWriter, r *http.Request)
	GetAllBikes(w http.ResponseWriter, r *http.Request)
	ViewPendingApplications(w http.ResponseWriter, r *http.Request)
	ViewPendingApplicationsById(w http.ResponseWriter, r *http.Request)
	ViewPendingRidersApplicationByName(w http.ResponseWriter, r *http.Request)
	AcceptApplication(w http.ResponseWriter, r *http.Request)
}

type adminHandler struct {
	srv adminService.AdminSrv
}

func (a *adminHandler) AcceptApplication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	id := chi.URLParam(r, "id")
	log.Println(id)
	if id == "" {
		log.Println("id is nil")
		return
	}
	var data adminEntity.AcceptRiderApplication
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding rider apllication: %v", err), http.StatusBadRequest)
		return
	}

	err = a.srv.AcceptApplication(id, &data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding rider apllication: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("application accepted")
}

func (a *adminHandler) ViewPendingRidersApplicationByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	name := chi.URLParam(r, "name")
	log.Println(name)
	if name == "" {
		log.Println("name is nil")
		return
	}
	application, err := a.srv.ViewPendingApplicationsByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding rider apllication: %v", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(application)
}

func (a *adminHandler) ViewPendingApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	riders, err := a.srv.ViewAllPendingApplications()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding all pending applications: %v", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(riders)
}

func (a *adminHandler) ViewPendingApplicationsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := chi.URLParam(r, "id")
	log.Println(id)
	if id == "" {
		log.Println("id is nil")
		return
	}
	application, err := a.srv.ViewPendingApplicationsById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding rider apllication: %v", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(application)
}

func (a *adminHandler) ViewAllRiders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	riders, err := a.srv.ViewAllRiders()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding all riders: %v", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(riders)
}

func (a *adminHandler) ViewRiderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]

	rider, err := a.srv.ViewRiderById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding by Id: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rider)
}

func (a *adminHandler) ModifyRider(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var update riderEntity.UpdateRiderReqAdmin
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding into struct: %v", err), http.StatusBadRequest)
		return
	}
	id := strings.Split(r.URL.Path, "/")[3]

	nRider, err := a.srv.EditRider(id, &update)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error editing rider: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nRider)
}

func (a *adminHandler) CreateRider(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newRider riderEntity.CreateRiderReq
	err := json.NewDecoder(r.Body).Decode(&newRider)
	if err != nil {
		http.Error(w, "Error Decoding into struct", http.StatusBadRequest)
		return
	}

	bike, err := a.srv.CreateRider(&newRider)
	if err != nil {
		http.Error(w, "Error creating rider", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bike)
}

func (a *adminHandler) CreateBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newBike bikesEntity.CreateBike
	err := json.NewDecoder(r.Body).Decode(&newBike)
	if err != nil {
		http.Error(w, "Error Decoding into struct", http.StatusBadRequest)
		return
	}

	bike, err := a.srv.PersistBike(&newBike)
	if err != nil {
		http.Error(w, "Error saving bikes", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bike)
}

func (a *adminHandler) ModifyBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var nBike bikesEntity.UpdateBike

	err := json.NewDecoder(r.Body).Decode(&nBike)
	if err != nil {
		http.Error(w, "Error Decoding into struct", http.StatusBadRequest)
		return
	}

	id := strings.Split(r.URL.Path, "/")[3]
	bike, err := a.srv.EditBike(id, &nBike)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error modifying bike: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bike)

}

func (a *adminHandler) AssignBikeToRider(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var bikeData bikesEntity.AssignBikeReq

	err := json.NewDecoder(r.Body).Decode(&bikeData)
	if err != nil {
		http.Error(w, "Error Decoding into struct", http.StatusBadRequest)
		return
	}

	err = a.srv.AssignBikeToRider(bikeData.RiderId, bikeData.BikeId)
	if err != nil {
		http.Error(w, "Error assigning bikes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Assigned successfully, implement bike by Id tho")
}

func (a *adminHandler) UpdateBikeHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var bikeHistory bikesEntity.UpdateBikeHistory

	err := json.NewDecoder(r.Body).Decode(&bikeHistory)
	if err != nil {
		http.Error(w, "Error Decoding into struct", http.StatusBadRequest)
		return
	}

	history, err := a.srv.UpdateBikeHistory(&bikeHistory)
	if err != nil {
		http.Error(w, "Error updating bike history", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history)
}

func (a *adminHandler) GetBikeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]

	bike, err := a.srv.FindBikeById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding bike: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bike)

}

func (a *adminHandler) DeleteBike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]

	err := a.srv.DeleteBike(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting bike: %v", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("deleted")
}

func (a *adminHandler) GetAllBikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	bikes, err := a.srv.FindAllBikes()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding all bike: %v", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bikes)
}

func (a *adminHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var todo adminEntity.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "bad request check again", http.StatusBadRequest)
		return
	}

	created, err := a.srv.CreateTodo(&todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating todo: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(created)
}

func (a *adminHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]

	err := a.srv.DeleteTodo(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *adminHandler) FavoriteTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[4]

	err := a.srv.FavoriteTodo(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error favoriting: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *adminHandler) MarkAsDone(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[4]

	err := a.srv.MarkTodoAsDone(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error favoriting: %v", err), http.StatusInternalServerError)
		return
	}
}

func (a *adminHandler) FindTodoByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	title := strings.Split(r.URL.Path, "/")[3]
	todo, err := a.srv.FindTodoByTitle(title)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding documents: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (a *adminHandler) FindAllTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	todo, err := a.srv.FindAllTodo()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding documents: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (a *adminHandler) FndAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	all, err := a.srv.FindAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding all documents: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(all)
}

func (a *adminHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var LoginReq adminEntity.AdminLoginReq

	err := json.NewDecoder(r.Body).Decode(&LoginReq)
	if err != nil {
		http.Error(w, "bad request check again", http.StatusBadRequest)
		return
	}

	user, err := a.srv.Login(&LoginReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("error login in %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *adminHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var SignUp adminEntity.Admin

	err := json.NewDecoder(r.Body).Decode(&SignUp)
	if err != nil {
		http.Error(w, "bad request check again", http.StatusBadRequest)
		return
	}
	user, err := a.srv.CreateAdmin(&SignUp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *adminHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := strings.Split(r.URL.Path, "/")[2]

	err := a.srv.Delete(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("no user with such email: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *adminHandler) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	log.Println("inside here")
	id := strings.Split(r.URL.Path, "/")[2]

	var UpdateData adminEntity.AdminAccess

	err := json.NewDecoder(r.Body).Decode(&UpdateData)
	if err != nil {
		http.Error(w, "bad request check again", http.StatusBadRequest)
		return
	}

	user, err := a.srv.Edit(id, &UpdateData)
	if err != nil {
		http.Error(w, "Error Updating", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *adminHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	email := strings.Split(r.URL.Path, "/")[3]

	user, err := a.srv.FindByEmail(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("no user with such email: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (a *adminHandler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]

	user, err := a.srv.FindById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("no user with such id: %v", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (a *adminHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	email := strings.Split(r.URL.Path, "/")[3]

	var req adminEntity.ChangePasswordReq
	req.Email = email

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request check again", http.StatusBadRequest)
		return
	}

	err = a.srv.ChangePassword(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error changing password: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Password changed successfully")
}

func NewAdminHandler(srv adminService.AdminSrv) AdminHandler {
	return &adminHandler{srv: srv}
}
