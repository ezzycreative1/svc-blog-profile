package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/pkg/db"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"
)

type app struct {
	fiber     *fiber.App
	cfg       *config.Group
	logger    mlog.Logger
	database  *gorm.DB
	validator mvalidator.Validator
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if c, ok := hint.Context.Value(sentry.RequestContextKey).(*fiber.Ctx); ok {
					// You have access to the original Context if it panicked
					fmt.Println(utils.ImmutableString(c.Hostname()))
				}
			}
			fmt.Println(event)
			return event
		},
		Debug:            true,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	// load config
	cfg := config.LoadConfig()

	// load logger
	logger := mlog.New("info", "stdout")

	// init database
	//database := db.NewDatabase(&cfg.Database)
	database := db.NewPsqlConnection(&cfg.BlogDatabase)
	instDB, _ := database.DB()
	defer instDB.Close()

	// init validator
	mValidator := mvalidator.New()

	// create echo app
	f := fiber.New()
	//e.HideBanner = true

	// fill app
	application := app{
		fiber:     f,
		cfg:       cfg,
		logger:    logger,
		database:  database,
		validator: mValidator,
	}

	// set common middleware
	LoadRoute(&application)

	// Start listen server on goroutine
	go func() {
		if err := f.Listen(fmt.Sprintf("0.0.0.0:%v", cfg.Blog.HTTPPort)); err != nil && err != http.ErrServerClosed {
			logger.Info("shutting down the server")
			panic(fmt.Sprintf("server startup panic: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// gracefull shutdown stage ===============================================

	logger.Info("shutdown server...")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := f.Shutdown(); err != nil {
		log.Fatal(err)
	}
	// cleanup app ...
	logger.Info("Running cleanup tasks...")
	logger.Info("Done cleanup tasks...")
	logger.Sync()
}
