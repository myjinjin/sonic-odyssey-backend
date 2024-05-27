package usecase

import (
	"errors"
	"testing"

	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:generate mockery --name=UserRepository --dir=../domain/repositories --output=./mocks
//go:generate mockery --name=PasswordHasher --dir=../../infrastructure/hash --output=./mocks
//go:generate mockery --name=EmailHasher --dir=../../infrastructure/hash --output=./mocks
//go:generate mockery --name=Encryptor --dir=../../infrastructure/encryption --output=./mocks
//go:generate mockery --name=EmailSender --dir=../../infrastructure/email --output=./mocks

func TestUserUsecase_SignUp_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := "hashed_email"
	encryptedEmail := "encrypted_email"
	hashedPassword := "hashed_password"

	emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	passwordHasher.On("HashPassword", input.Password).Return(hashedPassword, nil)
	emailEncryptor.On("Encrypt", input.Email).Return(encryptedEmail, nil)

	userRepo.On("Create", mock.MatchedBy(func(user *entities.User) bool {
		user.ID = 1
		return user.Email == encryptedEmail &&
			user.EmailHash == hashedEmail &&
			user.PasswordHash == hashedPassword &&
			user.Name == input.Name &&
			user.Nickname == input.Nickname
	})).Return(nil)

	emailSender.On("SendEmail", input.Email, "Welcome to the sonic odyssey~!", "welcome.html", mock.AnythingOfType("email.WelcomeData")).Return(nil)

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.NotZero(t, output.UserID)

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_EmailAlreadyExists(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := "hashed_email"

	emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
	userRepo.On("FindByEmailHash", hashedEmail).Return(&entities.User{
		ID:        1,
		Email:     "test@example.com",
		EmailHash: hashedEmail,
		Name:      "Test User",
		Nickname:  "testuser",
	}, nil)

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrEmailAlreadyExists, err)
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_NicknameAlreadyExists(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := "hashed_email"

	emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(&entities.User{
		ID:        1,
		Email:     "test2@example.com",
		EmailHash: "other_hashed_email",
		Name:      "Test User",
		Nickname:  "testuser",
	}, nil)

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrNicknameAlreadyExists, err)
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_InvalidPassword(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	// Test cases for invalid passwords
	invalidPasswords := []string{
		"short",
		"nouppercasE",
		"NOLOWERCASE",
		"NoNumber",
		"NoSpecialChar",
	}

	for _, password := range invalidPasswords {
		input := SignUpInput{
			Email:    "test@example.com",
			Password: password,
			Name:     "Test User",
			Nickname: "testuser",
		}

		// Expectations
		hashedEmail := "hashed_email"

		emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
		userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
		userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)

		// Execute
		output, err := userUsecase.SignUp(input)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password")
		assert.Nil(t, output)
	}

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_PasswordHashingFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := "hashed_email"

	emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	passwordHasher.On("HashPassword", input.Password).Return("", errors.New("hashing failed"))

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "hashing failed")
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_EmailEncryptionFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := "hashed_email"
	hashedPassword := "hashed_password"

	emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	passwordHasher.On("HashPassword", input.Password).Return(hashedPassword, nil)
	emailEncryptor.On("Encrypt", input.Email).Return("", errors.New("encryption failed"))

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "encryption failed")
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_EmailSendingFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := "hashed_email"
	hashedPassword := "hashed_password"
	encryptedEmail := "encrypted_email"

	emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	passwordHasher.On("HashPassword", input.Password).Return(hashedPassword, nil)
	emailEncryptor.On("Encrypt", input.Email).Return(encryptedEmail, nil)
	userRepo.On("Create", mock.AnythingOfType("*entities.User")).Return(nil)
	emailSender.On("SendEmail", input.Email, "Welcome to the sonic odyssey~!", "welcome.html", mock.AnythingOfType("email.WelcomeData")).Return(errors.New("sending failed"))

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sending failed")
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_UserCreationFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordHasher := &mocks.PasswordHasher{}
	emailHasher := &mocks.EmailHasher{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordHasher, emailHasher, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := "hashed_email"
	hashedPassword := "hashed_password"
	encryptedEmail := "encrypted_email"

	emailHasher.On("HashEmail", input.Email).Return(hashedEmail)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	passwordHasher.On("HashPassword", input.Password).Return(hashedPassword, nil)
	emailEncryptor.On("Encrypt", input.Email).Return(encryptedEmail, nil)
	userRepo.On("Create", mock.AnythingOfType("*entities.User")).Return(errors.New("creation failed"))

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "creation failed")
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	passwordHasher.AssertExpectations(t)
	emailHasher.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}
