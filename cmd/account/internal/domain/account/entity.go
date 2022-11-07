package account

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID          uuid.UUID
	Email       Email
	Password    Password
	Active      bool
	LastLoginAt time.Time
}

// New creates a new account entity.
func New() Entity {
	return Entity{
		ID: uuid.New(),
	}
}

func (e *Entity) IsActive() bool {
	return e.Active
}

// Activate sets the account entity active field to true.
func (e *Entity) Activate() {
	e.Active = true
}

// Deactivate sets the account entity active field to true.
func (e *Entity) Deactivate() {
	e.Active = false
}

func (e *Entity) LoginStamp() {
	e.LastLoginAt = time.Now()
}

// UpdateEmail checks the validity of the email address
// and updates the account entity email field value.
func (e *Entity) UpdateEmail(email string) error {
	newEmail := Email(email)
	if ok, err := newEmail.IsValid(); !ok {
		return err
	}

	e.Email = newEmail
	return nil
}

// UpdatePassword checks the validity of the password
// and updates the account entity password field value.
func (e *Entity) UpdatePassword(password string) error {
	newPassword := Password(password)
	if ok, err := newPassword.IsValid(); !ok {
		return err
	}

	e.Password = newPassword
	return nil
}
