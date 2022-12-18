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
	"github.com/gofiber/fiber/v2"
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
	// load config
	cfg := config.LoadConfig()

	// load logger
	logger := mlog.New("info", "stdout")

	// init database
	database := db.NewDatabase(&cfg.Database)
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
			panic(fmt.Sprintf("echo server startup panic: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// gracefull shutdown stage ===============================================

	logger.Info("shutdown echo server...")
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
