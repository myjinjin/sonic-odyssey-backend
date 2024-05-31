package usecase

import (
	"testing"
	"time"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/email"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_SignUp_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

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

	emailSender.On("SendEmail", input.Email, email.TemplateWelcome, mock.AnythingOfType("email.WelcomeData")).Return(nil)

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
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

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
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

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
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

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
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

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
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

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
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

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

func TestUserUsecase_SendPasswordRecoveryEmail_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	userEmail := "test@example.com"
	baseURL := "http://localhost:8080"

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(userEmail)
	user := &entities.User{ID: 1, Name: "Test User"}

	userRepo.On("FindByEmailHash", hashedEmail).Return(user, nil)
	passwordResetRepo.On("Create", mock.AnythingOfType("*entities.PasswordResetFlow")).Return(nil)
	emailSender.On("SendEmail", userEmail, email.TemplatePasswordReset, mock.AnythingOfType("email.PasswordResetData")).Return(nil)

	// Execute
	err := userUsecase.SendPasswordRecoveryEmail(baseURL, userEmail)
	assert.NoError(t, err)

	// Verify
	userRepo.AssertExpectations(t)
	passwordResetRepo.AssertExpectations(t)
}

func TestUserUsecase_SendPasswordRecoveryEmail_UserNotFound(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	userEmail := "test@example.com"
	baseURL := "http://localhost:8080"

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(userEmail)

	userRepo.On("FindByEmailHash", hashedEmail).Return(nil, repositories.ErrNotFound)

	// Execute
	err := userUsecase.SendPasswordRecoveryEmail(baseURL, userEmail)

	// Assert
	assert.ErrorIs(t, err, ErrUserNotFound)

	// Verify
	userRepo.AssertExpectations(t)
	passwordResetRepo.AssertNotCalled(t, "Create")
}

func TestUserUsecase_SendPasswordRecoveryEmail_CreateResetFlowError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	userEmail := "test@example.com"
	baseURL := "http://localhost:8080"

	// Expectations
	hashedEmail := hash.SHA256EmailHasher().HashEmail(userEmail)
	user := &entities.User{ID: 1, Name: "Test User"}

	userRepo.On("FindByEmailHash", hashedEmail).Return(user, nil)
	passwordResetRepo.On("Create", mock.AnythingOfType("*entities.PasswordResetFlow")).Return(repositories.ErrCreate)

	// Execute
	err := userUsecase.SendPasswordRecoveryEmail(baseURL, userEmail)
	assert.ErrorIs(t, err, ErrCreatingRecord)

	// Verify
	userRepo.AssertExpectations(t)
	passwordResetRepo.AssertExpectations(t)
}

func TestUserUsecase_ResetPassword_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	password := "newPassword123!"
	flowID := "flow123"
	user := &entities.User{ID: 1, Name: "Test User"}
	flow := &entities.PasswordResetFlow{
		ID:        1,
		UserID:    user.ID,
		User:      *user,
		FlowID:    flowID,
		ExpiresAt: func() *time.Time { t := time.Now().Add(time.Hour); return &t }(),
	}

	// Expectations
	passwordResetRepo.On("FindByFlowID", flowID).Return(flow, nil)
	userRepo.On("Update", mock.AnythingOfType("*entities.User")).Return(nil).Run(func(args mock.Arguments) {
		updatedUser := args.Get(0).(*entities.User)
		user.PasswordHash = updatedUser.PasswordHash
	})
	passwordResetRepo.On("DeleteByFlowID", flowID).Return(nil)

	// Execute
	err := userUsecase.ResetPassword(password, flowID)
	assert.NoError(t, err)

	// Verify
	passwordResetRepo.AssertExpectations(t)
	userRepo.AssertExpectations(t)
	assert.True(t, hash.BCryptPasswordHasher().CheckPasswordHash(password, user.PasswordHash))
}

func TestUserUsecase_ResetPassword_FlowNotFound(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	password := "newPassword123!"
	flowID := "flow123"

	// Expectations
	passwordResetRepo.On("FindByFlowID", flowID).Return(nil, repositories.ErrNotFound)

	// Execute
	err := userUsecase.ResetPassword(password, flowID)
	assert.Error(t, err)
	assert.Equal(t, ErrPasswordResetFlowNotFound, err)

	// Verify
	passwordResetRepo.AssertExpectations(t)
}

func TestUserUsecase_ResetPassword_FlowExpired(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	password := "newPassword123!"
	flowID := "flow123"
	user := &entities.User{ID: 1, Name: "Test User"}
	flow := &entities.PasswordResetFlow{
		ID:        1,
		UserID:    user.ID,
		User:      *user,
		FlowID:    flowID,
		ExpiresAt: func() *time.Time { t := time.Now().Add(-time.Hour); return &t }(),
	}

	// Expectations
	passwordResetRepo.On("FindByFlowID", flowID).Return(flow, nil)

	// Execute
	err := userUsecase.ResetPassword(password, flowID)
	assert.Error(t, err)
	assert.Equal(t, ErrPasswordResetFlowExpired, err)

	// Verify
	passwordResetRepo.AssertExpectations(t)
}

func TestUserUsecase_ResetPassword_InvalidPassword(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	password := "short"
	flowID := "flow123"
	user := &entities.User{ID: 1, Name: "Test User"}
	flow := &entities.PasswordResetFlow{
		ID:        1,
		UserID:    user.ID,
		User:      *user,
		FlowID:    flowID,
		ExpiresAt: func() *time.Time { t := time.Now().Add(time.Hour); return &t }(),
	}

	// Expectations
	passwordResetRepo.On("FindByFlowID", flowID).Return(flow, nil)

	// Execute
	err := userUsecase.ResetPassword(password, flowID)
	assert.Error(t, err)
	assert.Equal(t, ErrPasswordTooShort, err)

	// Verify
	passwordResetRepo.AssertExpectations(t)
}

func TestUserUsecase_ResetPassword_DeleteFlowError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	passwordResetRepo := &mocks.PasswordResetFlowRepository{}
	emailEncryptor := &mocks.Encryptor{}
	emailSender := &mocks.EmailSender{}

	userUsecase := NewUserUsecase(userRepo, passwordResetRepo, emailEncryptor, emailSender)

	password := "newPassword123!"
	flowID := "flow123"
	user := &entities.User{ID: 1, Name: "Test User"}
	flow := &entities.PasswordResetFlow{
		ID:        1,
		UserID:    user.ID,
		User:      *user,
		FlowID:    flowID,
		ExpiresAt: func() *time.Time { t := time.Now().Add(time.Hour); return &t }(),
	}

	// Expectations
	passwordResetRepo.On("FindByFlowID", flowID).Return(flow, nil)
	userRepo.On("Update", mock.AnythingOfType("*entities.User")).Return(nil).Run(func(args mock.Arguments) {
		updatedUser := args.Get(0).(*entities.User)
		user.PasswordHash = updatedUser.PasswordHash
	})
	passwordResetRepo.On("DeleteByFlowID", flowID).Return(repositories.ErrDelete)

	// Execute
	err := userUsecase.ResetPassword(password, flowID)
	assert.Error(t, err)
	assert.Equal(t, ErrDeletingRecord, err)

	// Verify
	passwordResetRepo.AssertExpectations(t)
	userRepo.AssertExpectations(t)
}
