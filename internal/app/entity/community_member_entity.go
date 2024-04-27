package entity

type Member struct {
	ID          string `gorm:"primaryKey"`
	UserID      string
	CommunityID string
}
