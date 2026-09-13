package repositories

import (
	"errors"
	"fmt"
	"time"

	"mini-project-pbi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TransaksiRepository defines database interaction contract for transactions
type TransaksiRepository interface {
	CreateTransactionWithLog(transaksi *models.Transaksi, items []models.TransactionItemInput) (*models.Transaksi, error)
	FindByUserID(userID uint) ([]models.Transaksi, error)
	FindByIDAndUserID(id, userID uint) (*models.Transaksi, error)
}

type transaksiRepository struct {
	db *gorm.DB
}

// NewTransaksiRepository creates a new TransaksiRepository instance
func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepository{db: db}
}

// CreateTransactionWithLog executes atomic transaction to deduct stock, snapshot log_produk, and persist transaction & details
func (r *transaksiRepository) CreateTransactionWithLog(transaksi *models.Transaksi, items []models.TransactionItemInput) (*models.Transaksi, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var totalHarga float64
	var details []models.DetailTransaksi

	for _, item := range items {
		var produk models.Produk
		// Lock the product row for update to avoid race conditions
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Photos").First(&produk, item.ProductID).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("Produk dengan ID %d tidak ditemukan", item.ProductID)
			}
			return nil, err
		}

		// Validate stock availability
		if produk.Stok < item.Kuantitas {
			tx.Rollback()
			return nil, fmt.Errorf("Stok produk '%s' tidak mencukupi (tersedia: %d, diminta: %d)", produk.NamaProduk, produk.Stok, item.Kuantitas)
		}

		// Deduct stock and save
		produk.Stok -= item.Kuantitas
		if err := tx.Save(&produk).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Create snapshot in log_produks (Ketentuan 16 & 17)
		logProduk := models.LogProduk{
			ProdukID:      produk.ID,
			NamaProduk:    produk.NamaProduk,
			Slug:          produk.Slug,
			HargaReseller: produk.HargaReseller,
			HargaKonsumen: produk.HargaKonsumen,
			Deskripsi:     produk.Deskripsi,
			TokoID:        produk.TokoID,
			CategoryID:    produk.CategoryID,
		}
		if err := tx.Create(&logProduk).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		subtotal := produk.HargaKonsumen * float64(item.Kuantitas)
		totalHarga += subtotal

		details = append(details, models.DetailTransaksi{
			LogProdukID: logProduk.ID,
			Kuantitas:   item.Kuantitas,
			HargaTotal:  subtotal,
			TokoID:      produk.TokoID,
		})
	}

	// Generate unique invoice code
	transaksi.KodeInvoice = fmt.Sprintf("INV-%d-%d", transaksi.UserID, time.Now().UnixNano()/1000000)
	transaksi.HargaTotal = totalHarga

	// Save transaction header
	if err := tx.Create(transaksi).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save transaction details
	for i := range details {
		details[i].TransaksiID = transaksi.ID
		if err := tx.Create(&details[i]).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Fetch complete populated transaction for return
	var createdTrx models.Transaksi
	if err := r.db.
		Preload("AlamatKirim").
		Preload("DetailTrx.Product.Category").
		Preload("DetailTrx.Product.Toko").
		Preload("DetailTrx.Toko").
		First(&createdTrx, transaksi.ID).Error; err != nil {
		return transaksi, nil
	}

	return &createdTrx, nil
}

// FindByUserID retrieves all transactions for a given user with isolated scope (Ketentuan 15)
func (r *transaksiRepository) FindByUserID(userID uint) ([]models.Transaksi, error) {
	var transaksis []models.Transaksi
	err := r.db.
		Where("user_id = ?", userID).
		Preload("AlamatKirim").
		Preload("DetailTrx.Product.Category").
		Preload("DetailTrx.Product.Toko").
		Preload("DetailTrx.Toko").
		Order("created_at desc").
		Find(&transaksis).Error
	if err != nil {
		return nil, err
	}
	return transaksis, nil
}

// FindByIDAndUserID retrieves single transaction owned by the specified user (Ketentuan 15)
func (r *transaksiRepository) FindByIDAndUserID(id, userID uint) (*models.Transaksi, error) {
	var trx models.Transaksi
	err := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		Preload("AlamatKirim").
		Preload("DetailTrx.Product.Category").
		Preload("DetailTrx.Product.Toko").
		Preload("DetailTrx.Toko").
		First(&trx).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}
