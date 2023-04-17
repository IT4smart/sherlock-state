package state

import (
	"fmt"
)

// State defines the state of an object
type State interface {
	SetState(state string, obj *Object)
	GetState() string
}

// Object defines a struct that holds actual and desired states
type Object struct {
	ActualState  *ActualState
	DesiredState *DesiredState
}

// ActualState defines the actual state of the object
type ActualState struct {
	State string `json:"actual"`
}

func (as *ActualState) SetState(state string, obj *Object) {
	switch state {
	case "draft":
		obj.ActualState = &ActualState{State: "draft"}
	case "warm":
		obj.ActualState = &ActualState{State: "warm"}
	case "online":
		obj.ActualState = &ActualState{State: "online"}
	case "offline":
		obj.ActualState = &ActualState{State: "offline"}
	default:
		fmt.Printf("Error: cannot transition to state %s from %s\n", state, as.State)
	}
}

func (as *ActualState) GetState() string {
	return as.State
}

// DesiredState defines the desired state of the object
type DesiredState struct {
	State string `json:"desired"`
}

func (ds *DesiredState) SetState(state string, obj *Object) {
	obj.DesiredState.State = state
}

func (ds *DesiredState) GetState() string {
	return ds.State
}
