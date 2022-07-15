package user

import (
	pass "github.com/BetuelSA/go-helpers/password"
)

// Usecase represents users usecases
// Extends User entity interface
type Usecase interface {
	Repository
	Login(email, password string) (*User, error)
	ChangePassword(id uint, oldPassword, newPassword string) (*User, error)
}

type usecase struct {
	repository  Repository
	passwordSvc pass.Service
}

// NewUsecase creates a new usecase. Implements the Usecase interface
func NewUsecase(repo Repository, passSvc pass.Service) Usecase {
	return &usecase{
		repository: repo,
		passSvc:    passSvc,
	}
}

// Create a new user
func (u *usecase) Create(user *User) (*User, error) {
	return nil, nil
}

// GetByID retrieves a user from repo by ID
func (u *usecase) GetByID(id uint) (*User, error) {
	return nil, nil
}

func (u *usecase) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (u *usecase) GetAll() ([]*User, error) {
	return nil, nil
}

func (u *usecase) Update(user *User) (*User, error) {
	return nil, nil
}

func (u *usecase) Delete(user *User) error {
	return nil
}

func (u *usecase) Login(email, password string) (*User, error) {
	return nil, nil
}

func (u *usecase) ChangePassword(id uint, oldPassword, newPassword string) (*User, error) {
	return nil, nil
}
