package adminService

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"sls/internal/Repository/adminRepository"
	"sls/internal/entity/adminEntity"
	"sls/internal/entity/bikesEntity"
	"sls/internal/entity/riderEntity"
	"sls/internal/service/cryptoService"
	"sls/internal/service/riderService"
	"sls/internal/service/timeService"
	"sls/internal/service/validationService"
	"time"
)

const (
	ADMINACCESSLEVEL = "ADMIN"
)

type AdminSrv interface {
	Login(req *adminEntity.AdminLoginReq) (*adminEntity.AdminAccess, error)
	CreateAdmin(admin *adminEntity.Admin) (*adminEntity.Admin, error)
	Delete(id string) error
	Edit(id string, req *adminEntity.AdminAccess) (*adminEntity.AdminAccess, error)
	FindById(id string) (*adminEntity.AdminAccess, error)
	FindByEmail(email string) (*adminEntity.AdminAccess, error)
	ChangePassword(req *adminEntity.ChangePasswordReq) error
	FindAll() ([]*adminEntity.AdminList, error)

	// TODO SERVICE

	CreateTodo(todo *adminEntity.Todo) (*adminEntity.Todo, error)
	DeleteTodo(id string) error
	FindAllTodo() ([]*adminEntity.Todo, error)
	FindTodoByTitle(title string) (*adminEntity.Todo, error)
	FavoriteTodo(id string) error
	MarkTodoAsDone(id string) error

	// Rider -Bikers Service
	// create Rider: inject the rider service and create a rider :)

	CreateRider(req *riderEntity.CreateRiderReq) (*riderEntity.CreateRiderRes, error)
	EditRider(id string, req *riderEntity.UpdateRiderReqAdmin) (*riderEntity.UpdateRiderReqAdmin, error)
	PersistBike(bike *bikesEntity.CreateBike) (*bikesEntity.CreateBike, error)
	EditBike(id string, bike *bikesEntity.UpdateBike) (*bikesEntity.UpdateBike, error)
	AssignBikeToRider(id, riderId string) error
	UpdateBikeHistory(bike *bikesEntity.UpdateBikeHistory) (*bikesEntity.UpdateBike, error)
	FindBikeById(id string) (*adminEntity.AdminViewBikeDetails, error)
	DeleteBike(id string) error
	FindAllBikes() ([]*adminEntity.AdminViewBike, error)
	ViewRiderById(id string) (*adminEntity.AdminViewRider, error)
	ViewAllRiders() ([]*adminEntity.AdminViewAllRider, error)
	ViewAllPendingApplications() ([]*adminEntity.AdminViewRiderApplication, error)
	ViewPendingApplicationsById(id string) (*adminEntity.AdminViewRiderApplicationById, error)
	ViewPendingApplicationsByName(name string) (*adminEntity.AdminViewRiderApplicationById, error)
	AcceptApplication(id string, application *adminEntity.AcceptRiderApplication) error
}
type adminSrv struct {
	adminRepo      adminRepository.AdminInterface
	cryptoSrv      cryptoService.CryptoSrv
	timeSrv        timeService.TimeSrv
	vldSrv         validationService.ValidationService
	adminBikerRepo adminRepository.AdminBikers
	riderSrv       riderService.RiderSrv
}

func (a *adminSrv) AcceptApplication(id string, application *adminEntity.AcceptRiderApplication) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	application.AccountStatus = "ACTIVE"
	application.DateJoined = a.timeSrv.Now().Format(time.RFC3339)
	return a.adminBikerRepo.AcceptApplication(ctx, id, application)
}

func (a *adminSrv) ViewPendingApplicationsByName(name string) (*adminEntity.AdminViewRiderApplicationById, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	return a.adminBikerRepo.FindApplicationByName(ctx, name)
}

func (a *adminSrv) ViewAllPendingApplications() ([]*adminEntity.AdminViewRiderApplication, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	return a.adminBikerRepo.ViewAllPendingRiders(ctx)
}

func (a *adminSrv) ViewPendingApplicationsById(id string) (*adminEntity.AdminViewRiderApplicationById, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	return a.adminBikerRepo.ViewPendingRiderById(ctx, id)
}

func (a *adminSrv) ViewAllRiders() ([]*adminEntity.AdminViewAllRider, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	return a.adminBikerRepo.ViewAllRiders(ctx)
}

func (a *adminSrv) ViewRiderById(id string) (*adminEntity.AdminViewRider, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	return a.adminBikerRepo.ViewRiderById(ctx, id)
}

