package repositories

import (
	"fmt"
	"letsgo-flight-provider/internal/core/entities"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGRepository struct {
	DB *gorm.DB
}

func NewGormDatabase() *PGRepository {
	db, _ := GormInit()
	return &PGRepository{DB: db}
}

func GormInit() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.Flight{})
	if err != nil {
		return nil, err
	}

	var count int64
	db.Model(&entities.Flight{}).Count(&count)
	if count == 0 {
		err := runSQLFile(db)
		if err != nil {
			return nil, err
		}

	}
	return db, nil
}

func runSQLFile(db *gorm.DB) error {
	filePath := filepath.Join("flights.sql")

	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = db.Exec(string(sqlBytes)).Error
	if err != nil {
		return err
	}

	return nil
}
