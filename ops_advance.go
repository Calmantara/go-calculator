package main

import "math"

type (
	// advance operation

	// absolute operation
	// would be changed
	// negative number into positive
	// included Inf number
	AbsOps struct {
		state *State
	}

	// negation operation
	// would be change
	// positive number into negative one
	// vice versa
	NegOps struct {
		state *State
	}

	// square root operation
	// would be return NaN
	// if state number is negative
	SqrtOps struct {
		state *State
	}

	// square function
	// to operate
	// state number * state number
	SqrOps struct {
		state *State
	}

	// cube root to operate
	// cube root of state number
	CubertOps struct {
		state *State
	}

	// cube opration for
	// state number
	CubeOps struct {
		state *State
	}

	// this operation would be
	// replay all event recorded in history
	// and recursively operate if any repeat in
	// recorded history
	RepeatOps struct {
		ops        map[OperationType]Operation
		opsHistory OperationHistory
		state      *State
	}
)

// Abs
func NewAbs(state *State) Operation {
	return &AbsOps{
		state: state,
	}
}

func (o *AbsOps) Do(in Input) {
	o.state.Number = math.Abs(o.state.Number)
}

// Negative
func NewNeg(state *State) Operation {
	return &NegOps{
		state: state,
	}
}

func (o *NegOps) Do(in Input) {
	o.state.Number *= -1
}

// Square Root
func NewSqrt(state *State) Operation {
	return &SqrtOps{
		state: state,
	}
}

func (o *SqrtOps) Do(in Input) {
	o.state.Number = math.Sqrt(o.state.Number)
}

// Square
func NewSqr(state *State) Operation {
	return &SqrOps{
		state: state,
	}
}

func (o *SqrOps) Do(in Input) {
	o.state.Number *= o.state.Number
}

// Cube Root
func NewCubert(state *State) Operation {
	return &CubertOps{
		state: state,
	}
}

func (o *CubertOps) Do(in Input) {
	o.state.Number = math.Cbrt(o.state.Number)
}

// Cube Root
func NewCube(state *State) Operation {
	return &CubeOps{
		state: state,
	}
}

func (o *CubeOps) Do(in Input) {
	o.state.Number = o.state.Number * o.state.Number * o.state.Number
}

// Repeat
func NewRepeat(state *State, opsHistory OperationHistory, ops map[OperationType]Operation) Operation {
	return &RepeatOps{
		ops:        ops,
		opsHistory: opsHistory,
		state:      state,
	}
}

func (o *RepeatOps) Do(in Input) {
	// get from history event
	histories := o.opsHistory.GetLastRecord(in.Number, in.Offset)
	for _, history := range histories {
		// check stop recursive
		if history.Offset == 0 &&
			history.Operation == REPEAT {
			continue
		}
		// if history valid, exec opeartion
		operations := o.ops[history.Operation]
		operations.Do(history)
	}
}
