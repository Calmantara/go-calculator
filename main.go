package main

import (
	"fmt"
	"strings"
)

func setup() (*State, OperationHistory, map[OperationType]Operation) {
	state, opsHistory := &State{Number: 0}, NewOperationHistory()
	ops := map[OperationType]Operation{
		ADD:       NewAdd(state),
		SUBSTRACT: NewSubstract(state),
		MULTIPLY:  NewMultiply(state),
		DIVIDE:    NewDivide(state),
		CANCEL:    NewCancel(state),
		ABS:       NewAbs(state),
		NEG:       NewNeg(state),
		SQRT:      NewSqrt(state),
		SQR:       NewSqr(state),
		CUBERT:    NewCubert(state),
		CUBE:      NewCube(state),
		EXIT:      NewExit(),
	}
	repeatOps := NewRepeat(state, opsHistory, ops)
	ops[REPEAT] = repeatOps
	return state, opsHistory, ops
}

func exec(input Input, opsHistory OperationHistory, ops map[OperationType]Operation) {
	// validate operation input from user
	lowerInput := strings.ToLower(input.Operation.String())
	op, ok := ops[OperationType(lowerInput)]
	if !ok {
		// continue if invalid input
		fmt.Println("invalid operation type")
		return
	}
	input.Operation = OperationType(lowerInput)
	// execute operation and store to state
	op.Do(input)
	// add to history
	opsHistory.Store(input)
}

func main() {
	state, opsHistory, ops := setup()
	for {
		input := Input{}

		//binding input
		fmt.Scanln(&input.Operation, &input.Number)
		// exec command
		exec(input, opsHistory, ops)
		// print state
		fmt.Printf("%.1f\n", state.Number)
	}
}
