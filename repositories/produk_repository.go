package repositories

import (
	"strings"

	"mini-project-pbi/models"

	"gorm.io/gorm"
)

// ProdukFilter defines parameters for filtering and paginating product listings
type ProdukFilter struct {
	NamaProduk string
	CategoryID uint
	TokoID     uint
	MinHarga   float64
	MaxHarga   float64
	Page       int
	Limit      int
}

// ProdukRepository defines database operations for product entities
type ProdukRepository interface {
	FindAll(filter ProdukFilter) ([]models.Produk, int64, error)
	FindByID(id uint) (*models.Produk, error)
	Create(produk *models.Produk) error
	Update(produk *models.Produk) error
	Delete(id uint) error
	CreatePhotos(photos []models.FotoProduk) error
	DeletePhotosByProdukID(produkID uint) error
}

type produkRepository struct {
	db *gorm.DB
}

// NewProdukRepository creates a new ProdukRepository instance
func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &produkRepository{db: db}
}

func (r *produkRepository) FindAll(filter ProdukFilter) ([]models.Produk, int64, error) {
	var produks []models.Produk
	var total int64

	query := r.db.Model(&models.Produk{})

	if filter.NamaProduk != "" {
		query = query.Where("LOWER(nama_produk) LIKE ?", "%"+strings.ToLower(filter.NamaProduk)+"%")
	}
	if filter.CategoryID != 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.TokoID != 0 {
		query = query.Where("toko_id = ?", filter.TokoID)
	}
	if filter.MinHarga > 0 {
		query = query.Where("harga_konsumen >= ?", filter.MinHarga)
	}
	if filter.MaxHarga > 0 {
		query = query.Where("harga_konsumen <= ?", filter.MaxHarga)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := filter.Page
	limit := filter.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	if err := query.Preload("Toko").Preload("Category").Preload("Photos").
		Offset(offset).Limit(limit).Find(&produks).Error; err != nil {
		return nil, 0, err
	}

	return produks, total, nil
}

func (r *produkRepository) FindByID(id uint) (*models.Produk, error) {
	var produk models.Produk
	if err := r.db.Preload("Toko").Preload("Category").Preload("Photos").
		First(&produk, id).Error; err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) Create(produk *models.Produk) error {
	return r.db.Create(produk).Error
}

func (r *produkRepository) Update(produk *models.Produk) error {
	return r.db.Save(produk).Error
}

func (r *produkRepository) Delete(id uint) error {
	// First delete related photos
	if err := r.DeletePhotosByProdukID(id); err != nil {
		return err
	}
	return r.db.Delete(&models.Produk{}, id).Error
}

func (r *produkRepository) CreatePhotos(photos []models.FotoProduk) error {
	if len(photos) == 0 {
		return nil
	}
	return r.db.Create(&photos).Error
}

func (r *produkRepository) DeletePhotosByProdukID(produkID uint) error {
	return r.db.Where("product_id = ?", produkID).Delete(&models.FotoProduk{}).Error
}
