package services

import (
	"errors"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
)

// AlamatService defines business logic contract for address management
type AlamatService interface {
	GetAlamatByUser(userID uint) ([]models.Alamat, error)
	GetAlamatByID(id, userID uint) (*models.Alamat, error)
	CreateAlamat(userID uint, req models.CreateAlamatRequest) (uint, error)
	UpdateAlamat(id, userID uint, req models.UpdateAlamatRequest) error
	DeleteAlamat(id, userID uint) error
}

type alamatService struct {
	alamatRepo repositories.AlamatRepository
}

// NewAlamatService creates a new AlamatService instance
func NewAlamatService(alamatRepo repositories.AlamatRepository) AlamatService {
	return &alamatService{alamatRepo: alamatRepo}
}

func (s *alamatService) GetAlamatByUser(userID uint) ([]models.Alamat, error) {
	alamats, err := s.alamatRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return alamats, nil
}

func (s *alamatService) GetAlamatByID(id, userID uint) (*models.Alamat, error) {
	alamat, err := s.alamatRepo.FindByIDAndUserID(id, userID)
	if err != nil || alamat == nil {
		return nil, errors.New("record not found")
	}
	return alamat, nil
}

func (s *alamatService) CreateAlamat(userID uint, req models.CreateAlamatRequest) (uint, error) {
	alamat := models.Alamat{
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
		UserID:       userID,
	}

	if err := s.alamatRepo.Create(&alamat); err != nil {
		return 0, err
	}
	return alamat.ID, nil
}

func (s *alamatService) UpdateAlamat(id, userID uint, req models.UpdateAlamatRequest) error {
	alamat, err := s.alamatRepo.FindByIDAndUserID(id, userID)
	if err != nil || alamat == nil {
		return errors.New("record not found")
	}

	if req.JudulAlamat != "" {
		alamat.JudulAlamat = req.JudulAlamat
	}
	if req.NamaPenerima != "" {
		alamat.NamaPenerima = req.NamaPenerima
	}
	if req.NoTelp != "" {
		alamat.NoTelp = req.NoTelp
	}
	if req.DetailAlamat != "" {
		alamat.DetailAlamat = req.DetailAlamat
	}

	return s.alamatRepo.Update(alamat)
}

func (s *alamatService) DeleteAlamat(id, userID uint) error {
	return s.alamatRepo.Delete(id, userID)
}
