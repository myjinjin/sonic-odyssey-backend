package usecase

import (
	"errors"
	"unicode"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/email"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/password"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain"
	"github.com/myjinjin/sonic-odyssey-backend/internal/repository"
)

type SignUpInput struct {
	Email    string
	Password string
	Name     string
	Nickname string
}

type SignUpOutput struct {
	UserID uint
}

type UserUsecase interface {
	SignUp(input SignUpInput) (*SignUpOutput, error)
}

type userUsecase struct {
	userRepo       repository.UserRepository
	passwordHasher password.PasswordHasher
	emailSender    email.EmailSender
}

func NewUserUsecase(userRepo repository.UserRepository, passwordHasher password.PasswordHasher) UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (u *userUsecase) SignUp(input SignUpInput) (*SignUpOutput, error) {
	existingUser, err := u.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	existingUser, err = u.userRepo.FindByNickname(input.Nickname)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("nickname already exists")
	}

	if err := validatePassword(input.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := u.passwordHasher.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Name:         input.Name,
		Nickname:     input.Nickname,
	}
	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	welcomeData := email.WelcomeData{Name: input.Name}
	if err := u.emailSender.SendEmail(input.Email, "welcome", welcomeData); err != nil {
		return nil, err
	}

	output := &SignUpOutput{
		UserID: user.ID,
	}

	return output, nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasUppercase := false
	hasLowercase := false
	hasNumber := false
	hasSpecialChar := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasUppercase {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLowercase {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecialChar {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
