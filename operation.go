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

type OperationType string

func (o OperationType) String() string {
	return string(o)
}

// define model and interface
type (
	Operation interface {
		Do(in Input)
	}

	OperationHistory interface {
		Store(in Input)
		GetLastRecord(n float64, offset int) []Input
	}
	// model
	OperationHistoryImpl struct {
		inArr []Input
	}

	State struct {
		Number float64
	}

	Input struct {
		Operation OperationType
		Number    float64
		Offset    int
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
