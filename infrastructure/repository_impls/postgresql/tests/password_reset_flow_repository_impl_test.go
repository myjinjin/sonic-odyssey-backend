package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestPasswordResetFlowRepository_Create(t *testing.T) {
	email := "passwordresetflow@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "passwordresetflow",
		Nickname:     "passwordresetflow",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	flowID := generateFlowID()

	testCases := []struct {
		name        string
		flow        *entities.PasswordResetFlow
		expectedErr error
	}{
		{
			name: "Success",
			flow: &entities.PasswordResetFlow{
				UserID: user.ID,
				FlowID: flowID,
				ExpiresAt: func() *time.Time {
					t := time.Now().Add(time.Hour * 2)
					return &t
				}(),
			},
			expectedErr: nil,
		},
		{
			name: "DuplicateFlowID",
			flow: &entities.PasswordResetFlow{
				UserID: user.ID,
				FlowID: flowID,
				ExpiresAt: func() *time.Time {
					t := time.Now().Add(time.Hour * 2)
					return &t
				}(),
			},
			expectedErr: repositories.ErrCreate,
		},
		{
			name: "ExpiresAtNotNull",
			flow: &entities.PasswordResetFlow{
				UserID: user.ID,
				FlowID: generateFlowID(),
			},
			expectedErr: repositories.ErrCreate,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := flowRepo.Create(tc.flow)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.PasswordResetFlow{})
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func TestPasswordResetFlowRepository_FindByFlowID(t *testing.T) {
	email := "passwordresetflow3@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "passwordresetflow3",
		Nickname:     "passwordresetflow3",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	paswordResetFlow := &entities.PasswordResetFlow{
		UserID: user.ID,
		FlowID: generateFlowID(),
		ExpiresAt: func() *time.Time {
			t := time.Now().Add(time.Hour * 2)
			return &t
		}(),
	}
	err = flowRepo.Create(paswordResetFlow)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		flowID      string
		expectedErr error
	}{
		{
			name:        "Success",
			flowID:      paswordResetFlow.FlowID,
			expectedErr: nil,
		},
		{
			name:        "NotFound",
			flowID:      "NOTEXIST",
			expectedErr: repositories.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := flowRepo.FindByFlowID(tc.flowID)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.PasswordResetFlow{})
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func TestPasswordResetFlowRepository_FindByUserID(t *testing.T) {
	email := "passwordresetflow4@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "passwordresetflow4",
		Nickname:     "passwordresetflow4",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	paswordResetFlow := &entities.PasswordResetFlow{
		UserID: user.ID,
		FlowID: generateFlowID(),
		ExpiresAt: func() *time.Time {
			t := time.Now().Add(time.Hour * 2)
			return &t
		}(),
	}
	err = flowRepo.Create(paswordResetFlow)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		userID      uint
		expectedErr error
	}{
		{
			name:        "Success",
			userID:      paswordResetFlow.UserID,
			expectedErr: nil,
		},
		{
			name:        "NotFound",
			userID:      100000,
			expectedErr: repositories.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := flowRepo.FindByUserID(tc.userID)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.PasswordResetFlow{})
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func TestPasswordResetFlowRepository_DeleteByFlowID(t *testing.T) {
	email := "passwordresetflow5@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "passwordresetflow5",
		Nickname:     "passwordresetflow5",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	paswordResetFlow := &entities.PasswordResetFlow{
		UserID: user.ID,
		FlowID: generateFlowID(),
		ExpiresAt: func() *time.Time {
			t := time.Now().Add(time.Hour * 2)
			return &t
		}(),
	}
	err = flowRepo.Create(paswordResetFlow)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		flowID      string
		expectedErr error
	}{
		{
			name:        "Success",
			flowID:      paswordResetFlow.FlowID,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := flowRepo.DeleteByFlowID(tc.flowID)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.PasswordResetFlow{})
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func generateFlowID() string {
	currentTime := time.Now().Unix()
	uuidValue := uuid.New()
	return fmt.Sprintf("%s:%d", uuidValue.String(), currentTime)
}
