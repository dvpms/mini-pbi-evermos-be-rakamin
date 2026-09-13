package services

import (
	"errors"
	"time"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
	"mini-project-pbi/utils"
)

// AuthService defines business logic interface for authentication
type AuthService interface {
	Register(req models.RegisterRequest) error
	Login(req models.LoginRequest) (*models.LoginResponseData, error)
}

type authService struct {
	userRepo repositories.UserRepository
	tokoRepo repositories.TokoRepository
}

// NewAuthService creates a new AuthService instance
func NewAuthService(userRepo repositories.UserRepository, tokoRepo repositories.TokoRepository) AuthService {
	return &authService{
		userRepo: userRepo,
		tokoRepo: tokoRepo,
	}
}

func (s *authService) Register(req models.RegisterRequest) error {
	// Check duplicate email (Ketentuan #3)
	existingEmail, _ := s.userRepo.FindByEmail(req.Email)
	if existingEmail != nil {
		return errors.New("Email sudah terdaftar")
	}

	// Check duplicate phone number (Ketentuan #3)
	existingPhone, _ := s.userRepo.FindByNoTelp(req.NoTelp)
	if existingPhone != nil {
		return errors.New("Nomor telepon sudah terdaftar")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.KataSandi)
	if err != nil {
		return errors.New("Gagal memproses kata sandi")
	}

	// Parse date of birth supporting common formats
	var tgl time.Time
	if req.TanggalLahir != "" {
		if parsed, err := time.Parse("02/01/2006", req.TanggalLahir); err == nil {
			tgl = parsed
		} else if parsed, err := time.Parse("2006-01-02", req.TanggalLahir); err == nil {
			tgl = parsed
		}
	}

	user := models.User{
		NamaUser:     req.Nama,
		KataSandi:    hashedPassword,
		NoTelp:       req.NoTelp,
		TanggalLahir: tgl,
		Pekerjaan:    req.Pekerjaan,
		Email:        req.Email,
		IDProvinsi:   req.IDProvinsi,
		IDKota:       req.IDKota,
		Role:         "user",
	}

	if err := s.userRepo.Create(&user); err != nil {
		return err
	}

	// Auto-create store for the user upon registration (Ketentuan #2, #6)
	toko := models.Toko{
		NamaToko: user.NamaUser,
		UrlFoto:  "",
		UserID:   user.ID,
	}
	if err := s.tokoRepo.Create(&toko); err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(req models.LoginRequest) (*models.LoginResponseData, error) {
	user, err := s.userRepo.FindByNoTelp(req.NoTelp)
	if err != nil || user == nil {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	if !utils.CheckPasswordHash(req.KataSandi, user.KataSandi) {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, errors.New("Gagal membuat token autentikasi")
	}

	tglStr := ""
	if !user.TanggalLahir.IsZero() {
		tglStr = user.TanggalLahir.Format("02/01/2006")
	}

	resp := &models.LoginResponseData{
		Nama:         user.NamaUser,
		NoTelp:       user.NoTelp,
		TanggalLahir: tglStr,
		Tentang:      "",
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi:   user.IDProvinsi,
		IDKota:       user.IDKota,
		Token:        token,
	}

	return resp, nil
}
