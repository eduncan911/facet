package domain

import (
	"github.com/eduncan911/es"
)

// UserNameChanged represents a user's username has been changed
type UserNameChanged struct {
	EventID     EventID `json:"eventId,omitempty"`
	UserID      UserID  `json:"userId,omitempty"`
	OldUserName string  `json:"oldUserName,omitempty"`
	NewUserName string  `json:"newUserName,omitempty"`
}

func NewUserNameChanged(event EventID) *UserNameChanged {
	return &UserNameChanged{}
}

func (e *UserNameChanged) Meta() *es.Info { return i("UserNameChanged", e.EventID) }
