package repository

import (
	"gorm.io/gorm"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/entity"
)

type CommunityRepository struct {
	db *gorm.DB
}

func NewCommunity(DB *gorm.DB) CommunityRepository {
	return CommunityRepository{
		db: DB,
	}
}

func (r *CommunityRepository) FindById(id string) (entity.Community, error) {
	var community entity.Community
	if err := r.db.Where("id = ?", id).First(&community).Error; err != nil {
		return community, err
	}
	return community, nil
}

func (r *CommunityRepository) FindAll() ([]entity.Community, error) {
	var community []entity.Community
	if err := r.db.Where("status = verified").Find(&community).Error; err != nil {
		return community, err
	}
	return community, nil
}

func (r *CommunityRepository) FindByInvoiceId(id string) (entity.Community, error) {
	var community entity.Community
	if err := r.db.Where("invoice_id = ?", id).First(&community).Error; err != nil {
		return community, err
	}
	return community, nil
}

func (r *CommunityRepository) Insert(Community entity.Community) error {
	if err := r.db.Create(&Community).Error; err != nil {
		return err
	}
	return nil
}

func (r *CommunityRepository) Update(Community entity.Community) error {
	if err := r.db.Where("id = ?", Community.ID).Save(&Community).Error; err != nil {
		return err
	}
	return nil
}

func (r *CommunityRepository) Delete(Community entity.Community) error {
	if err := r.db.Where("id = ?", Community.ID).Delete(&Community).Error; err != nil {
		return err
	}
	return nil
}
