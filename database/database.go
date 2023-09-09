package database

import (
	"fmt"
	"log"
	"os"
	"sustain/repositories"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
		viper.GetInt("db.port"),
		viper.GetString("db.sslmode"),
		viper.GetString("db.timezone"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		DryRun: false,
	})
	if err != nil {
		fmt.Println("FAIL !!")
		log.Fatal(err)
	}

	table := fmt.Sprintf("set search_path=%v", viper.GetString("db.tablename"))
	db.Exec(table)

	return db, nil
}

func AutoMigrate(am *gorm.DB) {
	am.Debug().AutoMigrate(
		&repositories.Users{},
		&repositories.Groups{},
		&repositories.GroupUserReletion{},
	)
}
