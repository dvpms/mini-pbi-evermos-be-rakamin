package repositories

import (
	"strings"

	"mini-project-pbi/models"

	"gorm.io/gorm"
)

// TokoRepository defines contract for store persistence operations
type TokoRepository interface {
	Create(toko *models.Toko) error
	FindByUserID(userID uint) (*models.Toko, error)
	FindByID(id uint) (*models.Toko, error)
	Update(toko *models.Toko) error
	FindAll(page, limit int, search string) ([]models.Toko, int64, error)
}

type tokoRepository struct {
	db *gorm.DB
}

// NewTokoRepository creates a new TokoRepository instance
func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoRepository{db: db}
}

func (r *tokoRepository) Create(toko *models.Toko) error {
	return r.db.Create(toko).Error
}

func (r *tokoRepository) FindByUserID(userID uint) (*models.Toko, error) {
	var toko models.Toko
	if err := r.db.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoRepository) FindByID(id uint) (*models.Toko, error) {
	var toko models.Toko
	if err := r.db.First(&toko, id).Error; err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoRepository) Update(toko *models.Toko) error {
	return r.db.Save(toko).Error
}

func (r *tokoRepository) FindAll(page, limit int, search string) ([]models.Toko, int64, error) {
	var tokos []models.Toko
	var total int64

	query := r.db.Model(&models.Toko{})
	if search != "" {
		query = query.Where("LOWER(nama_toko) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&tokos).Error; err != nil {
		return nil, 0, err
	}

	return tokos, total, nil
}
