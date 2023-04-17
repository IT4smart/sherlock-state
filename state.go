package state

import (
	"errors"
)

// Object defines a struct that holds actual and desired states
type State struct {
	ActualState  string `json:"actual"`
	DesiredState string `json:"desired"`
}

const (
	Draft   = "draft"
	Warm    = "warm"
	Online  = "online"
	Offline = "offline"
)

// SetActualState sets the actual state of the object
func (s *State) SetActualState(state string) error {
	if !isValidState(state) {
		return errors.New("invalid state")
	}
	s.ActualState = state
	return nil
}

// SetDesiredState sets the desired state of the object
func (s *State) SetDesiredState(state string) error {
	if !isValidState(state) {
		return errors.New("invalid state")
	}
	s.DesiredState = state
	return nil
}

// GetActualState gets the actual state of the object
func (s *State) GetActualState() string {
	return s.ActualState
}

// GetDesiredState gets the desired state of the object
func (s *State) GetDesiredState() string {
	return s.DesiredState
}

// isValidState returns true if the provided state is valid, and false otherwise
func isValidState(state string) bool {
	switch state {
	case Draft, Warm, Online, Offline:
		return true
	default:
		return false
	}
}
