package api

import (
	"fmt"
	"kredit/backend/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	host := os.Getenv("host")
	port := os.Getenv("port_db")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if os.Getenv("AUTO_MIGRATE") == "Y" {
		if err := db.AutoMigrate(model.BranchTab{}, model.CustomerDataTab{}, model.LoanDataTab{}, model.MstCompanyTab{}, model.SkalaRentalTab{}, model.StagingCustomer{}, model.StagingError{}, model.VehicleDataTab{}, model.GeneralParameter{}); err != nil {
			return nil, fmt.Errorf("failed to migrate database: %w", err)
		}
	}

	return db, err
}
