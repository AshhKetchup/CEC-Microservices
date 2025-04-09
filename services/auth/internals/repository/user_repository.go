package repository

import (
	"context"

	"github.com/yourproject/services/auth/internal/model"
	"github.com/yourproject/services/auth/internal/storage"
)

type FileUserRepository struct {
	store *storage.FileStore
}

func NewUserRepository(store *storage.FileStore) *FileUserRepository {
	return &FileUserRepository{store: store}
}

func (r *FileUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.store.SaveUser(user)
}

func (r *FileUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return r.store.FindUserByEmail(email)
}
