package server

import (
	"fmt"
	"os"

	"github.com/esc-chula/gearfest-backend/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadSupabase(config config.SupabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		os.Exit(0)
	}
	db = db.Exec(fmt.Sprintf("SET search_path TO %s", config.Schema)).Session(&gorm.Session{})
	return db
}
