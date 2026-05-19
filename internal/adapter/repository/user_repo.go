package repository

import (
	"be-ayaka/internal/core/customerrors"
	"be-ayaka/internal/core/entity"
	"be-ayaka/internal/core/port"
	"context"
	"errors"

	"gorm.io/gorm"
)

type userRepoPostgres struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.UserRepository {
	return &userRepoPostgres{
		db: db,
	}
}

func (r *userRepoPostgres) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrDataNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepoPostgres) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrInvalidCredentials
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepoPostgres) Create(ctx context.Context, user *entity.User) error {
	db := ExtractTx(ctx, r.db)
	return db.WithContext(ctx).Create(user).Error
}
