package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-architecture-mysql/api/config"
	"go-architecture-mysql/api/controllers"
	"go-architecture-mysql/api/middlewares"
	"go-architecture-mysql/api/repositories/repoimpl"
	"go-architecture-mysql/api/services/serviceimpl"
	_ "go-architecture-mysql/docs"
	"gorm.io/gorm"
)

func CreateHandler(db *gorm.DB) {
	r := gin.New()

	userRepo := repoimpl.CreateUserRepository(db)

	userService := serviceimpl.CreateUserService(userRepo)
	authService := serviceimpl.CreateAuthService(userRepo)

	r.Use(middlewares.SetupCorsMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	controllers.CreateAuthRoutes(api.Group("/Auth"), authService, userService)
	controllers.CreateUserRoutes(api.Group("/User"), userService)

	err := r.Run(fmt.Sprintf(":%v", config.Config.Server.Port))

	if err != nil {
		log.Fatal(err)
		return
	}
}