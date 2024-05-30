package tests

import (
	"testing"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	existingUser := &entities.User{
		Email:        "existing@example.com",
		EmailHash:    "existinghashedemail",
		PasswordHash: "hashedPassword!",
		Name:         "existingname",
		Nickname:     "existingnickname",
	}
	err := userRepo.Create(existingUser)
	assert.NoError(t, err)

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

func TestUserRepository_FindByID(t *testing.T) {
	email := "findbyid@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "findbyid",
		Nickname:     "findbyid",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		userID      uint
		expectedErr error
	}{
		{
			name:        "Success",
			userID:      user.ID,
			expectedErr: nil,
		},
		{
			name:        "NotFound",
			userID:      10000,
			expectedErr: repositories.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := userRepo.FindByID(tc.userID)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func TestUserRepository_FindByEmailHash(t *testing.T) {
	email := "findbyemailhash@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "findbyemailhash",
		Nickname:     "findbyemailhash",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		emailHash   string
		expectedErr error
	}{
		{
			name:        "Success",
			emailHash:   user.EmailHash,
			expectedErr: nil,
		},
		{
			name:        "NotFound",
			emailHash:   "notexist",
			expectedErr: repositories.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := userRepo.FindByEmailHash(tc.emailHash)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func TestUserRepository_FindByNickname(t *testing.T) {
	email := "findbynickname@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "findbynickname",
		Nickname:     "findbynickname",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		nickname    string
		expectedErr error
	}{
		{
			name:        "Success",
			nickname:    user.Nickname,
			expectedErr: nil,
		},
		{
			name:        "NotFound",
			nickname:    "notexist",
			expectedErr: repositories.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := userRepo.FindByNickname(tc.nickname)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func TestUserRepository_Update(t *testing.T) {
	email := "userupdate@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "userupdate",
		Nickname:     "userupdate",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	email2 := "userupdate2@example.com"
	user2 := &entities.User{
		Email:        email2,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email2),
		PasswordHash: hashedPassword,
		Name:         "userupdate2",
		Nickname:     "userupdate2",
	}
	err = userRepo.Create(user2)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		user        *entities.User
		expectedErr error
	}{
		{
			name: "Success",
			user: &entities.User{
				ID:       user.ID,
				Name:     "userupdate1",
				Nickname: "userupdate1",
			},
			expectedErr: nil,
		},
		{
			name: "DuplicateEmail",
			user: &entities.User{
				ID:        user.ID,
				EmailHash: user2.EmailHash,
			},
			expectedErr: repositories.ErrUpdate,
		},
		{
			name: "DuplicateNickname",
			user: &entities.User{
				ID:       user.ID,
				Nickname: user2.Nickname,
			},
			expectedErr: repositories.ErrUpdate,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := userRepo.Update(tc.user)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}

func TestUserRepository_Delete(t *testing.T) {
	email := "delete@example.com"
	password := "Password123!"
	hashedPassword, _ := hash.BCryptPasswordHasher().HashPassword(password)
	user := &entities.User{
		Email:        email,
		EmailHash:    hash.SHA256EmailHasher().HashEmail(email),
		PasswordHash: hashedPassword,
		Name:         "delete",
		Nickname:     "delete",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		userID      uint
		expectedErr error
	}{
		{
			name:        "Success",
			userID:      user.ID,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := userRepo.Delete(tc.userID)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

	t.Cleanup(func() {
		testdb.GetDB().Unscoped().Where("1 = 1").Delete(&entities.User{})
	})
}
