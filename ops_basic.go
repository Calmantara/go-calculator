package main

import "os"

type (
	// basic operation implementation
	AddOps struct {
		state *State
	}

	SubstractOps struct {
		state *State
	}

	MultiplyOps struct {
		state *State
	}

	DivideOps struct {
		state *State
	}

	CancelOps struct {
		state *State
	}

	ExitOps struct{}
)

// Add
func NewAdd(state *State) Operation {
	return &AddOps{
		state: state,
	}
}

func (o *AddOps) Do(in Input) {
	o.state.Number += in.Number
}

// Substract
func NewSubstract(state *State) Operation {
	return &SubstractOps{
		state: state,
	}
}

func (o *SubstractOps) Do(in Input) {
	o.state.Number -= in.Number
}

// Multiply
func NewMultiply(state *State) Operation {
	return &MultiplyOps{
		state: state,
	}
}

func (o *MultiplyOps) Do(in Input) {
	o.state.Number *= in.Number
}

// Divide
func NewDivide(state *State) Operation {
	return &DivideOps{
		state: state,
	}
}

func (o *DivideOps) Do(in Input) {
	o.state.Number /= in.Number
}

// Cancel
func NewCancel(state *State) Operation {
	return &CancelOps{
		state: state,
	}
}

func (o *CancelOps) Do(in Input) {
	o.state.Number = 0
}

// exit
func NewExit() Operation {
	return &ExitOps{}
}

func (o *ExitOps) Do(in Input) {
	os.Exit(0)
}
