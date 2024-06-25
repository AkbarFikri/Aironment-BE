package repository

import (
	"gorm.io/gorm"

	"github.com/AkbarFikri/Aironment-BE/internal/app/entity"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPost(DB *gorm.DB) PostRepository {
	return PostRepository{
		db: DB,
	}
}

func (r *PostRepository) FindByCategory(id int) ([]entity.Post, error) {
	var posts []entity.Post

	if err := r.db.Where("category = ?", id).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *PostRepository) FindAllImageByCommunity(id string) ([]entity.Post, error) {
	var posts []entity.Post

	if err := r.db.Where("community_id = ? AND image_url IS NOT NULL", id).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *PostRepository) Insert(Post entity.Post) error {
	if err := r.db.Create(&Post).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Update(Post entity.Post) error {
	if err := r.db.Where("id = ?", Post.ID).Save(&Post).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Delete(Post entity.Post) error {
	if err := r.db.Where("id = ?", Post.ID).Delete(&Post).Error; err != nil {
		return err
	}
	return nil
}
