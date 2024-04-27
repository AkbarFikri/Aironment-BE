package entity

type Post struct {
	ID          string `gorm:"primaryKey"`
	CommunityID string
	UserID      string
	Title       string
	Description string
	ImageUrl    string
	Community   Community
}
