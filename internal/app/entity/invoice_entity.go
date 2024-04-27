package entity

type Invoice struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Amount    uint64
	Purpose   string
	Status    string
	Community Community `gorm:"foreignKey:InvoiceID"`
}
