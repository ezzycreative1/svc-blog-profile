package config

import (
	"log"

	"github.com/ezzycreative1/svc-blog-profile/pkg/envar"
	"github.com/joho/godotenv"
)

type Group struct {
	Server Server `json:"server,omitempty"`
	//Database     Database   `json:"database,omitempty"`
	BlogDatabase PsqlConfig `json:"blog_database,omitempty"`
	Redis        Redis      `json:"redis,omitempty"`
	Jwt          Jwt        `json:"jwt,omitempty"`
	Blog         BlogConfig `json:"pokemon,omitempty"`
}

type Server struct {
	ENV string `json:"env"`
}

func LoadConfig() *Group {
	if err := godotenv.Load(); err != nil {
		// in prod we will not use this,use os env instead
		log.Print(".env notfound")
	}

	env = envar.GetEnv("ENV", "dev")

	return &Group{
		Server: Server{
			ENV: env,
		},
		Blog: BlogConfig{
			HTTPPort: envar.GetEnv("HTTP_PORT", 8080),
		},
		// Database: Database{
		// 	Engine:          envar.GetEnv("DATABASE_ENGINE", "mysqli"),
		// 	Host:            envar.GetEnv("DATABASE_HOST", "localhost"),
		// 	Port:            envar.GetEnv("DATABASE_PORT", 3306),
		// 	Username:        envar.GetEnv("DATABASE_USERNAME", "root"),
		// 	Password:        envar.GetEnv("DATABASE_PASSWORD", ""),
		// 	Schema:          envar.GetEnv("DATABASE_SCHEMA", "inventory"),
		// 	MaxIdle:         envar.GetEnv("DATABASE_MAX_IDLE", 20),
		// 	MaxConn:         envar.GetEnv("DATABASE_MAX_CONN", 100),
		// 	ConnMaxLifetime: envar.GetEnv("DATABASE_CONN_LIFETIME", 180),
		// 	Environment:     env,
		// },
		BlogDatabase: PsqlConfig{
			Host:            envar.GetEnv("POSTGRES_HOST", "localhost"),
			Port:            envar.GetEnv("POSTGRES_PORT", 5432),
			User:            envar.GetEnv("POSTGRES_USERNAME", "root"),
			Password:        envar.GetEnv("POSTGRES_PASSWORD", ""),
			Schema:          envar.GetEnv("POSTGRES_SCHEMA", "myblog"),
			MaxIdleConns:    envar.GetEnv("POSTGRES_MAX_IDLE", 20),
			MaxOpenConns:    envar.GetEnv("POSTGRES_MAX_CONN", 100),
			ConnMaxLifetime: envar.GetEnv("POSTGRES_CONN_LIFETIME", 180),
			Environment:     env,
		},
		Redis: Redis{
			Host:     envar.GetEnv("REDIS_HOST", "localhost"),
			Port:     envar.GetEnv("REDIS_PORT", 31113),
			Username: envar.GetEnv("REDIS_USERNAME", ""),
			Password: envar.GetEnv("REDIS_PASSWORD", ""),
			DB:       envar.GetEnv("REDIS_DB", 4),
			UseTLS:   envar.GetEnv("REDIS_USE_TLS", false),
		},
		Jwt: Jwt{
			Secret:           envar.GetEnv("JWT_SECRET", "secret"),
			AccessExpireMin:  envar.GetEnv("JWT_ACCESS_EXPIRE_MIN", 15),
			RefreshExpireMin: envar.GetEnv("JWT_REFRESH_EXPIRE_MIN", 60*24*3),
		},
	}
}
