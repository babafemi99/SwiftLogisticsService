package adminRepository

import (
	"context"
	"sls/internal/entity/adminEntity"
	"sls/internal/entity/bikesEntity"
	"sls/internal/entity/riderEntity"
)

type AdminInterface interface {

	//admin authentcation

	Persist(admin *adminEntity.Admin) error
	FetchByEmail(email string) (*adminEntity.Admin, error)
	FetchById(id string) (*adminEntity.Admin, error)
	EditData(id string, admin *adminEntity.AdminAccess) error
	Delete(email string) error
	EditPassword(email, password string) error
	Fetch() ([]*adminEntity.Admin, error)

	//admin to-do

	PersistTodo(todo *adminEntity.Todo) (*adminEntity.Todo, error)
	DeleteTodo(id string) error
	FavoriteTodo(id string) error
	DoTodo(id string) error
	FetchAll() ([]*adminEntity.Todo, error)
	FindByTitle(title string) (*adminEntity.Todo, error)
}

type AdminBikers interface {

	//admin edit-rider

	CreateBike(ctx context.Context, bike *bikesEntity.CreateBike) (*bikesEntity.CreateBike, error)
	EditRider(ctx context.Context, id string, admin *riderEntity.UpdateRiderReqAdmin) (*riderEntity.UpdateRiderReqAdmin, error)

	//Bikers Repository

	EditBike(ctx context.Context, id string, bike *bikesEntity.UpdateBike) (*bikesEntity.UpdateBike, error)
	DeleteBike(ctx context.Context, id string) error
	AssignBikeToRider(ctx context.Context, riderId, bikesId string) error
	UpdateBikeHistory(ctx context.Context, bike *bikesEntity.UpdateBikeHistory) error
	FindBikeById(ctx context.Context, id string) (*adminEntity.AdminViewBikeDetails, error)
	ViewAllBikes(ctx context.Context) ([]*adminEntity.AdminViewBike, error)
	ViewRiderById(ctx context.Context, id string) (*adminEntity.AdminViewRider, error)
	ViewAllRiders(ctx context.Context) ([]*adminEntity.AdminViewAllRider, error)
	ViewPendingRiderById(ctx context.Context, id string) (*adminEntity.AdminViewRiderApplicationById, error)
	ViewAllPendingRiders(ctx context.Context) ([]*adminEntity.AdminViewRiderApplication, error)
	AcceptApplication(ctx context.Context, xx string, application *adminEntity.AcceptRiderApplication) error
	FindApplicationByName(ctx context.Context, name string) (*adminEntity.AdminViewRiderApplicationById, error)
}
