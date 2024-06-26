package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/email"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase/mocks"
	"github.com/myjinjin/sonic-odyssey-backend/pkg/utils"
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

func TestUserUsecase_GetUserByID_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}

	userUsecase := NewUserUsecase(userRepo, nil, emailEncryptor, nil)

	userID := uint(1)
	encryptedEmail := "encrypted_email"
	decryptedEmail := "test@example.com"

	user := &entities.User{
		ID:       userID,
		Email:    encryptedEmail,
		Name:     "Test User",
		Nickname: "testuser",
		UserProfile: &entities.UserProfile{
			ProfileImageURL: "https://example.com/profile.png",
			Bio:             "Test bio",
			Website:         "https://example.com",
		},
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(user, nil)
	emailEncryptor.On("Decrypt", encryptedEmail).Return(decryptedEmail, nil)

	// Execute
	output, err := userUsecase.GetUserByID(userID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, userID, output.ID)
	assert.Equal(t, decryptedEmail, output.Email)
	assert.Equal(t, user.Name, output.Name)
	assert.Equal(t, user.Nickname, output.Nickname)
	assert.Equal(t, user.UserProfile.ProfileImageURL, output.ProfileImageURL)
	assert.Equal(t, user.UserProfile.Bio, output.Bio)
	assert.Equal(t, user.UserProfile.Website, output.Website)

	// Verify
	userRepo.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
}

func TestUserUsecase_GetUserByID_UserNotFound(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}

	userUsecase := NewUserUsecase(userRepo, nil, emailEncryptor, nil)

	userID := uint(1)

	// Expectations
	userRepo.On("FindByID", userID).Return(nil, repositories.ErrNotFound)

	// Execute
	output, err := userUsecase.GetUserByID(userID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	emailEncryptor.AssertNotCalled(t, "Decrypt")
}

func TestUserUsecase_GetUserByID_FindError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}

	userUsecase := NewUserUsecase(userRepo, nil, emailEncryptor, nil)

	userID := uint(1)

	// Expectations
	userRepo.On("FindByID", userID).Return(nil, repositories.ErrFind)

	// Execute
	output, err := userUsecase.GetUserByID(userID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrFindingRecord, err)
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	emailEncryptor.AssertNotCalled(t, "Decrypt")
}

func TestUserUsecase_GetUserByID_DecryptionError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	emailEncryptor := &mocks.Encryptor{}

	userUsecase := NewUserUsecase(userRepo, nil, emailEncryptor, nil)

	userID := uint(1)
	encryptedEmail := "encrypted_email"

	user := &entities.User{
		ID:    userID,
		Email: encryptedEmail,
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(user, nil)
	emailEncryptor.On("Decrypt", encryptedEmail).Return("", errors.New("decryption error"))

	// Execute
	output, err := userUsecase.GetUserByID(userID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrDecryptingEmail, err)
	assert.Nil(t, output)

	// Verify
	userRepo.AssertExpectations(t)
	emailEncryptor.AssertExpectations(t)
}

func TestUserUsecase_PatchUser_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := &PatchUserInput{
		Name:     utils.ToPtr("Updated Name"),
		Nickname: utils.ToPtr("updatednickname"),
		Bio:      utils.ToPtr("Updated Bio"),
		Website:  utils.ToPtr("https://example.com"),
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(&entities.User{
		ID:       userID,
		Name:     "Test User",
		Nickname: "testuser",
		UserProfile: &entities.UserProfile{
			Bio:     "Test Bio",
			Website: "https://test.com",
		},
	}, nil)
	userRepo.On("FindByNickname", *input.Nickname).Return(&entities.User{
		ID: userID,
	}, nil)
	userRepo.On("Update", mock.MatchedBy(func(user *entities.User) bool {
		return user.ID == userID &&
			user.Name == *input.Name &&
			user.Nickname == *input.Nickname &&
			user.UserProfile.Bio == *input.Bio &&
			user.UserProfile.Website == *input.Website
	})).Return(nil)

	// Execute
	err := userUsecase.PatchUser(userID, input)

	// Assert
	assert.NoError(t, err)

	// Verify
	userRepo.AssertExpectations(t)
}

func TestUserUsecase_PatchUser_UserNotFound(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := &PatchUserInput{
		Name: utils.ToPtr("Updated Name"),
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(nil, repositories.ErrNotFound)

	// Execute
	err := userUsecase.PatchUser(userID, input)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)

	// Verify
	userRepo.AssertExpectations(t)
}

func TestUserUsecase_PatchUser_FindingRecordError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := &PatchUserInput{
		Name: utils.ToPtr("Updated Name"),
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(nil, repositories.ErrFind)

	// Execute
	err := userUsecase.PatchUser(userID, input)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrFindingRecord, err)

	// Verify
	userRepo.AssertExpectations(t)
}

func TestUserUsecase_PatchUser_UpdatingRecordError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := &PatchUserInput{
		Name: utils.ToPtr("Updated Name"),
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(&entities.User{
		ID:   userID,
		Name: "Test User",
	}, nil)

	userRepo.On("Update", mock.Anything).Return(repositories.ErrUpdate)

	// Execute
	err := userUsecase.PatchUser(userID, input)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrUpdatingRecord, err)

	// Verify
	userRepo.AssertExpectations(t)
}

