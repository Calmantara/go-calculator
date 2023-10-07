package main

import "os"

type (
	// basic operation implementation

	// addition basic ops
	// the operation implemented is state + n
	AddOps struct {
		state *State
	}

	// subtract basic ops
	// the operation implemented is state - n
	SubstractOps struct {
		state *State
	}

	// multiply basic ops
	// the operation implemented is state * n
	MultiplyOps struct {
		state *State
	}

	// divide basic ops
	// the operation implemented is state / n
	// would be return Inf if divided by 0
	DivideOps struct {
		state *State
	}

	// cancel basic ops
	// the operation implemented is
	// forcing state number to be 0
	CancelOps struct {
		state *State
	}

	// exit operation to exit the program
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
