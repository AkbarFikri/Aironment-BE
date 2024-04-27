package entity

type Post struct {
	ID          string `gorm:"primaryKey"`
	Category    int
	CommunityID string
	UserID      string
	Link        string
	Title       string
	Description string
	ImageUrl    string
	Community   Community
}
