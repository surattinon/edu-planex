package service

import (
	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (s *CategoryService) ListAllCat() ([]model.CurriculumCategory, error) {
	var cats []model.CurriculumCategory
	if err := s.db.Preload("Courses").Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (s *CategoryService) FindCatByID(id uint) (*model.CurriculumCategory, error) {
	var cat model.CurriculumCategory
	if err := s.db.Preload("Courses").First(&cat, "cat_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}
