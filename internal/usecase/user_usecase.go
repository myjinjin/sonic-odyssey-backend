package usecase

import (
	"errors"
	"fmt"
	"time"
	"unicode"

	"github.com/google/uuid"
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
	SignUp(SignUpInput) (*SignUpOutput, error)
	SendPasswordRecoveryEmail(baseURL, email string) error
	ResetPassword(password, flowID string) error
}

type userUsecase struct {
	userRepo          repositories.UserRepository
	passwordResetRepo repositories.PasswordResetFlowRepository

	emailEncryptor encryption.Encryptor
	emailSender    email.EmailSender
}

func NewUserUsecase(userRepo repositories.UserRepository, passwordResetRepo repositories.PasswordResetFlowRepository, emailEncryptor encryption.Encryptor, emailSender email.EmailSender) UserUsecase {
	return &userUsecase{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
		emailEncryptor:    emailEncryptor,
		emailSender:       emailSender,
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
		err := u.emailSender.SendEmail(input.Email, email.TemplateWelcome, welcomeData)
		if err != nil {
			logging.Log().Error("failed to send welcome email",
				zap.Error(err),
				zap.String("email", input.Email),
				zap.String("name", input.Name),
			)

			retries := 3
			for i := 0; i < retries; i++ {
				time.Sleep(time.Second * time.Duration(i+1))
				err = u.emailSender.SendEmail(input.Email, email.TemplateWelcome, welcomeData)
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

func (u *userUsecase) SendPasswordRecoveryEmail(baseURL, userEmail string) error {
	user, err := u.userRepo.FindByEmailHash(hash.SHA256EmailHasher().HashEmail(userEmail))
	if err != nil {
		if errors.Is(err, repositories.ErrFind) {
			return ErrFindingRecord
		}
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrUserNotFound
		}
	}

	passwordResetFlow := &entities.PasswordResetFlow{
		UserID:    user.ID,
		FlowID:    generateFlowID(),
		ExpiresAt: func() *time.Time { t := time.Now().Add(time.Hour * 2); return &t }(),
	}
	if err := u.passwordResetRepo.Create(passwordResetFlow); err != nil {
		return ErrCreatingRecord
	}

	resetLink := fmt.Sprintf("%s/password/recovery?flow_id=%s", baseURL, passwordResetFlow.FlowID)

	go func() {
		passwordResetData := email.PasswordResetData{Name: user.Name, ResetLink: resetLink}
		err := u.emailSender.SendEmail(userEmail, email.TemplatePasswordReset, passwordResetData)
		if err != nil {
			logging.Log().Error("failed to send password reset email",
				zap.Error(err),
				zap.String("email", userEmail),
				zap.String("name", user.Name),
			)

			retries := 3
			for i := 0; i < retries; i++ {
				time.Sleep(time.Second * time.Duration(i+1))
				err = u.emailSender.SendEmail(userEmail, email.TemplatePasswordReset, passwordResetData)
				if err == nil {
					break
				}
				logging.Log().Error("failed to send password reset email, retrying...",
					zap.Error(err),
					zap.String("email", userEmail),
					zap.String("name", user.Name),
					zap.Int("retry", i+1),
				)
			}
		}
	}()

	return nil
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

func (u *userUsecase) ResetPassword(password, flowID string) error {
	flow, err := u.passwordResetRepo.FindByFlowID(flowID)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrPasswordResetFlowNotFound
		}
		if errors.Is(err, repositories.ErrFind) {
			return ErrFindingRecord
		}
	}
	if flow.ExpiresAt.Before(time.Now()) {
		return ErrPasswordResetFlowExpired
	}

	if err := validatePassword(password); err != nil {
		return err
	}

	user := &flow.User
	hashedPassword, err := hash.BCryptPasswordHasher().HashPassword(password)
	if err != nil {
		return ErrHashingPassword
	}

	user.PasswordHash = hashedPassword
	if err := u.userRepo.Update(user); err != nil {
		return ErrUpatingRecord
	}

	if err := u.passwordResetRepo.DeleteByFlowID(flowID); err != nil {
		return ErrDeletingRecord
	}

	return nil
}

func generateFlowID() string {
	currentTime := time.Now().Unix()
	uuidValue := uuid.New()
	return fmt.Sprintf("%s:%d", uuidValue.String(), currentTime)
}
