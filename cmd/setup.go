package cmd

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sls/internal/Repository/adminRepository/mongoAdminRepo"
	"sls/internal/Repository/adminRepository/psqlAdminRepo"
	psqlRepo2 "sls/internal/Repository/riderRepository/psqlRepo"
	mongoRepo2 "sls/internal/Repository/settingsRepository/mongoRepo"
	"sls/internal/Repository/userRepository/psqlRepo"
	"sls/internal/datasource/mongoSrc"
	"sls/internal/datasource/psqlSrc"
	"sls/internal/handler/adminHandler"
	"sls/internal/handler/riderHandler"
	"sls/internal/handler/settingsHandler"
	"sls/internal/handler/userHandler"
	"sls/internal/service/adminService"
	"sls/internal/service/cryptoService"
	"sls/internal/service/riderService"
	"sls/internal/service/settingsService"
	"sls/internal/service/timeService"
	"sls/internal/service/tokenService/jwtService"
	"sls/internal/service/userService"
	"sls/internal/service/validationService"
	"time"
)

func Setup() {
	log := logrus.New()
	//config, err := util.LoadConfig(".")
	//if err != nil {
	//	log.Println("unable to load env", err)
	//	return
	//}
	//log.Println(config.DBURL)

	//create postgres Service
	psqlData, err := psqlSrc.NewPsqlSrc(log, "postgres://postgres:sxSMQmFtI2nhe3677a9e@sls.czsuryogwqjo.us-east-1.rds.amazonaws.com:5432/sls")
	if err != nil {
		log.Fatalf("Error Starting Database: %v", err)
	}

	psqlData.LoadDB("./internal/datasource/create.sql")
	log.Info("DB LOADED")
	conn := psqlData.GetConn()
	defer psqlData.CloseConn()

	// create mongo service
	mongoCl, err := mongoSrc.NewMongoSrc("mongodb://localhost:27017/sls")
	if err != nil {
		return
	}
	defer mongoCl.CloseClient()
	log.Println("Mongo connected successfully")

	// creating a Repository services
	userRepo := psqlRepo.NewPsqlUserRepository(conn)
	adminRepo := mongoAdminRepo.NewMongoAdminRepository(mongoCl.Client)
	adminBikeRepo := psqlAdminRepo.NewPsqlAdminBikeRepository(conn)
	settingsRepo := mongoRepo2.NewMongoRepository(mongoCl.Client)
	riderRepo := psqlRepo2.NewPsqlRiderRepo(conn)

	//create utility services

	cryptoSrv := cryptoService.NewCryptoSrv()
	timeSrv := timeService.NewTimesrv()
	validationSrv := validationService.NewValidationSrv(timeSrv)
	tokenSrv := jwtService.NewJWTMaker("ELECTRICITYVIBESONAFREQUENCYTILL")

	// setting up user,admin and settings service
	userSrv := userService.NewUserSrv(log, cryptoSrv, timeSrv, validationSrv, userRepo)
	adminSrv := adminService.NewAdminSrv(adminRepo, cryptoSrv, timeSrv, validationSrv, adminBikeRepo)
	settingSrv := settingsService.NewSettingsService(settingsRepo, validationSrv)
	riderSrv := riderService.NewRiderSrv(timeSrv, cryptoSrv, validationSrv, riderRepo)

	// creating controller services
	userController := userHandler.NewUserController(userSrv, tokenSrv, validationSrv)
	adminController := adminHandler.NewAdminHandler(adminSrv)
	settingsHandler := settingsHandler.NewSettingsHandler(settingSrv)
	riderHandler := riderHandler.NewRiderHttpHandler(riderSrv)

	//create middleware srv
	//middle := middleware.NewMiddleware(tokenSrv)

	//creating and defining routes'
	//router := chiRouter.NewChiRouter(tokenSrv, log)

	router := chi.NewRouter()

	// USER ROUTES
	router.Post("/user/signup", userController.SignUp)
	router.Post("/user/login", userController.Login)
	router.Post("/user/password/reset", userController.ResetPassword)

	//ARouter := router.Group(func(r chi.Router) {
	//	r.Use(middle.Authentication)
	//	r.Use(middle.MapToId)
	//})

	//ARouter.Get("/user/ping/{id}", userController.PING)
	//ARouter.Delete("/user/delete/{id}", userController.DeleteAccount)
	//ARouter.Post("/user/password/change/{id}", userController.ChangePassword)
	//ARouter.Post("/user/update/{id}", userController.UpdateProfile)

	//
	router.Get("/user/ping/{id}", userController.PING)
	router.Delete("/user/delete/{id}", userController.DeleteAccount)
	router.Post("/user/password/change/{id}", userController.ChangePassword)
	router.Post("/user/update/{id}", userController.UpdateProfile)

	//ADMIN ROUTES
	router.Post("/admin/signup", adminController.Create)                           //works
	router.Post("/admin/login", adminController.Login)                             //works
	router.Get("/admin/admins", adminController.FndAll)                            //works
	router.Get("/admin/id/{id}", adminController.FindById)                         //works
	router.Delete("/admin/{id}", adminController.Delete)                           //works
	router.Post("/admin/{id}", adminController.Edit)                               //works
	router.Get("/admin/email/{email}", adminController.FindByEmail)                //works
	router.Patch("/admin/change-password/{email}", adminController.ChangePassword) //works

	//ADMIN TO-DO ROUTES
	router.Post("/admin/todo", adminController.CreateTodo)                 //works
	router.Delete("/admin/todo", adminController.DeleteTodo)               //works
	router.Get("/admin/todo", adminController.FindAllTodo)                 //works
	router.Get("/admin/todo/{title}", adminController.FindTodoByTitle)     //works
	router.Post("/admin/todo/favorite/{id}", adminController.FavoriteTodo) //works
	router.Post("/admin/todo/done/{id}", adminController.MarkAsDone)       //works

	//ADMIN BIKES-RIDER
	router.Post("/admin/rider", adminController.CreateRider)                                //works
	router.Post("/admin/bike", adminController.CreateBike)                                  // works
	router.Post("/admin/rider/{id}", adminController.ModifyRider)                           //works
	router.Get("/admin/rider/{id}", adminController.ViewRiderById)                          //works
	router.Get("/admin/riders", adminController.ViewAllRiders)                              //works
	router.Post("/admin/bike/{id}", adminController.ModifyBike)                             //works
	router.Post("/admin/bike/assign", adminController.AssignBikeToRider)                    //works
	router.Post("/admin/bike/history", adminController.UpdateBikeHistory)                   //works
	router.Get("/admin/bike/{id}", adminController.GetBikeById)                             //works
	router.Delete("/admin/bike/{id}", adminController.DeleteBike)                           // works
	router.Get("/admin/bikes", adminController.GetAllBikes)                                 // works
	router.Get("/admin/pending", adminController.ViewPendingApplications)                   // works
	router.Get("/admin/pending/{id}", adminController.ViewPendingApplicationsById)          // works
	router.Get("/admin/pending/{name}", adminController.ViewPendingRidersApplicationByName) // works
	router.Post("/admin/pending/{id}", adminController.AcceptApplication)                   // works

	//SETTINGS ROUTER
	router.Post("/admin/settings", settingsHandler.CreateSetting)     //works
	router.Post("/admin/settings/{id}", settingsHandler.EditSettings) //works
	router.Get("/admin/settings", settingsHandler.GetSettings)        //works

	router.Post("/rider", riderHandler.SignUp)
	//SERVER SETTINGS
	port := ":8080"
	srv := http.Server{
		Addr:        port,
		Handler:     router,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Errorf("ERROR STARTING SERVER: %v", err)
			os.Exit(1)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Infof("Closing now, We've gotten signal: %v", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
