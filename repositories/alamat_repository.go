package repositories

import (
	"mini-project-pbi/models"

	"gorm.io/gorm"
)

// AlamatRepository defines contract for address persistence operations
type AlamatRepository interface {
	FindByUserID(userID uint) ([]models.Alamat, error)
	FindByIDAndUserID(id, userID uint) (*models.Alamat, error)
	Create(alamat *models.Alamat) error
	Update(alamat *models.Alamat) error
	Delete(id, userID uint) error
}

type alamatRepository struct {
	db *gorm.DB
}

// NewAlamatRepository creates a new AlamatRepository instance
func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &alamatRepository{db: db}
}

func (r *alamatRepository) FindByUserID(userID uint) ([]models.Alamat, error) {
	var alamats []models.Alamat
	if err := r.db.Where("user_id = ?", userID).Find(&alamats).Error; err != nil {
		return nil, err
	}
	return alamats, nil
}

func (r *alamatRepository) FindByIDAndUserID(id, userID uint) (*models.Alamat, error) {
	var alamat models.Alamat
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&alamat).Error; err != nil {
		return nil, err
	}
	return &alamat, nil
}

func (r *alamatRepository) Create(alamat *models.Alamat) error {
	return r.db.Create(alamat).Error
}

func (r *alamatRepository) Update(alamat *models.Alamat) error {
	return r.db.Save(alamat).Error
}

func (r *alamatRepository) Delete(id, userID uint) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Alamat{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
