package tests

import (
	"testing"

	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	t.Run("Setup", func(t *testing.T) {
		existingUser := &entities.User{
			Email:        "existing@example.com",
			EmailHash:    "existinghashedemail",
			PasswordHash: "hashedPassword!",
			Name:         "existingname",
			Nickname:     "existingnickname",
		}
		err := userRepo.Create(existingUser)
		assert.NoError(t, err)
	})

	testCases := []struct {
		name        string
		user        *entities.User
		expectedErr error
	}{
		{
			name: "Success",
			user: &entities.User{
				Email:        "useremail1@example.com",
				EmailHash:    "useremail1hashed",
				PasswordHash: "hashedPassword1!",
				Name:         "username1",
				Nickname:     "usernickname1",
			},
			expectedErr: nil,
		},
		{
			name: "DuplicateEmail",
			user: &entities.User{
				Email:        "existing@example.com",
				EmailHash:    "existinghashedemail",
				PasswordHash: "hashedPassword2!",
				Name:         "username2",
				Nickname:     "usernickname2",
			},
			expectedErr: repositories.ErrCreate,
		},
		{
			name: "DuplicateNickname",
			user: &entities.User{
				Email:        "useremail3@example.com",
				EmailHash:    "useremail3hashed",
				PasswordHash: "hashedPassword3!",
				Name:         "username3",
				Nickname:     "existingnickname",
			},
			expectedErr: repositories.ErrCreate,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := userRepo.Create(tc.user)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}
