package usecase

import (
	"unicode"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/email"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/password"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
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
	userRepo       repositories.UserRepository
	passwordHasher password.PasswordHasher
	emailSender    email.EmailSender
}

func NewUserUsecase(userRepo repositories.UserRepository, passwordHasher password.PasswordHasher) UserUsecase {
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
		return nil, ErrEmailAlreadyExists
	}

	existingUser, err = u.userRepo.FindByNickname(input.Nickname)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrNicknameAlreadyExists
	}

	if err := validatePassword(input.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := u.passwordHasher.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
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
		return ErrPasswordTooShort
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
		return ErrPasswordNoUppercase
	}
	if !hasLowercase {
		return ErrPasswordNoLowercase
	}
	if !hasNumber {
		return ErrPasswordNoNumber
	}
	if !hasSpecialChar {
		return ErrPasswordNoSpecialChar
	}

	return nil
}
