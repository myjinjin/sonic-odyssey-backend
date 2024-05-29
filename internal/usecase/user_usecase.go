package usecase

import (
	"errors"
	"time"
	"unicode"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/email"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/encryption"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/logging"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"go.uber.org/zap"
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
	emailEncryptor encryption.Encryptor
	emailSender    email.EmailSender
}

func NewUserUsecase(userRepo repositories.UserRepository, emailEncryptor encryption.Encryptor, emailSender email.EmailSender) UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		emailEncryptor: emailEncryptor,
		emailSender:    emailSender,
	}
}

func (u *userUsecase) SignUp(input SignUpInput) (*SignUpOutput, error) {
	hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
	existingUser, err := u.userRepo.FindByEmailHash(hashedEmail)
	if err != nil && errors.Is(err, repositories.ErrFind) {
		return nil, ErrFindingRecord
	}
	if existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	existingUser, err = u.userRepo.FindByNickname(input.Nickname)
	if err != nil && errors.Is(err, repositories.ErrFind) {
		return nil, ErrFindingRecord
	}
	if existingUser != nil {
		return nil, ErrNicknameAlreadyExists
	}

	if err := validatePassword(input.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := hash.BCryptPasswordHasher().HashPassword(input.Password)
	if err != nil {
		return nil, ErrHashingPassword
	}

	encryptedEmail, err := u.emailEncryptor.Encrypt(input.Email)
	if err != nil {
		return nil, ErrEncryptingEmail
	}

	user := &entities.User{
		Email:        encryptedEmail,
		EmailHash:    hashedEmail,
		PasswordHash: hashedPassword,
		Name:         input.Name,
		Nickname:     input.Nickname,
		UserProfile:  &entities.UserProfile{},
	}
	if err := u.userRepo.Create(user); err != nil {
		return nil, ErrCreatingRecord
	}

	go func() {
		welcomeData := email.WelcomeData{Name: input.Name}
		err := u.emailSender.SendEmail(input.Email, "Welcome to the sonic odyssey~!", "welcome.html", welcomeData)
		if err != nil {
			logging.Log().Error("failed to send welcome email",
				zap.Error(err),
				zap.String("email", input.Email),
				zap.String("name", input.Name),
			)

			retries := 3
			for i := 0; i < retries; i++ {
				time.Sleep(time.Second * time.Duration(i+1))
				err = u.emailSender.SendEmail(input.Email, "Welcome to the sonic odyssey~!", "welcome.html", welcomeData)
				if err == nil {
					break
				}
				logging.Log().Error("failed to send welcome email, retrying...",
					zap.Error(err),
					zap.String("email", input.Email),
					zap.String("name", input.Name),
					zap.Int("retry", i+1),
				)
			}
		}
	}()

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
