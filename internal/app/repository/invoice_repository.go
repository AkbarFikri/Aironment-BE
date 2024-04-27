package repository

import (
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/entity"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoice(DB *gorm.DB) InvoiceRepository {
	return InvoiceRepository{
		db: DB,
	}
}

func (r *InvoiceRepository) Insert(Invoice entity.Invoice) error {
	if err := r.db.Create(&Invoice).Error; err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepository) Update(Invoice entity.Invoice) error {
	if err := r.db.Where("id = ?", Invoice.ID).Save(&Invoice).Error; err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepository) Delete(Invoice entity.Invoice) error {
	if err := r.db.Where("id = ?", Invoice.ID).Delete(&Invoice).Error; err != nil {
		return err
	}
	return nil
}
