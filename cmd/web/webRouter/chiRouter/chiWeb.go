package chiRouter

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sls/cmd/web/webRouter"
	"sls/internal/service/tokenService"
	"time"
)

var (
	router = chi.NewRouter()
)

type chiRouter struct {
	tokenSrv tokenService.TokenSrv
	log      *logrus.Logger
}

func (c *chiRouter) USE(f func(next http.Handler) http.Handler) {
	router.Use(f)
}

func NewChiRouter(tokenSrv tokenService.TokenSrv, log *logrus.Logger) webRouter.WebRouter {
	return &chiRouter{tokenSrv: tokenSrv, log: log}
}

func (c chiRouter) DELETE(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	router.Delete(uri, f)
}

func (c chiRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	router.Get(uri, f)
}

func (c chiRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	router.Post(uri, f)
}

func (c chiRouter) SERVE(port string) {
	srv := http.Server{
		Addr:        port,
		Handler:     router,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			c.log.Errorf("ERROR STARTING SERVER: %v", err)
			os.Exit(1)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	c.log.Infof("Closing now, We've gotten signal: %v\n", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
