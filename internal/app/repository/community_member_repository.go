package repository

import (
	"gorm.io/gorm"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/entity"

)

type CommunityMemberRepository struct {
	db *gorm.DB
}

func NewCommunityMember(DB *gorm.DB) CommunityMemberRepository {
	return CommunityMemberRepository{
		db: DB,
	}
}

func (r *CommunityMemberRepository) FindByCommunityId(id string) ([]entity.Member, error) {
	var members []entity.Member
	if err := r.db.Where("community_id = ?", id).Find(&members).Error; err != nil {
		return members, err
	}
	return members, nil
}

func (r *CommunityMemberRepository) Insert(Member entity.Member) error {
	if err := r.db.Create(&Member).Error; err != nil {
		return err
	}
	return nil
}

func (r *CommunityMemberRepository) Update(Member entity.Member) error {
	if err := r.db.Where("id = ?", Member.ID).Save(&Member).Error; err != nil {
		return err
	}
	return nil
}

func (r *CommunityMemberRepository) Delete(Member entity.Member) error {
	if err := r.db.Where("id = ?", Member.ID).Delete(&Member).Error; err != nil {
		return err
	}
	return nil
}
