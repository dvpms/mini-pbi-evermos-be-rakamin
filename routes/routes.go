package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/repositories"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes handles dependency injection and registers all modular application routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Static directory for uploaded image assets
	router.Static("/uploads", "./uploads")

	// 1. Data Access Layer (Repositories)
	userRepo := repositories.NewUserRepository(db)
	tokoRepo := repositories.NewTokoRepository(db)
	alamatRepo := repositories.NewAlamatRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	produkRepo := repositories.NewProdukRepository(db)
	transaksiRepo := repositories.NewTransaksiRepository(db)

	// 2. Business Logic Layer (Services)
	authService := services.NewAuthService(userRepo, tokoRepo)
	userService := services.NewUserService(userRepo)
	alamatService := services.NewAlamatService(alamatRepo)
	provCityService := services.NewProvCityService()
	tokoService := services.NewTokoService(tokoRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	produkService := services.NewProdukService(produkRepo, tokoRepo)
	transaksiService := services.NewTransaksiService(transaksiRepo, alamatRepo, produkRepo)

	// 3. Presentation Layer (Handlers)
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	alamatHandler := handlers.NewAlamatHandler(alamatService)
	provCityHandler := handlers.NewProvCityHandler(provCityService)
	tokoHandler := handlers.NewTokoHandler(tokoService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	produkHandler := handlers.NewProdukHandler(produkService)
	transaksiHandler := handlers.NewTransaksiHandler(transaksiService)

	// 4. Modular Route Registrations
	RegisterAuthRoutes(router, authHandler)
	RegisterUserRoutes(router, userHandler, alamatHandler)
	RegisterTokoRoutes(router, tokoHandler)
	RegisterCategoryRoutes(router, categoryHandler)
	RegisterProductRoutes(router, produkHandler)
	RegisterTrxRoutes(router, transaksiHandler)
	RegisterProvCityRoutes(router, provCityHandler)
}
