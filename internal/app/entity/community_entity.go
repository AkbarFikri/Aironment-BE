package entity

type Community struct {
	ID             string `gorm:"primaryKey"`
	Name           string
	UserID         string
	InvoiceID      string
	Description    string
	ProfilePicture string
	CoverPicture   string
	Status         string
	Posts          []Post   `gorm:"foreignKey:CommunityID"`
	Members        []Member `gorm:"foreignKey:CommunityID"`
}
