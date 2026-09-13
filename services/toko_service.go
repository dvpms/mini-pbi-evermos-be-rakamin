package services

import (
	"errors"
	"mime/multipart"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
	"mini-project-pbi/utils"
)

// TokoService defines business logic for store operations
type TokoService interface {
	GetMyToko(userID uint) (*models.Toko, error)
	GetTokoByID(id uint) (*models.Toko, error)
	GetAllToko(page, limit int, nama string) ([]models.TokoSimpleResponse, int64, error)
	UpdateToko(tokoID, userID uint, namaToko string, photoFile *multipart.FileHeader) error
}

type tokoService struct {
	tokoRepo repositories.TokoRepository
}

// NewTokoService creates a new TokoService instance
func NewTokoService(tokoRepo repositories.TokoRepository) TokoService {
	return &tokoService{tokoRepo: tokoRepo}
}

func (s *tokoService) GetMyToko(userID uint) (*models.Toko, error) {
	toko, err := s.tokoRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("Toko tidak ditemukan")
	}
	return toko, nil
}

func (s *tokoService) GetTokoByID(id uint) (*models.Toko, error) {
	toko, err := s.tokoRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Toko tidak ditemukan")
	}
	return toko, nil
}

func (s *tokoService) GetAllToko(page, limit int, nama string) ([]models.TokoSimpleResponse, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	tokos, total, err := s.tokoRepo.FindAll(page, limit, nama)
	if err != nil {
		return nil, 0, err
	}

	var results []models.TokoSimpleResponse
	for _, t := range tokos {
		results = append(results, models.TokoSimpleResponse{
			ID:       t.ID,
			NamaToko: t.NamaToko,
			UrlFoto:  t.UrlFoto,
		})
	}

	if results == nil {
		results = []models.TokoSimpleResponse{}
	}

	return results, total, nil
}

func (s *tokoService) UpdateToko(tokoID, userID uint, namaToko string, photoFile *multipart.FileHeader) error {
	toko, err := s.tokoRepo.FindByID(tokoID)
	if err != nil {
		return errors.New("Toko tidak ditemukan")
	}

	// User isolation rule (Ketentuan 13)
	if toko.UserID != userID {
		return errors.New("Anda tidak memiliki akses untuk mengubah toko ini")
	}

	if namaToko != "" {
		toko.NamaToko = namaToko
	}

	if photoFile != nil {
		filename, err := utils.SaveUploadedFile(photoFile, utils.UploadDir)
		if err != nil {
			return err
		}
		toko.UrlFoto = filename
	}

	return s.tokoRepo.Update(toko)
}
