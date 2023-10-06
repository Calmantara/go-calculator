package main

import (
	"fmt"
	"strings"
)

func main() {
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

	for {
		input := Input{}

		//binding input
		fmt.Scanln(&input.Operation, &input.Number)
		// validate operation input from user
		lowerInput := strings.ToLower(input.Operation.String())
		op, ok := ops[OperationType(lowerInput)]
		if !ok {
			// continue if invalid input
			fmt.Println("invalid operation type")
			continue
		}
		// execute operation and store to state
		op.Do(input)
		// add to history
		opsHistory.Store(input)
		// print state
		fmt.Printf("%.1f\n", state.Number)
	}
}
