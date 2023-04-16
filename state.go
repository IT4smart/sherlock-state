package state

import (
	"fmt"
	"time"
)

// State defines the state of an object
type State interface {
	SetState(state string, obj *Object)
	GetState() string
}

// Object defines a struct that holds actual and desired states
type Object struct {
	ActualState  State
	DesiredState State
}

// DraftState defines a state where the object is in draft mode
type DraftState struct{}

func (ds *DraftState) SetState(state string, obj *Object) {
	switch state {
	case "warm":
		fmt.Println("Object is warming up...")
		time.Sleep(3 * time.Second)
		obj.ActualState = &WarmState{}
	default:
		fmt.Printf("Error: cannot transition to state %s from draft\n", state)
	}
}

func (ds *DraftState) GetState() string {
	return "draft"
}

// WarmState defines a state where the object is warming up
type WarmState struct{}

func (ws *WarmState) SetState(state string, obj *Object) {
	switch state {
	case "online":
		fmt.Println("Object is now online!")
		obj.ActualState = &OnlineState{}
	case "offline":
		obj.ActualState = &OfflineState{}
	default:
		fmt.Printf("Error: cannot transition to state %s from warm\n", state)
	}
}

func (ws *WarmState) GetState() string {
	return "warm"
}

// OnlineState defines a state where the object is online
type OnlineState struct{}

func (os *OnlineState) SetState(state string, obj *Object) {
	switch state {
	case "offline":
		fmt.Println("Object is going offline...")
		time.Sleep(3 * time.Second)
		obj.ActualState = &OfflineState{}
	case "warm":
		fmt.Println("Object is going to warm state...")
		obj.ActualState = &WarmState{}
	default:
		fmt.Printf("Error: cannot transition to state %s from online\n", state)
	}
}

func (os *OnlineState) GetState() string {
	return "online"
}

// OfflineState defines a state where the object is offline
type OfflineState struct{}

func (os *OfflineState) SetState(state string, obj *Object) {
	switch state {
	case "warm":
		fmt.Println("Object is going to warm state...")
		time.Sleep(3 * time.Second)
		obj.ActualState = &WarmState{}
	default:
		fmt.Printf("Error: cannot transition to state %s from offline\n", state)
	}
}

func (os *OfflineState) GetState() string {
	return "offline"
}
