package main

import "fmt"

type ServerState int // Create a new type called ServerState

// Define the possible states as constants
// using iota to automatically assign values
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Create a map to associate each state with its string representation
var stateNames = map[ServerState]string{
	StateIdle: 		"idle",
	StateConnected: "connected",
	StateError:    	"error",
	StateRetrying:  "retrying",
}

// Implement the Stringer interface for ServerState
// This allows us to print the state as a string
// when we use fmt.Println or similar functions
// because String() is called automatically
func (ss ServerState) String() string { // Method : Implement the Stringer interface
	return stateNames[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

}

// transition function takes a ServerState and returns the next state
func transition(s ServerState) ServerState {
	switch s {
		case StateIdle:
			return StateConnected
		case StateConnected, StateRetrying:
			return StateIdle
		case StateError:
			return StateError
		default: 
			panic(fmt.Errorf("unknow state %s", s))
	}
}