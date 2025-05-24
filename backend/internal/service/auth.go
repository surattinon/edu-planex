package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/surattinon/edu-planex/backend/internal/model"
	"github.com/surattinon/edu-planex/backend/internal/repository"
)

// AuthService handles signup and login logic
type AuthService struct {
	userRepo  repository.UserRepo
	jwtSecret string
}

// NewAuthService creates a new AuthService
func NewAuthService(userRepo repository.UserRepo, jwtSecret string) *AuthService {
	return &AuthService{userRepo, jwtSecret}
}

// Signup registers services
func (a *AuthService) AdvisorSignup(ctx context.Context, advisor_id string, password string, fname, lname string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	advisor := &model.Advisor{FName: fname, LName: lname, AdvisorID: advisor_id, Password: string(hashed)}
	return a.userRepo.CreateAdvisor(ctx, advisor)
}

func (s *AuthService) StudentSignup(ctx context.Context, fname, lname string, student_id string, password string, advisor_id string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student := &model.Student{FName: fname, LName: lname, StudentID: student_id, Password: string(hashed), AdvisorID: advisor_id}
	return s.userRepo.CreateStudent(ctx, student)
}

func (s *AuthService) Login(ctx context.Context, id string, password string) (string, error) {
	// Try advisor by ID
	if adv, _ := s.userRepo.FindByAdvisorID(ctx, id); adv != nil {
		if err := bcrypt.CompareHashAndPassword([]byte(adv.Password), []byte(password)); err != nil {
			return s.makeToken(adv.AdvisorID, "advisor")
		}
	}
	// Try student by ID
	if st, _ := s.userRepo.FindByStudentID(ctx, id); st != nil {
		if err := bcrypt.CompareHashAndPassword([]byte(st.Password), []byte(password)); err != nil {
			return s.makeToken(st.StudentID, "student")
		}
	}
	return "", errors.New("invalid credentials")
}

func (s *AuthService) makeToken(sub string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"exp":  time.Now().Add(72 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}
