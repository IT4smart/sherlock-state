package state_test

import (
	"testing"

	state "github.com/IT4smart/sherlock-state/v2"
	"github.com/stretchr/testify/assert"
)

func TestActualState_SetState(t *testing.T) {
	// create a new object with an actual state of "offline"
	obj := &state.Object{
		ActualState: &state.ActualState{State: "offline"},
	}

	// set the state to "online"
	obj.ActualState.SetState("online", obj)

	// check that the actual state is now "online"
	assert.Equal(t, "online", obj.ActualState.State)

	// set the state to "offline" again
	obj.ActualState.SetState("offline", obj)

	// check that the actual state is now "offline" again
	assert.Equal(t, "offline", obj.ActualState.State)

	obj.ActualState.SetState("draft", obj)

	assert.Equal(t, "draft", obj.ActualState.State)
}

func TestDesiredState_SetState(t *testing.T) {
	// Initialize the actual state and the desired state
	actualState := &state.ActualState{State: "offline"}
	desiredState := &state.DesiredState{State: "online"}

	// Initialize the object with the actual and desired states
	obj := &state.Object{ActualState: actualState, DesiredState: desiredState}

	// Call the SetState function to change the desired state to "warm"
	obj.DesiredState.SetState("warm", obj)

	// Check if the desired state is "warm"
	if obj.DesiredState.GetState() != "warm" {
		t.Errorf("SetState failed, expected state: %s, got: %s", "warm", obj.DesiredState.GetState())
	}

	// Call the SetState function to change the desired state to "online"
	obj.DesiredState.SetState("online", obj)

	// Check if the desired state is "online"
	if obj.DesiredState.GetState() != "online" {
		t.Errorf("SetState failed, expected state: %s, got: %s", "online", obj.DesiredState.GetState())
	}
}
