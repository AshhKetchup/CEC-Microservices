package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/yourproject/services/auth/internal/model"
)

type FileStore struct {
	mu       sync.Mutex
	filePath string
}

func NewFileStore(dataDir string) *FileStore {
	return &FileStore{
		filePath: dataDir + "/users.json",
	}
}

func (fs *FileStore) SaveUser(user *model.User) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	users, err := fs.readUsers()
	if err != nil {
		return err
	}

	// Check for existing user
	for _, u := range users {
		if u.Email == user.Email {
			return ErrUserExists
		}
	}

	users = append(users, *user)
	return fs.writeUsers(users)
}

func (fs *FileStore) FindUserByEmail(email string) (*model.User, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	users, err := fs.readUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, ErrUserNotFound
}

func (fs *FileStore) readUsers() ([]model.User, error) {
	var users []model.User

	file, err := os.Open(fs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	return users, err
}

func (fs *FileStore) writeUsers(users []model.User) error {
	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(users)
}

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)
