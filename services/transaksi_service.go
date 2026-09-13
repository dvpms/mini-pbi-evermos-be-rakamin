package services

import (
	"errors"
	"fmt"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
)

// TransaksiService defines business logic contract for transaction handling
type TransaksiService interface {
	CreateTrx(userID uint, req models.CreateTransaksiRequest) (uint, error)
	GetAllTrx(userID uint) ([]models.Transaksi, error)
	GetTrxByID(id, userID uint) (*models.Transaksi, error)
}

type transaksiService struct {
	transaksiRepo repositories.TransaksiRepository
	alamatRepo    repositories.AlamatRepository
	produkRepo    repositories.ProdukRepository
}

// NewTransaksiService creates a new TransaksiService instance
func NewTransaksiService(
	transaksiRepo repositories.TransaksiRepository,
	alamatRepo repositories.AlamatRepository,
	produkRepo repositories.ProdukRepository,
) TransaksiService {
	return &transaksiService{
		transaksiRepo: transaksiRepo,
		alamatRepo:    alamatRepo,
		produkRepo:    produkRepo,
	}
}

func (s *transaksiService) CreateTrx(userID uint, req models.CreateTransaksiRequest) (uint, error) {
	// Guard clause: items cannot be empty
	if len(req.DetailTrx) == 0 {
		return 0, errors.New("Detail transaksi tidak boleh kosong")
	}

	// Validate delivery address belongs to the authenticated user
	alamat, err := s.alamatRepo.FindByIDAndUserID(req.AlamatKirim, userID)
	if err != nil || alamat == nil {
		return 0, errors.New("Alamat pengiriman tidak valid atau bukan milik Anda")
	}

	transaksi := models.Transaksi{
		UserID:           userID,
		AlamatPengiriman: req.AlamatKirim,
		MethodBayar:      req.MethodBayar,
	}

	createdTrx, err := s.transaksiRepo.CreateTransactionWithLog(&transaksi, req.DetailTrx)
	if err != nil {
		return 0, err
	}

	return createdTrx.ID, nil
}

func (s *transaksiService) GetAllTrx(userID uint) ([]models.Transaksi, error) {
	transaksis, err := s.transaksiRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	if transaksis == nil {
		transaksis = []models.Transaksi{}
	}
	return transaksis, nil
}

func (s *transaksiService) GetTrxByID(id, userID uint) (*models.Transaksi, error) {
	trx, err := s.transaksiRepo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, fmt.Errorf("No Data Trx")
	}
	return trx, nil
}
