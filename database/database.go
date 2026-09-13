package database

import (
	"fmt"
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
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
			log.Fatal("Konfigurasi database di file .env belum lengkap. Harap pastikan DB_USER, DB_PASS, DB_HOST, DB_PORT, dan DB_NAME telah diisi.")
		}

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
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

	// Run auto migrations conditionally if enabled in environment
	if os.Getenv("AUTO_MIGRATE") == "true" {
		Migrate(db)
	} else {
		log.Println("Database auto-migration skipped (AUTO_MIGRATE is not true)")
	}

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

	// Seed initial categories if empty
	var catCount int64
	db.Model(&models.Category{}).Count(&catCount)
	if catCount == 0 {
		db.Create(&models.Category{NamaCategory: "Elektronik"})
		db.Create(&models.Category{NamaCategory: "Pakaian"})
		log.Println("Initial categories seeded")
	}
}


