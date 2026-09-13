package services

import (
	"errors"
	"mime/multipart"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
	"mini-project-pbi/utils"
)

// ProdukService defines business logic contract for product operations
type ProdukService interface {
	GetAllProduk(filter repositories.ProdukFilter) ([]models.Produk, int64, error)
	GetProdukByID(id uint) (*models.Produk, error)
	CreateProduk(userID uint, req models.CreateProdukRequest, photoHeaders []*multipart.FileHeader) (uint, error)
	UpdateProduk(id, userID uint, req models.UpdateProdukRequest, photoHeaders []*multipart.FileHeader) error
	DeleteProduk(id, userID uint) error
}

type produkService struct {
	produkRepo repositories.ProdukRepository
	tokoRepo   repositories.TokoRepository
}

// NewProdukService creates a new ProdukService instance
func NewProdukService(produkRepo repositories.ProdukRepository, tokoRepo repositories.TokoRepository) ProdukService {
	return &produkService{
		produkRepo: produkRepo,
		tokoRepo:   tokoRepo,
	}
}

func (s *produkService) GetAllProduk(filter repositories.ProdukFilter) ([]models.Produk, int64, error) {
	return s.produkRepo.FindAll(filter)
}

func (s *produkService) GetProdukByID(id uint) (*models.Produk, error) {
	produk, err := s.produkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("No Data Product")
	}
	return produk, nil
}

func (s *produkService) CreateProduk(userID uint, req models.CreateProdukRequest, photoHeaders []*multipart.FileHeader) (uint, error) {
	toko, err := s.tokoRepo.FindByUserID(userID)
	if err != nil {
		return 0, errors.New("Toko tidak ditemukan. Silakan buat toko terlebih dahulu")
	}

	slug := utils.GenerateSlug(req.NamaProduk)

	produk := models.Produk{
		NamaProduk:    req.NamaProduk,
		Slug:          slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		TokoID:        toko.ID,
		CategoryID:    req.CategoryID,
	}

	if err := s.produkRepo.Create(&produk); err != nil {
		return 0, err
	}

	var photos []models.FotoProduk
	for _, header := range photoHeaders {
		if header == nil {
			continue
		}
		filename, err := utils.SaveUploadedFile(header, utils.UploadDir)
		if err != nil {
			return 0, err
		}
		photos = append(photos, models.FotoProduk{
			ProductID: produk.ID,
			URL:       filename,
		})
	}

	if len(photos) > 0 {
		if err := s.produkRepo.CreatePhotos(photos); err != nil {
			return 0, err
		}
	}

	return produk.ID, nil
}

func (s *produkService) UpdateProduk(id, userID uint, req models.UpdateProdukRequest, photoHeaders []*multipart.FileHeader) error {
	produk, err := s.produkRepo.FindByID(id)
	if err != nil {
		return errors.New("No Data Product")
	}

	toko, err := s.tokoRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("Toko tidak ditemukan")
	}

	// User isolation rule (Ketentuan #14)
	if produk.TokoID != toko.ID {
		return errors.New("Anda tidak memiliki akses untuk mengubah produk ini")
	}

	if req.NamaProduk != "" {
		produk.NamaProduk = req.NamaProduk
		produk.Slug = utils.GenerateSlug(req.NamaProduk)
	}
	if req.CategoryID != 0 {
		produk.CategoryID = req.CategoryID
	}
	if req.HargaReseller > 0 {
		produk.HargaReseller = req.HargaReseller
	}
	if req.HargaKonsumen > 0 {
		produk.HargaKonsumen = req.HargaKonsumen
	}
	if req.Stok > 0 {
		produk.Stok = req.Stok
	}
	if req.Deskripsi != "" {
		produk.Deskripsi = req.Deskripsi
	}

	if err := s.produkRepo.Update(produk); err != nil {
		return err
	}

	if len(photoHeaders) > 0 {
		_ = s.produkRepo.DeletePhotosByProdukID(id)

		var photos []models.FotoProduk
		for _, header := range photoHeaders {
			if header == nil {
				continue
			}
			filename, err := utils.SaveUploadedFile(header, utils.UploadDir)
			if err != nil {
				return err
			}
			photos = append(photos, models.FotoProduk{
				ProductID: produk.ID,
				URL:       filename,
			})
		}
		if len(photos) > 0 {
			if err := s.produkRepo.CreatePhotos(photos); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *produkService) DeleteProduk(id, userID uint) error {
	produk, err := s.produkRepo.FindByID(id)
	if err != nil {
		return errors.New("No Data Product")
	}

	toko, err := s.tokoRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("Toko tidak ditemukan")
	}

	// User isolation rule (Ketentuan #14)
	if produk.TokoID != toko.ID {
		return errors.New("Anda tidak memiliki akses untuk menghapus produk ini")
	}

	return s.produkRepo.Delete(id)
}
