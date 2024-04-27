package entity

type User struct {
	ID         string `gorm:"primaryKey"`
	Email      string `gorm:"not null;unique"`
	Password   string
	FullName   string
	Posts      []Post      `gorm:"foreignKey:UserID"`
	Communitys []Community `gorm:"foreignKey:UserID"`
	Members    []Member    `gorm:"foreignKey:UserID"`
}
