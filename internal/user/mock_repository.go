package user

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type memoryRepository struct {
	users map[string]User
	mutex sync.RWMutex
}

func NewMemoryRepository() Repository {
	return &memoryRepository{
		users: make(map[string]User),
	}
}

// Create implements Repository.
func (m *memoryRepository) Create(newUser NewUser) (User, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	id, _ := uuid.NewRandom()
	_user := User{
		ID:          id.String(),
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
	}

	m.users[id.String()] = _user
	return _user, nil
}

// GetByEmail implements Repository.
func (m *memoryRepository) GetByEmail(email string) (User, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, value := range m.users {
		if value.Email == email {
			return value, nil
		}
	}

	return User{}, fmt.Errorf("user with email: %s was not found", email)
}

// GetByID implements Repository.
func (m *memoryRepository) GetByID(id string) (User, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_user, found := m.users[id]
	if !found {
		return User{}, fmt.Errorf("user with id: %s not found", id)
	}

	return _user, nil
}

// Delete implements Repository.
func (m *memoryRepository) Delete(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, found := m.users[id]
	if !found {
		return fmt.Errorf("user with id: %s not found", id)
	}

	m.users[id] = User{}

	return nil
}
