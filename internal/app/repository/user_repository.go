package repository

import (
	"gorm.io/gorm"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/entity"

)

type UserRepository struct {
	db *gorm.DB
}

func NewUser(DB *gorm.DB) UserRepository {
	return UserRepository{
		db: DB,
	}
}

func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindById(id string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Insert(user entity.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(user entity.User) error {
	if err := r.db.Where("id = ?", user.ID).Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(user entity.User) error {
	if err := r.db.Where("id = ?", user.ID).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}