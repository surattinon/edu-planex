package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/surattinon/edu-planex/backend/internal/model"
)

// StudentRepo defines methods for user CRUD
type UserRepo interface {
	CreateStudent(ctx context.Context, u *model.Student) error
	FindByStudentID(ctx context.Context, stuid string) (*model.Student, error)
	CreateAdvisor(ctx context.Context, u *model.Advisor) error
	FindByAdvisorID(ctx context.Context, advid string) (*model.Advisor, error)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo constructs a UserRepo backed by GORM
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}


func (r *userRepo) CreateStudent(ctx context.Context, u *model.Student) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *userRepo) FindByStudentID(ctx context.Context, stuid string) (*model.Student, error) {
	var student model.Student
	if err := r.db.WithContext(ctx).Table("students").Where("student_id = ?", stuid).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *userRepo) CreateAdvisor(ctx context.Context, u *model.Advisor) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *userRepo) FindByAdvisorID(ctx context.Context, advid string) (*model.Advisor, error) {
	var advisor model.Advisor
	if err := r.db.WithContext(ctx).Table("advisors").Where("advisor_id = ?", advid).First(&advisor).Error; err != nil {
		return nil, err
	}
	return &advisor, nil
}
