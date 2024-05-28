package usecase

import (
	"testing"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_SignUp_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
	encryptedEmail := "encrypted_email"

	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	emailEncryptor.On("Encrypt", input.Email).Return(encryptedEmail, nil)

	userRepo.On("Create", mock.MatchedBy(func(user *entities.User) bool {
		user.ID = 1
		return user.Email == encryptedEmail &&
			user.EmailHash == hashedEmail &&
			hash.BCryptPasswordHasher().CheckPasswordHash(input.Password, user.PasswordHash) &&
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
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_EmailAlreadyExists(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
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
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_NicknameAlreadyExists(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
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
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_InvalidPassword(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, emailEncryptor, emailSender)

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
		hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
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
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_PasswordHashingFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "TOOLONGtofkjsldfjdsSsddsDGFdsfsdfsdfdfsVjlfdkjgvljkfdjkPassword123456789!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrHashingPassword.Error())
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_EmailEncryptionFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	emailEncryptor.On("Encrypt", input.Email).Return("", ErrEncryptingEmail)

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrEncryptingEmail.Error())
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}

func TestUserUsecase_SignUp_UserCreationFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, emailEncryptor, emailSender)

	input := SignUpInput{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "Test User",
		Nickname: "testuser",
	}

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(input.Email)
	encryptedEmail := "encrypted_email"

	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, nil)
	userRepo.On("FindByNickname", input.Nickname).Return(nil, nil)
	emailEncryptor.On("Encrypt", input.Email).Return(encryptedEmail, nil)
	userRepo.On("Create", mock.AnythingOfType("*entities.User")).Return(ErrCreatingRecord)

	// Execute
	output, err := userUsecase.SignUp(input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrCreatingRecord.Error())
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
	emailSender.AssertExpectations(t)
}
