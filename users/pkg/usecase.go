package user

import (
	"strings"

	"github.com/BetuelSA/go-helpers/errors"
	pass "github.com/BetuelSA/go-helpers/password"
	"github.com/go-playground/validator/v10"
)

const (
	minPasswordLength int = 6
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
	// Verify email uniqueness
	user.Email = strings.TrimSpace(user.Email)
	_, err = u.GetByEmail(user.Email)
	if err != nil {
		return nil, errors.BadRequest.Newf("user with email %s already exists", user.Email)
	}

	// Verify password lengh
	if len(user.Password) < minPasswordLength {
		return nil, errors.BadRequest.Newf("password must have at least %d characters", minPasswordLength)
	}

	hash, err := u.passwordSvc.Hash(user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "can't obtain password hash")
	}
	user.Hash = hash
	user.Surname = strings.TrimSpace(user.Surname)
	user.Name = strings.TrimSpace(user.Name)

	validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidateErrors)
		return nil, errors.BadRequest.Wrap(validationErrors, "error during user data validation")
	}

	user, err = u.repository.Create(user)
	if err != nil {
		return nil, errors.Wrap(err, "error creating a new user")
	}

	return user, nil
}

// GetByID retrieves a user from repo by ID
func (u *usecase) GetByID(id uint) (*User, error) {
	user, err := u.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching a user")
	}

	return user, nil
}

// GetByEmail retrieves a user from repo by email address
func (u *usecase) GetByEmail(email string) (*User, error) {
	user, err := u.repository.GetByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching a user")
	}

	return user, nil
}

// GetAll retrieves every user
func (u *usecase) GetAll() ([]*User, error) {
	users, err := u.repository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "error fetching all users")
	}

	return users, nil
}

// Update modify an existing user
func (u *usecase) Update(user *User) (*User, error) {
	// Trim spaces
	user.Surname = strings.TrimSpace(user.Surname)
	user.Name = strings.TrimSpace(user.Name)
	user.Password = strings.TrimSpace(user.Password)

	formerUser := &User{}

	// Verify email uniqueness
	formerUser, err = u.GetByEmail(user.Email)
	if (err == nil) && (formerUser.ID != user.ID) {
		return nil, errors.BadRequest.Newf("user with email %s already exists", user.Email)
	}

	// Verify password

}

// Delete a user
func (u *usecase) Delete(user *User) error {
	return nil
}

// Login validates a user by email/password
func (u *usecase) Login(email, password string) (*User, error) {
	return nil, nil
}

// ChangePassword permits that a user can change her password
func (u *usecase) ChangePassword(id uint, oldPassword, newPassword string) (*User, error) {
	return nil, nil
}
