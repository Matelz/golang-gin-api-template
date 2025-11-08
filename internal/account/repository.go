package account

type Repository interface {
	Create(newAccount NewAccount) (Account, error)
	GetByID(id string) (Account, error)
	GetByOwnerID(ownerID string) (Account, error)
	UpdateBalance(id string, amount float32) (float32, error)
	Delete(id string) error
}
