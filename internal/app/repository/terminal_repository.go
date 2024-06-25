package repository

import (
	"github.com/AkbarFikri/Aironment-BE/internal/app/entity"
	"gorm.io/gorm"
)

type TerminalRepository struct {
	db *gorm.DB
}

func NewTerminal(DB *gorm.DB) TerminalRepository {
	return TerminalRepository{
		db: DB,
	}
}

func (r *TerminalRepository) Insert(terminal entity.Terminal) error {
	if err := r.db.Create(&terminal).Error; err != nil {
		return err
	}
	return nil
}

func (r *TerminalRepository) Update(terminal entity.Terminal) error {
	if err := r.db.Where("id = ?", terminal.ID).Save(&terminal).Error; err != nil {
		return err
	}
	return nil
}

func (r *TerminalRepository) Delete(terminal entity.Terminal) error {
	if err := r.db.Where("id = ?", terminal.ID).Delete(&terminal).Error; err != nil {
		return err
	}
	return nil
}
