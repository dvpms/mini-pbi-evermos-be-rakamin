package services

import (
	"errors"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
)

// CategoryService defines business logic contract for category operations
type CategoryService interface {
	GetAllCategories() ([]models.CategorySimpleResponse, error)
	GetCategoryByID(id uint) (*models.CategorySimpleResponse, error)
	CreateCategory(req models.CategoryRequest) (uint, error)
	UpdateCategory(id uint, req models.CategoryRequest) error
	DeleteCategory(id uint) error
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

// NewCategoryService creates a new CategoryService instance
func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) GetAllCategories() ([]models.CategorySimpleResponse, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var results []models.CategorySimpleResponse
	for _, cat := range categories {
		results = append(results, models.CategorySimpleResponse{
			ID:           cat.ID,
			NamaCategory: cat.NamaCategory,
		})
	}

	if results == nil {
		results = []models.CategorySimpleResponse{}
	}

	return results, nil
}

func (s *categoryService) GetCategoryByID(id uint) (*models.CategorySimpleResponse, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Kategori tidak ditemukan")
	}

	return &models.CategorySimpleResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (s *categoryService) CreateCategory(req models.CategoryRequest) (uint, error) {
	category := models.Category{
		NamaCategory: req.NamaCategory,
	}

	if err := s.categoryRepo.Create(&category); err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (s *categoryService) UpdateCategory(id uint, req models.CategoryRequest) error {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return errors.New("Kategori tidak ditemukan")
	}

	category.NamaCategory = req.NamaCategory
	return s.categoryRepo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	_, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return errors.New("Kategori tidak ditemukan")
	}

	return s.categoryRepo.Delete(id)
}