func TestUserUsecase_UpdatePassword_Success(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := UpdatePasswordInput{
		UserID:       userID,
		CurrPassword: "currentPassword1!",
		NewPassword:  "newPassword1!",
	}
	hashedCurrPassword, _ := hash.BCryptPasswordHasher().HashPassword(input.CurrPassword)

	// Expectations
	userRepo.On("FindByID", userID).Return(&entities.User{
		ID: userID, PasswordHash: hashedCurrPassword}, nil)
	userRepo.On("Update", mock.MatchedBy(func(user *entities.User) bool {
		hashedNewPassword, _ := hash.BCryptPasswordHasher().HashPassword(input.NewPassword)
		return user.ID == userID && hash.BCryptPasswordHasher().CheckPasswordHash(input.NewPassword, hashedNewPassword)
	})).Return(nil)

	// Execute
	err := userUsecase.UpdatePassword(input)

	// Assert
	assert.NoError(t, err)

	// Verify
	userRepo.AssertExpectations(t)
}

func TestUserUsecase_UpdatePassword_UserNotFound(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := UpdatePasswordInput{
		UserID:       userID,
		CurrPassword: "currentPassword1!",
		NewPassword:  "newPassword1!",
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(nil, repositories.ErrNotFound)

	// Execute
	err := userUsecase.UpdatePassword(input)

	// Assert
	assert.ErrorIs(t, err, ErrUserNotFound)

	// Verify
	userRepo.AssertExpectations(t)
	userRepo.AssertNotCalled(t, "Update")
}

func TestUserUsecase_UpdatePassword_FindingRecordError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := UpdatePasswordInput{
		UserID:       userID,
		CurrPassword: "currentPassword1!",
		NewPassword:  "newPassword1!",
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(nil, repositories.ErrFind)

	// Execute
	err := userUsecase.UpdatePassword(input)

	// Assert
	assert.ErrorIs(t, err, ErrFindingRecord)

	// Verify
	userRepo.AssertExpectations(t)
	userRepo.AssertNotCalled(t, "Update")
}

func TestUserUsecase_UpdatePassword_PasswordNotMatched(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := UpdatePasswordInput{
		UserID:       userID,
		CurrPassword: "currentPassword1!",
		NewPassword:  "newPassword1!",
	}

	// Expectations
	userRepo.On("FindByID", userID).Return(&entities.User{
		ID: userID, PasswordHash: "SOMETHINGINVALID"}, nil)

	// Execute
	err := userUsecase.UpdatePassword(input)

	// Assert
	assert.ErrorIs(t, err, ErrPasswordNotMatched)

	// Verify
	userRepo.AssertExpectations(t)
	userRepo.AssertNotCalled(t, "Update")
}
func TestUserUsecase_UpdatePassword_PasswordHashingFailed(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := UpdatePasswordInput{
		UserID:       userID,
		CurrPassword: "currentPassword1!",
		NewPassword:  "TOOLONGtofkjsldfjdsSsddsDGFdsfsdfsdfdfsVjlfdkjgvljkfdjkPassword123456789!",
	}
	hashedCurrPassword, _ := hash.BCryptPasswordHasher().HashPassword(input.CurrPassword)

	// Expectations
	userRepo.On("FindByID", userID).Return(&entities.User{
		ID: userID, PasswordHash: hashedCurrPassword}, nil)

	// Execute
	err := userUsecase.UpdatePassword(input)

	// Assert
	assert.ErrorIs(t, err, ErrHashingPassword)

	// Verify
	userRepo.AssertExpectations(t)
	userRepo.AssertNotCalled(t, "Update")
}

func TestUserUsecase_UpdatePassword_InvalidPassword(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := UpdatePasswordInput{
		UserID:       userID,
		CurrPassword: "currentPassword1!",
		NewPassword:  "short",
	}
	hashedCurrPassword, _ := hash.BCryptPasswordHasher().HashPassword(input.CurrPassword)

	// Expectations
	userRepo.On("FindByID", userID).Return(&entities.User{
		ID: userID, PasswordHash: hashedCurrPassword}, nil)

	// Execute
	err := userUsecase.UpdatePassword(input)

	// Assert
	assert.ErrorIs(t, err, ErrPasswordTooShort)

	// Verify
	userRepo.AssertExpectations(t)
	userRepo.AssertNotCalled(t, "Update")
}

func TestUserUsecase_UpdatePassword_UpdatingError(t *testing.T) {
	// Setup
	userRepo := &mocks.UserRepository{}
	userUsecase := NewUserUsecase(userRepo, nil, nil, nil)

	userID := uint(1)
	input := UpdatePasswordInput{
		UserID:       userID,
		CurrPassword: "currentPassword1!",
		NewPassword:  "newPassword1!",
	}
	hashedCurrPassword, _ := hash.BCryptPasswordHasher().HashPassword(input.CurrPassword)

	// Expectations
	userRepo.On("FindByID", userID).Return(&entities.User{
		ID: userID, PasswordHash: hashedCurrPassword}, nil)
	userRepo.On("Update", mock.MatchedBy(func(user *entities.User) bool {
		hashedNewPassword, _ := hash.BCryptPasswordHasher().HashPassword(input.NewPassword)
		return user.ID == userID && hash.BCryptPasswordHasher().CheckPasswordHash(input.NewPassword, hashedNewPassword)
	})).Return(repositories.ErrUpdate)

	// Execute
	err := userUsecase.UpdatePassword(input)

	// Assert
	assert.ErrorIs(t, err, ErrUpdatingRecord)

	// Verify
	userRepo.AssertExpectations(t)
}
