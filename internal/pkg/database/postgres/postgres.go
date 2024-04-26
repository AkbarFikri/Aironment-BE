package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

func NewPostgres() *gorm.DB {
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBSslMode := os.Getenv("SSL_MODE")

	DBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, DBSslMode,
	)

	db, err := gorm.Open(postgres.Open(DBDSN), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}
	migrate(db)
	return db
}

func migrate(db *gorm.DB) {

}
