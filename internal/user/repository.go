package user

type Repository interface {
	Create(newUser NewUser) (User, error)
	GetByID(id string) (User, error)
	GetByEmail(email string) (User, error)
	Delete(id string) error
}
