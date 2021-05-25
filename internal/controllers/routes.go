package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/trucktrace/internal/repository"
	service "github.com/trucktrace/internal/services"
)

type Controllers struct {
	services   *service.Services
	repository *repository.Repository
}

func GetNewController(services *service.Services, repository *repository.Repository) *Controllers {
	return &Controllers{services, repository}
}

func (controllers *Controllers) InitRoutes() *gin.Engine {
	router := gin.New()

	//applying global and general middlewares CORS
	// router.Use(middlewares.CORSMiddleware())
	groupControllers := GetGroupControllers()
	// Recovery middleware used to continue application run after panic oc log fatal
	router.Use(gin.Recovery())

	// creating Api Group Routes
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signup", groupControllers. )
			auth.POST("/signin")
			auth.GET("/logout", controllers.Logout)
		}

		resetPassword := api.Group("/resetPassword")
		{
			resetPassword.POST("", controllers.)
			resetPassword.GET("/:token", controllers.GetResetPassword)
		}

		changePassword := api.Group("/changePassword")
		{
			changePassword.POST("", controllers.PostChangePassword)
		}
	}

	rsaConfig := router.Group("/rsa")
	{
		rsaConfig.GET("/getconfig", controllers.GetConfig)
		// rsaConfig.GET("/tokenconfig", controllers.GetConfigByKeyValue)
		// }

		return router
	}
}
