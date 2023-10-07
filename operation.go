package main

import "math"

const (
	ADD       OperationType = "add"
	SUBSTRACT OperationType = "subtract"
	MULTIPLY  OperationType = "multiply"
	DIVIDE    OperationType = "divide"
	CANCEL    OperationType = "cancel"
	ABS       OperationType = "abs"
	NEG       OperationType = "neg"
	SQRT      OperationType = "sqrt"
	SQR       OperationType = "sqr"
	CUBERT    OperationType = "cubert"
	CUBE      OperationType = "cube"
	EXIT      OperationType = "exit"
	REPEAT    OperationType = "repeat"
)

// operation custom type
// so it would be more
// readable if you want
// to add more operations
type OperationType string

func (o OperationType) String() string {
	return string(o)
}

// define model and interface
type (

	// generic oeparion interface
	// use and / or composite this interface
	// to make new operation implementation
	// for this calculator
	Operation interface {
		Do(in Input)
	}

	// Operation history interface
	// to store and get last n record
	// this should be imutable record of event
	// in the implementation side
	OperationHistory interface {
		Store(in Input)
		GetLastRecord(n float64, offset int) []Input
	}

	// operation history model
	// implementation for event store
	// for all user input
	// would be appended only if
	// input is valid with predefined
	// operation types
	OperationHistoryImpl struct {
		inArr []Input
	}

	// state number or latest number
	// of calculator from all inputs
	State struct {
		Number float64
	}

	// input model that represent user input
	Input struct {
		// operation name filled
		// by user, predefined in constant
		Operation OperationType

		// input number from user
		// would be input n
		// from the operations function
		Number float64

		// index of recorded event
		Offset int
	}
)

func NewOperationHistory() OperationHistory {
	return &OperationHistoryImpl{inArr: []Input{}}
}

func (oh *OperationHistoryImpl) Store(in Input) {
	// set offset to record event
	in.Offset = len(oh.inArr)
	oh.inArr = append(oh.inArr, in)
	return
}

func (oh *OperationHistoryImpl) GetLastRecord(n float64, offset int) []Input {
	in := int(math.Ceil(n))
	if offset == 0 {
		offset = len(oh.inArr)
	}
	if offset < in {
		return oh.inArr[:offset]
	}
	return oh.inArr[offset-in : offset]
}
