package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/c-kissel/SellerMDM.git/internal/api/v1"
	"github.com/c-kissel/SellerMDM.git/internal/core/config"
	"github.com/c-kissel/SellerMDM.git/internal/service"
	"github.com/c-kissel/SellerMDM.git/internal/storage"
	"github.com/c-kissel/SellerMDM.git/specs"
	"github.com/go-chi/cors"

	"golang.org/x/sync/errgroup"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		err         error
		ctx, cancel = signal.NotifyContext(
			context.Background(),
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
	)
	defer cancel()

	SetupLogging()
	logrus.Info("Started")
	logrus.Debug("Started with args: ", os.Args)

	cfg, err := config.InitConfig(os.Args)
	if err != nil {
		logrus.Error("failed to init config: ", err.Error())
		return
	}
	logrus.Debug("Got configuration: ", cfg)

	// Storage
	var store service.Storer = storage.NewStorage()

	// Service
	var srv v1.Server = service.NewService(&store)

	// API Handler
	apiServer := v1.NewAPI(&srv)

	err = startHTTPServer(ctx, cfg, apiServer)
	if err != nil {
		logrus.Fatalf("starting server: %s", err.Error())
	}
}

func startHTTPServer(
	ctx context.Context,
	cfg *config.Config,
	apiServer specs.ServerInterface,
	middlewares ...specs.MiddlewareFunc,
) error {
	handler := specs.HandlerWithOptions(apiServer, specs.ChiServerOptions{
		BaseURL:     cfg.BasePath,
		Middlewares: middlewares,
	})

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedOrigins:   cfg.Allowed.Hosts, // Use this to allow specific origin hosts
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Handle("/*", handler)

	port := os.Getenv("PORT")

	httpServer := http.Server{
		Addr:    port,
		Handler: router,
	}

	group, ctx := errgroup.WithContext(ctx)

	logrus.Infof("Listening on %s", port)

	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		return httpServer.Shutdown(ctx)
	})

	return group.Wait()
}

func SetupLogging() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}
