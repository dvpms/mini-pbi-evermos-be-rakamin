package database

import (
	"log"
	"os"
	"time"

	"mini-project-pbi/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	_ = godotenv.Load()
	if os.Getenv("DATABASE_URL") == "" {
		_ = godotenv.Load("../.env")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in environment")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connection established successfully")

	// Run auto migrations
	Migrate(db)

	return db
}

// Migrate executes auto-migration for all database entities
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Toko{},
		&models.Alamat{},
		&models.Category{},
		&models.Produk{},
		&models.FotoProduk{},
		&models.LogProduk{},
		&models.Transaksi{},
		&models.DetailTransaksi{},
	)
	if err != nil {
		log.Fatalf("Database auto-migration failed: %v", err)
	}
	log.Println("Database auto-migration completed successfully")
}


