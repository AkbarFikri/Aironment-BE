package entity

type Terminal struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Latitude  string
	Longitude string
}
