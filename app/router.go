package main

import (
	"net/http"

	"github.com/ezzycreative1/svc-blog-profile/app/v1/handler"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/usecase"
	"github.com/ezzycreative1/svc-blog-profile/internal/infrastructure/persistence/postgresql"
	"github.com/ezzycreative1/svc-blog-profile/internal/infrastructure/persistence/redis"
	"github.com/ezzycreative1/svc-blog-profile/pkg/web"
	"github.com/gofiber/fiber/v2"
)

const (
	keyTransaction = "blog-ctx"
	timeout        = 10
)

func LoadRoute(app *app) {
	// init dependency
	userRepos := postgresql.NewUserRepository(app.database, keyTransaction, timeout)
	queueUserRepos := redis.NewQueueUserRepo(*app.redis)
	userUCase := usecase.NewUserUsecase(
		userRepos,
		queueUserRepos,
		app.logger,
	)
	userHandler := handler.NewUserHandler(
		userUCase,
		app.validator,
		app.logger,
		*app.cfg,
	)

	// init additional middleware here or directly in route (ex. JWT, api key)
	// ...

	// set route =============================================================
	// route for check health
	app.fiber.Get("v1/health/ping", func(c *fiber.Ctx) error {
		return web.ResponseFormatter(c, http.StatusOK, "Success", map[string]any{"status": "ok"}, nil)
	})

	g := app.fiber.Group("/v1/blog")
	//router role
	g.Post("/user", userHandler.Register)
	g.Post("/user/login", userHandler.Login)
	// g.GET("/roles", pokemonHandler.FetchRoles)
	// g.PUT("/role/:id", pokemonHandler.UpdateRole)
	// g.GET("/role/:id", pokemonHandler.GetRoleByID)
	// g.DELETE("/role/:id", pokemonHandler.DeleteRole)
	//router user
	// g.POST("/user", pokemonHandler.StoreUser)
	// g.GET("/users", pokemonHandler.FetchUsers)
	// g.PUT("/user/:id", pokemonHandler.UpdateUser)
	// g.GET("/user/:id", pokemonHandler.GetUserByID)
	// g.DELETE("/user/:id", pokemonHandler.DeleteUser)
	// g.POST("/user/login", pokemonHandler.Login)

	// g.GET("/list", pokemonHandler.FetchPokemons)
	// g.POST("/create", pokemonHandler.StorePokemon)
}
