package state_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	state "github.com/IT4smart/sherlock-state/v2"
)

func TestDraftState(t *testing.T) {
	// Create a new object in draft state
	obj := &state.Object{
		ActualState:  &state.DraftState{},
		DesiredState: &state.DraftState{},
	}

	assert.Equal(t, "draft", obj.ActualState.GetState())

}

func TestDraft_SetState_Warm(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.DraftState{},
		DesiredState: &state.DraftState{},
	}

	obj.ActualState.SetState("warm", obj)

	assert.Equal(t, "warm", obj.ActualState.GetState())
}

func TestWarmState(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.WarmState{},
		DesiredState: &state.WarmState{},
	}

	assert.Equal(t, "warm", obj.ActualState.GetState())
}

func TestWarm_SetState_Online(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.WarmState{},
		DesiredState: &state.WarmState{},
	}

	obj.ActualState.SetState("online", obj)

	assert.Equal(t, "online", obj.ActualState.GetState())
}

func TestWarm_SetState_Offline(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.WarmState{},
		DesiredState: &state.WarmState{},
	}

	obj.ActualState.SetState("offline", obj)

	assert.Equal(t, "offline", obj.ActualState.GetState())
}

func TestOnlineState(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.OnlineState{},
		DesiredState: &state.OnlineState{},
	}

	assert.Equal(t, "online", obj.ActualState.GetState())
}

func TestOnline_SetState_Offline(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.OnlineState{},
		DesiredState: &state.OnlineState{},
	}

	obj.ActualState.SetState("offline", obj)

	assert.Equal(t, "offline", obj.ActualState.GetState())
}

func TestOnline_SetState_Warm(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.OnlineState{},
		DesiredState: &state.OnlineState{},
	}

	obj.ActualState.SetState("warm", obj)

	assert.Equal(t, "warm", obj.ActualState.GetState())
}

func TestOfflineState(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.OfflineState{},
		DesiredState: &state.OfflineState{},
	}

	assert.Equal(t, "offline", obj.ActualState.GetState())
}

func TestOffline_SetState_Warm(t *testing.T) {
	obj := &state.Object{
		ActualState:  &state.OfflineState{},
		DesiredState: &state.OfflineState{},
	}

	obj.ActualState.SetState("warm", obj)

	assert.Equal(t, "warm", obj.ActualState.GetState())
}
