package account

import (
	"errors"
	"fmt"

	"devmaua.com/devbank/internal/shared"
	"devmaua.com/devbank/internal/user"
)

type Service struct {
	repo        Repository
	userService user.Service
}

func NewService(repo Repository, userService user.Service) *Service {
	return &Service{
		repo:        repo,
		userService: userService,
	}
}

func (s *Service) OpenAccount(ownerId string) (Account, error) {
	if ownerId == "" {
		return Account{}, shared.ErrFieldRequired("owner_id")
	}

	a, _ := s.repo.GetByOwnerID(ownerId)
	if (a != Account{}) {
		return a, nil
	}

	if !s.userService.UserExists(ownerId) {
		return Account{}, fmt.Errorf("user with id: %s does not exist", ownerId)
	}

	return s.repo.Create(NewAccount{
		Balance: 0.0,
		OwnerID: ownerId,
	})
}

func (s *Service) GetAccount(accountID string) (Account, error) {
	if accountID == "" {
		return Account{}, shared.ErrFieldRequired("account_id")
	}

	return s.repo.GetByID(accountID)
}

func (s *Service) DepositAmount(accountID string, amount float32) (float32, error) {
	if accountID == "" {
		return 0.0, shared.ErrFieldRequired("account_id")
	}

	if amount <= 0.0 {
		return 0.0, shared.ErrInvalidField("amount")
	}

	_account, err := s.repo.GetByID(accountID)
	if err != nil {
		return 0.0, err
	}

	updated_balance, err := s.repo.UpdateBalance(accountID, _account.Balance+amount)
	if err != nil {
		return 0.0, nil
	}

	return updated_balance, nil
}

func (s *Service) WithdrawAmount(accountID string, amount float32) (float32, error) {
	if accountID == "" {
		return 0.0, shared.ErrFieldRequired("account_id")
	}

	if amount <= 0.0 {
		return 0.0, shared.ErrInvalidField("amount")
	}

	_account, err := s.repo.GetByID(accountID)
	if err != nil {
		return 0.0, err
	}

	if amount > _account.Balance {
		return 0.0, errors.New("not enough balance to withdraw")
	}

	updated_balance, err := s.repo.UpdateBalance(accountID, _account.Balance-amount)
	if err != nil {
		return 0.0, nil
	}

	return updated_balance, nil
}
