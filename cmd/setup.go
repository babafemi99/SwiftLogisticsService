package cmd

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sls/cmd/middleware"
	"sls/internal/Repository/userRepository/psqlRepo"
	"sls/internal/datasource/psqlSrc"
	"sls/internal/handler/userHandler"
	"sls/internal/service/cryptoService"
	"sls/internal/service/timeService"
	"sls/internal/service/tokenService/jwtService"
	"sls/internal/service/userService"
	"sls/internal/service/validationService"
	"time"
)

func Setup() {
	log := logrus.New()

	// Creating a data source
	psqlData, err := psqlSrc.NewPsqlSrc(log, "postgres://postgres:mysecretpassword@localhost:5432/slsstore")
	if err != nil {
		log.Fatalf("Error Starting Database: %v", err)
	}

	psqlData.LoadDB("./internal/datasource/create.sql")
	log.Info("DB LOADED")
	conn := psqlData.GetConn()
	defer psqlData.CloseConn()

	// creating a Repo service
	repository := psqlRepo.NewPsqlRepository(conn)

	// setting up user service
	cryptoSrv := cryptoService.NewCryptoSrv()
	timeSrv := timeService.NewTimesrv()
	validationSrv := validationService.NewValidationSrv()
	userSrv := userService.NewUserSrv(log, cryptoSrv, timeSrv, validationSrv, repository)

	// creating token service
	tokenSrv := jwtService.NewJWTMaker("ELECTRICITYVIBESONAFREQUENCYTILL")

	// creating controller service
	userController := userHandler.NewUserController(userSrv, tokenSrv, validationSrv)

	//create middleware srv
	middle := middleware.NewMiddleware(tokenSrv)

	//creating router
	//router := chiRouter.NewChiRouter(tokenSrv, log)

	router := chi.NewRouter()

	// USER ROUTES

	router.Post("/user/signup", userController.SignUp)
	router.Post("/user/login", userController.Login)
	router.Post("/user/password/reset", userController.ResetPassword)

	ARouter := router.Group(func(r chi.Router) {
		r.Use(middle.Authentication)
		r.Use(middle.MapToId)
	})

	ARouter.Get("/user/ping/{id}", userController.PING)
	ARouter.Delete("/user/delete/{id}", userController.DeleteAccount)
	ARouter.Post("/user/password/change/{id}", userController.ChangePassword)
	ARouter.Post("/user/update/{id}", userController.UpdateProfile)

	port := ":9090"
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
