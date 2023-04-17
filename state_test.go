package state_test

import (
	"testing"

	state "github.com/IT4smart/sherlock-state"
)

func TestSetActualState(t *testing.T) {
	state := &state.State{}
	err := state.SetActualState("draft")
	if err != nil {
		t.Errorf("Error setting actual state: %v", err)
	}
	if state.ActualState != "draft" {
		t.Errorf("Actual state not set correctly: expected %s, but got %s", "draft", state.ActualState)
	}

	// test setting invalid state
	err = state.SetActualState("invalid")
	if err == nil {
		t.Errorf("Expected error setting invalid actual state, but got nil")
	}
}

func TestSetDesiredState(t *testing.T) {
	validState := "warm"
	invalidState := "invalid"

	s := state.State{}

	// Test setting a valid state
	err := s.SetDesiredState(validState)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	if s.DesiredState != validState {
		t.Errorf("Expected desired state to be %s but got %s", validState, s.DesiredState)
	}

	// Test setting an invalid state
	err = s.SetDesiredState(invalidState)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
	if s.DesiredState == invalidState {
		t.Errorf("Expected desired state to remain unchanged but got %s", s.DesiredState)
	}
}

func TestGetActualState(t *testing.T) {
	s := &state.State{ActualState: state.Online}
	if astate := s.GetActualState(); astate != state.Online {
		t.Errorf("expected %q, but got %q", state.Online, astate)
	}
}

func TestGetDesiredState(t *testing.T) {
	s := &state.State{DesiredState: state.Online}
	if astate := s.GetDesiredState(); astate != state.Online {
		t.Errorf("expected %q, but got %q", state.Online, astate)
	}
}
