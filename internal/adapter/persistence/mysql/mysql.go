package mysql

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"toktok-backend-v1.0.1/internal/config"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)

	gormConfig := &gorm.Config{
		TranslateError: true,
	}

	gormDB, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(config.Database.ConnMaxLifeTime))) * time.Millisecond)

	// new instance
	database := Database{
		DB: gormDB,
	}

	return &database, nil
}