func (a *adminSrv) EditRider(id string, req *riderEntity.UpdateRiderReqAdmin) (*riderEntity.UpdateRiderReqAdmin, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	err := a.vldSrv.Validate(req)
	if err != nil {
		return nil, err
	}
	req.DateUpdated = a.timeSrv.Now().Format(time.RFC3339)
	riderRes, err := a.adminBikerRepo.EditRider(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return riderRes, nil
}

func (a *adminSrv) CreateRider(req *riderEntity.CreateRiderReq) (*riderEntity.CreateRiderRes, error) {
	create, err := a.riderSrv.Create(req)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (a *adminSrv) FindAllBikes() ([]*adminEntity.AdminViewBike, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	return a.adminBikerRepo.ViewAllBikes(ctx)
}

func (a *adminSrv) PersistBike(bike *bikesEntity.CreateBike) (*bikesEntity.CreateBike, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	err := a.vldSrv.Validate(bike)
	if err != nil {
		return nil, err
	}

	bike.BikeId = uuid.New().String()
	createBike, err := a.adminBikerRepo.CreateBike(ctx, bike)
	if err != nil {
		return nil, err
	}

	return createBike, nil
}

func (a *adminSrv) EditBike(id string, bike *bikesEntity.UpdateBike) (*bikesEntity.UpdateBike, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	err := a.vldSrv.Validate(bike)
	if err != nil {
		return nil, err
	}

	editBike, err := a.adminBikerRepo.EditBike(ctx, id, bike)
	if err != nil {
		return nil, err
	}
	return editBike, nil
}

func (a *adminSrv) AssignBikeToRider(riderId, bikeId string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	return a.adminBikerRepo.AssignBikeToRider(ctx, riderId, bikeId)
}

func (a *adminSrv) UpdateBikeHistory(bike *bikesEntity.UpdateBikeHistory) (*bikesEntity.UpdateBike, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	err := a.vldSrv.Validate(bike)
	if err != nil {
		return nil, err
	}

	err = a.adminBikerRepo.UpdateBikeHistory(ctx, bike)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (a *adminSrv) FindBikeById(id string) (*adminEntity.AdminViewBikeDetails, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	return a.adminBikerRepo.FindBikeById(ctx, id)
}

func (a *adminSrv) DeleteBike(id string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	return a.adminBikerRepo.DeleteBike(ctx, id)
}

func (a *adminSrv) FindAll() ([]*adminEntity.AdminList, error) {

	var AdminList []*adminEntity.AdminList

	admins, err := a.adminRepo.Fetch()
	if err != nil {
		return nil, err
	}
	for _, admin := range admins {

		res := adminEntity.AdminList{
			FirstName:      admin.FirstName,
			LastName:       admin.LastName,
			Email:          admin.Email,
			ProfilePicture: admin.ProfilePicture,
		}

		AdminList = append(AdminList, &res)
	}

	return AdminList, nil
}

func (a *adminSrv) Login(req *adminEntity.AdminLoginReq) (*adminEntity.AdminAccess, error) {
	err := a.vldSrv.Validate(req)
	if err != nil {
		return nil, err
	}

	userDetails, err := a.adminRepo.FetchByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = a.cryptoSrv.ComparePassword(userDetails.Password, req.Password)
	if err != nil {
		return nil, err
	}

	user := &adminEntity.AdminAccess{
		Id:             userDetails.Id,
		FirstName:      userDetails.Email,
		LastName:       userDetails.LastName,
		Email:          userDetails.Email,
		PhoneNumber:    userDetails.PhoneNumber,
		Gender:         userDetails.Gender,
		Position:       userDetails.Position,
		ProfilePicture: userDetails.ProfilePicture,
		AccessLevel:    userDetails.AccessLevel,
		UpdatedAt:      userDetails.UpdatedAt,
		Role:           userDetails.Role,
	}

	return user, nil
}

func (a *adminSrv) CreateAdmin(admin *adminEntity.Admin) (*adminEntity.Admin, error) {
	err := a.vldSrv.Validate(admin)
	if err != nil {
		return nil, err
	}

	_, err = a.adminRepo.FetchByEmail(admin.Email)
	if err == nil {
		return nil, errors.New("user already Exists")
	}

	password, err := a.cryptoSrv.HashPassword(admin.Password)
	if err != nil {
		return nil, err
	}

	admin.Id = uuid.New().String()
	admin.Password = password
	admin.AccessLevel = ADMINACCESSLEVEL
	admin.CreatedAt = a.timeSrv.Now().Format(time.RFC3339)

	err = a.adminRepo.Persist(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil

}

func (a *adminSrv) Delete(email string) error {
	return a.adminRepo.Delete(email)
}

func (a *adminSrv) Edit(id string, req *adminEntity.AdminAccess) (*adminEntity.AdminAccess, error) {
	err := a.vldSrv.Validate(req)
	if err != nil {
		return nil, err
	}

	req.UpdatedAt = time.Now().Format(time.RFC3339)
	log.Println(id)
	err = a.adminRepo.EditData(id, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (a *adminSrv) FindById(id string) (*adminEntity.AdminAccess, error) {
	userDetails, err := a.adminRepo.FetchById(id)
	if err != nil {
		return nil, err
	}

	user := &adminEntity.AdminAccess{
		Id:             userDetails.Id,
		FirstName:      userDetails.Email,
		LastName:       userDetails.LastName,
		Email:          userDetails.Email,
		PhoneNumber:    userDetails.PhoneNumber,
		Gender:         userDetails.Gender,
		Position:       userDetails.Position,
		ProfilePicture: userDetails.ProfilePicture,
		AccessLevel:    userDetails.AccessLevel,
		UpdatedAt:      userDetails.UpdatedAt,
		Role:           userDetails.Role,
	}
	return user, nil
}

func (a *adminSrv) FindByEmail(email string) (*adminEntity.AdminAccess, error) {

	userDetails, err := a.adminRepo.FetchByEmail(email)
	if err != nil {
		return nil, err
	}

	user := &adminEntity.AdminAccess{
		Id:             userDetails.Id,
		FirstName:      userDetails.Email,
		LastName:       userDetails.LastName,
		Email:          userDetails.Email,
		PhoneNumber:    userDetails.PhoneNumber,
		Gender:         userDetails.Gender,
		Position:       userDetails.Position,
		ProfilePicture: userDetails.ProfilePicture,
		AccessLevel:    userDetails.AccessLevel,
		UpdatedAt:      userDetails.UpdatedAt,
		Role:           userDetails.Role,
	}
	return user, nil
}

func (a *adminSrv) ChangePassword(req *adminEntity.ChangePasswordReq) error {

	// validate the request first
	err := a.vldSrv.Validate(req)
	if err != nil {
		return err
	}

	// fetch the user with that particular email
	user, err := a.adminRepo.FetchByEmail(req.Email)
	if err != nil {
		return err
	}

	//compare the old password with the user password
	err = a.cryptoSrv.ComparePassword(user.Password, req.OldPassword)
	if err != nil {
		log.Println("error1")
		return err
	}

	pwd, err := a.cryptoSrv.HashPassword(req.NewPassword)
	if err != nil {
		log.Println("error2")
		return err
	}

	//update the new password
	err = a.adminRepo.EditPassword(req.Email, pwd)
	if err != nil {
		return err
	}

	return nil

}

// to-do impl for the admin panel

func (a *adminSrv) CreateTodo(todo *adminEntity.Todo) (*adminEntity.Todo, error) {
	err := a.vldSrv.Validate(todo)
	if err != nil {
		return nil, err
	}

	todo.Status = "UNDONE"
	todo.Id = uuid.New().String()

	todo, err = a.adminRepo.PersistTodo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (a *adminSrv) DeleteTodo(id string) error {
	return a.adminRepo.DeleteTodo(id)
}

func (a *adminSrv) FindAllTodo() ([]*adminEntity.Todo, error) {
	return a.adminRepo.FetchAll()
}

func (a *adminSrv) FindTodoByTitle(title string) (*adminEntity.Todo, error) {
	return a.adminRepo.FindByTitle(title)
}

func (a *adminSrv) FavoriteTodo(id string) error {
	return a.adminRepo.FavoriteTodo(id)
}

func (a *adminSrv) MarkTodoAsDone(id string) error {
	return a.adminRepo.DoTodo(id)
}

func NewAdminSrv(adminRepo adminRepository.AdminInterface, cryptoSrv cryptoService.CryptoSrv, timeSrv timeService.TimeSrv,
	vldSrv validationService.ValidationService, bikers adminRepository.AdminBikers) AdminSrv {
	return &adminSrv{adminRepo: adminRepo, cryptoSrv: cryptoSrv, timeSrv: timeSrv, vldSrv: vldSrv, adminBikerRepo: bikers}
}
