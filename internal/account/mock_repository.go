package account

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type memoryRepository struct {
	accounts map[string]Account
	mutex    sync.RWMutex
}

func NewMemoryRepository() Repository {
	return &memoryRepository{
		accounts: make(map[string]Account),
	}
}

// Create implements Repository.
func (m *memoryRepository) Create(newAccount NewAccount) (Account, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	id, _ := uuid.NewRandom()
	_account := Account{
		ID:      id.String(),
		Balance: newAccount.Balance,
		OwnerID: newAccount.OwnerID,
	}

	m.accounts[id.String()] = _account
	return _account, nil
}

func (m *memoryRepository) UpdateBalance(id string, amount float32) (float32, error) {
	_account, found := m.accounts[id]
	if !found {
		return 0.0, fmt.Errorf("account with id: %s was not found", id)
	}

	_account.Balance = amount

	m.accounts[id] = _account

	return _account.Balance, nil
}

// Delete implements Repository.
func (m *memoryRepository) Delete(id string) error {
	panic("unimplemented")
}

// GetByID implements Repository.
func (m *memoryRepository) GetByID(id string) (Account, error) {
	_account, found := m.accounts[id]
	if !found {
		return Account{}, fmt.Errorf("account with id: %s was not found", id)
	}

	return _account, nil
}

func (m *memoryRepository) GetByOwnerID(ownerID string) (Account, error) {
	for _, value := range m.accounts {
		if value.OwnerID == ownerID {
			return value, nil
		}
	}

	return Account{}, fmt.Errorf("account with ownerID: %s was not found", ownerID)
}
