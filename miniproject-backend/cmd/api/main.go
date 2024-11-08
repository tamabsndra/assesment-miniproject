package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/tamabsndra/miniproject/miniproject-backend/config"
	_ "github.com/tamabsndra/miniproject/miniproject-backend/docs"
	"github.com/tamabsndra/miniproject/miniproject-backend/handlers"
	"github.com/tamabsndra/miniproject/miniproject-backend/middleware"
	"github.com/tamabsndra/miniproject/miniproject-backend/pkg/database"
	"github.com/tamabsndra/miniproject/miniproject-backend/pkg/redis"
	"github.com/tamabsndra/miniproject/miniproject-backend/repository"
	"github.com/tamabsndra/miniproject/miniproject-backend/services"
)

// @title           Backend API
// @version         1.0
// @description     A REST API using Gin framework with JWT authentication.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

    redisClient, err := redis.NewRedisClient(cfg)
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
    defer redisClient.Close()

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

    tokenService := services.NewTokenService(redisClient, cfg.TokenExpiry, cfg.JWTSecret)
    authService := services.NewAuthService(userRepo, cfg.JWTSecret)
    postService := services.NewPostService(postRepo)

    authHandler := handlers.NewAuthHandler(authService, tokenService)
    postHandler := handlers.NewPostHandler(postService)

	router := gin.Default()

	router.Use(middleware.CORS())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
    {
        api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)
		api.POST("/validate-token", authHandler.ValidateToken)

        protected := api.Group("")
        protected.Use(middleware.AuthMiddleware(cfg.JWTSecret, tokenService))
        {
            protected.POST("/logout", authHandler.Logout)
			protected.GET("/me", authHandler.GetMe)

            protected.POST("/posts", postHandler.Create)
            protected.GET("/posts", postHandler.GetAll)
			protected.GET("/post-detail", postHandler.GetPostDetail)
            protected.GET("/posts/:id", postHandler.GetByID)
			protected.GET("/posts/my/:id", postHandler.GetByUserID)
			protected.PUT("/posts/:id", postHandler.Update)
			protected.DELETE("/posts/:id", postHandler.Delete)
        }
    }

    log.Printf("Server starting on port %s", cfg.ServerPort)
    if err := router.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
