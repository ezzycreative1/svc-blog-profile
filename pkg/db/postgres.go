package db

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPsqlConnection(dbConfig *config.PsqlConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Schema,
		dbConfig.Password,
	)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}
	if strings.TrimSpace(strings.ToLower(dbConfig.Environment)) == "prod" {
		gormConfig.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             6 * time.Second, // Slow SQL threshold
				LogLevel:                  logger.Error,    // Log level
				IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,           // Disable color
			},
		)
	}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnMaxLifetime))

	return db
}
