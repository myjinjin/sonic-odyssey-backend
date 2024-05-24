package postgresql

import (
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type UserRepository struct {
	db            *gorm.DB
	encryptionKey string
}

func NewUserRepository(db *gorm.DB, encryptionKey string) repositories.UserRepository {
	return &UserRepository{db: db, encryptionKey: encryptionKey}
}

func (r *UserRepository) Create(user *entities.User) error {
	encryptedEmail, err := r.encryptEmail(user.Email)
	if err != nil {
		return err
	}
	user.Email = encryptedEmail

	return r.db.Create(user).Error
}
func (r *UserRepository) FindByID(id uint) (*entities.User, error) {
	user := new(entities.User)
	err := r.db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	user.Email, err = r.decryptEmail(user.Email)
	return user, err
}
func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	encryptedEmail, err := r.encryptEmail(email)
	if err != nil {
		return nil, err
	}

	user := new(entities.User)
	if err = r.db.Where("email = ?", encryptedEmail).First(user).Error; err != nil {
		return nil, err
	}
	user.Email = email
	return user, err
}

func (r *UserRepository) FindByNickname(nickname string) (*entities.User, error) {
	user := new(entities.User)
	err := r.db.Where("nickname = ?", nickname).First(user).Error
	if err != nil {
		return nil, err
	}
	user.Email, err = r.decryptEmail(user.Email)
	return user, err
}

func (r *UserRepository) Update(user *entities.User) error {
	encryptedEmail, err := r.encryptEmail(user.Email)
	if err != nil {
		return err
	}
	user.Email = encryptedEmail
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&entities.User{}, id).Error
}

func (r *UserRepository) encryptEmail(email string) (string, error) {
	var encryptedEmail string
	err := r.db.Raw("SELECT pgp_sym_encrypt(?, ?) AS encrypted_email", email, r.encryptionKey).Scan(&encryptedEmail).Error
	return encryptedEmail, err
}

func (r *UserRepository) decryptEmail(encryptedEmail string) (string, error) {
	var email string
	err := r.db.Raw("SELECT pgp_sym_decrypt(?::bytea, ?)::text AS email", encryptedEmail, r.encryptionKey).Scan(&email).Error
	return email, err
}
