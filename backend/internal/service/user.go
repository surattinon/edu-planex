package service

import (
	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetProfile() (*model.UserProfile, error) {
	var u model.UserProfile
	if err := s.db.First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) UpdateProfile(input *model.UserProfile) (*model.UserProfile, error) {
	// SQL: UPDATE user_profile SET display_name=?, email=?, avatar_url=?, updated_at=? WHERE user_id=?
	if err := s.db.
		Model(&model.UserProfile{}).
		Where("user_id = ?", input.UserID).
		Updates(input).
		Error; err != nil {
		return nil, err
	}

	// Return the fresh copy
	return s.GetProfile()
}
