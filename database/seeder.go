package database

import (
	"log"
	"time"

	"mini-project-pbi/models"
	"mini-project-pbi/utils"

	"gorm.io/gorm"
)

// SeedAll executes all database seeders
func SeedAll(db *gorm.DB) {
	SeedCategories(db)
	SeedAdminUser(db)
}

// SeedCategories seeds initial default product categories if empty
func SeedCategories(db *gorm.DB) {
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count == 0 {
		categories := []models.Category{
			{NamaCategory: "Elektronik"},
			{NamaCategory: "Pakaian"},
			{NamaCategory: "Aksesoris"},
			{NamaCategory: "Rumah Tangga"},
			{NamaCategory: "Kesehatan & Kecantikan"},
		}
		if err := db.Create(&categories).Error; err != nil {
			log.Printf("Gagal melakukan seed kategori: %v\n", err)
		} else {
			log.Println("Default categories seeded successfully")
		}
	}
}

// SeedAdminUser seeds default admin account (Ketentuan #8 & Tugas 6)
func SeedAdminUser(db *gorm.DB) {
	var adminCount int64
	db.Model(&models.User{}).Where("role = ?", "admin").Count(&adminCount)
	if adminCount == 0 {
		hashedPassword, err := utils.HashPassword("admin123")
		if err != nil {
			log.Printf("Gagal hash password admin: %v\n", err)
			return
		}

		adminUser := models.User{
			NamaUser:     "Admin Evermos",
			KataSandi:    hashedPassword,
			NoTelp:       "081199887766",
			TanggalLahir: time.Date(1995, time.January, 1, 0, 0, 0, 0, time.UTC),
			Pekerjaan:    "Administrator",
			Email:        "admin@evermos.com",
			IDProvinsi:   "32",
			IDKota:       "3273",
			Role:         "admin",
		}

		if err := db.Create(&adminUser).Error; err != nil {
			log.Printf("Gagal melakukan seed admin user: %v\n", err)
			return
		}

		// Otomatis buat toko untuk admin (Ketentuan #6)
		adminToko := models.Toko{
			NamaToko: "Evermos Official Store",
			UserID:   adminUser.ID,
		}
		_ = db.Create(&adminToko).Error

		log.Println("Default admin user ('admin@evermos.com' / 'admin123') and official store seeded successfully")
	}
}
