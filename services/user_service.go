package services

import (
	"errors"
	"time"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
	"mini-project-pbi/utils"
)

// UserService defines business logic contract for user profile management
type UserService interface {
	GetProfile(userID uint) (*models.User, error)
	UpdateProfile(userID uint, req models.UpdateProfileRequest) error
}

type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetProfile(userID uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("User tidak ditemukan")
	}
	return user, nil
}

func (s *userService) UpdateProfile(userID uint, req models.UpdateProfileRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil || user == nil {
		return errors.New("User tidak ditemukan")
	}

	if req.Nama != "" {
		user.NamaUser = req.Nama
	}
	if req.NoTelp != "" && req.NoTelp != user.NoTelp {
		existing, _ := s.userRepo.FindByNoTelp(req.NoTelp)
		if existing != nil && existing.ID != user.ID {
			return errors.New("Nomor telepon sudah digunakan")
		}
		user.NoTelp = req.NoTelp
	}
	if req.Email != "" && req.Email != user.Email {
		existing, _ := s.userRepo.FindByEmail(req.Email)
		if existing != nil && existing.ID != user.ID {
			return errors.New("Email sudah digunakan")
		}
		user.Email = req.Email
	}
	if req.KataSandi != "" {
		hash, err := utils.HashPassword(req.KataSandi)
		if err == nil {
			user.KataSandi = hash
		}
	}
	if req.JenisKelamin != "" {
		user.JenisKelamin = req.JenisKelamin
	}
	if req.Pekerjaan != "" {
		user.Pekerjaan = req.Pekerjaan
	}
	if req.IDProvinsi != "" {
		user.IDProvinsi = req.IDProvinsi
	}
	if req.IDKota != "" {
		user.IDKota = req.IDKota
	}
	if req.TanggalLahir != "" {
		if parsed, err := time.Parse("02/01/2006", req.TanggalLahir); err == nil {
			user.TanggalLahir = parsed
		} else if parsed, err := time.Parse("2006-01-02", req.TanggalLahir); err == nil {
			user.TanggalLahir = parsed
		}
	}

	return s.userRepo.Update(user)
}
