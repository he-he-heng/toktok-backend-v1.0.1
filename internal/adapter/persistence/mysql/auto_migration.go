package mysql

import "toktok-backend-v1.0.1/internal/core/domain"

func (d *Database) AutoMigrate() error {
	return d.DB.AutoMigrate(&domain.User{}, &domain.Avatar{}, &domain.Friendship{}, &domain.Message{})
}
