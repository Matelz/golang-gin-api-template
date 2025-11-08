package user

import (
	"regexp"

	"devmaua.com/devbank/internal/shared"
)

type Service struct {
	repo Repository
}

var phoneRegex = regexp.MustCompile(`(?:(?:(\+|00)?(55))\s?)?(?:\(?(\d{2})\)?\s?)(|\d{2})(|-)?(?:(9\d|[2-9])\d{3}[-|.|\s]?(\d{4}))`)
var emailRegex = regexp.MustCompile(`([\w._%+-]+)(@|\s@\s|\sat\s|\[at\])([\w.-]+)\.([\w]{2,})`)

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateNewUser(newUser NewUser) (User, error) {
	u, _ := s.repo.GetByEmail(newUser.Email)
	if (u != User{}) {
		return u, nil
	}

	if newUser.FirstName == "" {
		return User{}, shared.ErrFieldRequired("first_name")
	}

	if newUser.LastName == "" {
		return User{}, shared.ErrFieldRequired("last_name")
	}

	if newUser.PhoneNumber == "" {
		return User{}, shared.ErrFieldRequired("phone_number")
	} else if !phoneRegex.Match([]byte(newUser.PhoneNumber)) {
		return User{}, shared.ErrInvalidField("phone_number")
	}

	if newUser.Email == "" {
		return User{}, shared.ErrFieldRequired("email")
	} else if !emailRegex.Match([]byte(newUser.Email)) {
		return User{}, shared.ErrInvalidField("email")
	}

	return s.repo.Create(newUser)
}

func (s *Service) GetUserByID(id string) (User, error) {
	if id == "" {
		return User{}, shared.ErrFieldRequired("id")
	}

	return s.repo.GetByID(id)
}

func (s *Service) UserExists(id string) bool {
	u, err := s.repo.GetByID(id)
	if (err != nil || u == User{}) {
		return false
	}

	return true
}
