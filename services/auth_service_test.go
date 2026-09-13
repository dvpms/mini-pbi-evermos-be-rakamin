package services_test

import (
	"fmt"
	"mini-project-pbi/database"
	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
	"mini-project-pbi/services"
	"testing"
	"time"
)

func TestAuthServiceRegisterAndLogin(t *testing.T) {
	db := database.ConnectDatabase()
	userRepo := repositories.NewUserRepository(db)
	tokoRepo := repositories.NewTokoRepository(db)
	authService := services.NewAuthService(userRepo, tokoRepo)

	uniqueSuffix := fmt.Sprintf("%d", time.Now().UnixNano())
	testEmail := fmt.Sprintf("test_%s@example.com", uniqueSuffix)
	testPhone := fmt.Sprintf("08%s", uniqueSuffix[len(uniqueSuffix)-10:])
	testPassword := "password123"

	// 1. Test Register
	regReq := models.RegisterRequest{
		Nama:         "Test User " + uniqueSuffix,
		KataSandi:    testPassword,
		NoTelp:       testPhone,
		TanggalLahir: "02/01/2000",
		Pekerjaan:    "Software Engineer",
		Email:        testEmail,
		IDProvinsi:   "11",
		IDKota:       "1101",
	}

	err := authService.Register(regReq)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Verify User exists in database
	user, err := userRepo.FindByEmail(testEmail)
	if err != nil || user == nil {
		t.Fatalf("User was not found in DB: %v", err)
	}

	// Verify Store was auto-created (Ketentuan #2, #6)
	toko, err := tokoRepo.FindByUserID(user.ID)
	if err != nil || toko == nil {
		t.Fatalf("Store was not auto-created for user %d: %v", user.ID, err)
	}
	if toko.NamaToko != user.NamaUser {
		t.Errorf("Expected store name '%s', got '%s'", user.NamaUser, toko.NamaToko)
	}

	// 2. Test Duplicate Email Registration (Ketentuan #3)
	errDupEmail := authService.Register(regReq)
	if errDupEmail == nil {
		t.Fatal("Expected error on duplicate email registration, got nil")
	}

	// 3. Test Login Success
	loginReq := models.LoginRequest{
		NoTelp:    testPhone,
		KataSandi: testPassword,
	}

	loginResp, err := authService.Login(loginReq)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}
	if loginResp.Token == "" {
		t.Fatal("Expected token in login response, got empty")
	}

	// 4. Test Login Wrong Password
	loginWrongReq := models.LoginRequest{
		NoTelp:    testPhone,
		KataSandi: "wrong_password",
	}
	_, errWrong := authService.Login(loginWrongReq)
	if errWrong == nil {
		t.Fatal("Expected error on wrong password, got nil")
	}
}
