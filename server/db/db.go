package db

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	// "time"

	"github.com/joho/godotenv"
	"github.com/vijaymehrotra/go-next-ts_chat/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {										
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=%s user=%s",
		config.Host, config.Port, config.Password, config.DBName, config.SSLMode, config.User)

	log.Println("Connecting to database...")
	var db *gorm.DB
	var err error
	
	// for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// if err == nil {
		// 	break
		// }
		// time.Sleep(2 * time.Second)
	// }
	
	if err != nil {
		return nil, err
	}
	DB=db
	log.Println("Database connected successfully")
	return db, nil
}

func InitilizeDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		log.Fatal(err)
	}

	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db,err := NewConnection(*config)
	if err != nil {
		log.Fatal(err)
	}

	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatal(err)
	}
}