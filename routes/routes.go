package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/middleware"
	"mini-project-pbi/repositories"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes registers all application routes and dependency injections
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Static folder for uploaded files
	router.Static("/uploads", "./uploads")

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	tokoRepo := repositories.NewTokoRepository(db)
	alamatRepo := repositories.NewAlamatRepository(db)

	// Services
	authService := services.NewAuthService(userRepo, tokoRepo)
	userService := services.NewUserService(userRepo)
	alamatService := services.NewAlamatService(alamatRepo)
	provCityService := services.NewProvCityService()

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	alamatHandler := handlers.NewAlamatHandler(alamatService)
	provCityHandler := handlers.NewProvCityHandler(provCityService)

	// Auth Routes (Public)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// User Routes (Protected JWT)
	user := router.Group("/user", middleware.AuthMiddleware())
	{
		// Profil
		user.GET("", userHandler.GetProfile)
		user.PUT("", userHandler.UpdateProfile)

		// Alamat Kirim
		user.GET("/alamat", alamatHandler.GetMyAlamat)
		user.GET("/alamat/:id", alamatHandler.GetAlamatByID)
		user.POST("/alamat", alamatHandler.CreateAlamat)
		user.PUT("/alamat/:id", alamatHandler.UpdateAlamat)
		user.DELETE("/alamat/:id", alamatHandler.DeleteAlamat)

		// Ping test
		user.GET("/ping", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Auth middleware working",
				"user_id": userID,
			})
		})
	}

	// Province & City Routes (Public)
	provCity := router.Group("/provcity")
	{
		provCity.GET("/listprovincies", provCityHandler.GetListProvincies)
		provCity.GET("/listcities/:prov_id", provCityHandler.GetListCities)
		provCity.GET("/detailprovince/:prov_id", provCityHandler.GetDetailProvince)
		provCity.GET("/detailcity/:city_id", provCityHandler.GetDetailCity)
	}
}
