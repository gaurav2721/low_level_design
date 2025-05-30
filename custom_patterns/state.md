package main

import "fmt"

// State interface
type State interface {
	PressPlay(p *Player)
}

// Context
type Player struct {
	state State
}

func (p *Player) SetState(s State) {
	p.state = s
}

func (p *Player) PressPlay() {
	p.state.PressPlay(p)
}

// Concrete State: Playing
type PlayingState struct{}

func (s *PlayingState) PressPlay(p *Player) {
	fmt.Println("Pausing music...")
	p.SetState(&PausedState{})
}

// Concrete State: Paused
type PausedState struct{}

func (s *PausedState) PressPlay(p *Player) {
	fmt.Println("Resuming music...")
	p.SetState(&PlayingState{})
}

func main() {
	player := &Player{}

	// Set initial state
	player.SetState(&PausedState{})

	player.PressPlay() // -> Resuming music...
	player.PressPlay() // -> Pausing music...
	player.PressPlay() // -> Resuming music...
}
