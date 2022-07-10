package app

import (
	"fmt"
	"github.com/risangpratama13/sirka-test/model"
	"github.com/risangpratama13/sirka-test/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config util.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	return db
}
