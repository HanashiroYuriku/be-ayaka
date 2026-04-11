package database

import (
	"fmt"
	"log"

	"be-ayaka/config"
	ayaka "be-ayaka/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// NewPostgresConnection creates a new PostgreSQL connection using GORM
func NewPostgresConnection(cfg *config.Config) *gorm.DB {
	// build dsn
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	// open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})

	if err != nil {
		ayaka.Log("SYSTEM", "ERROR", fmt.Sprintf("Failed to connect to database: %v", err))
		log.Fatal("Failed to connect to database: ", err)
	}

	ayaka.Log("SYSTEM", "INFOR", "Success connect to database!")

	return db
}