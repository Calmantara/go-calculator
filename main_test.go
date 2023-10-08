package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_setup(t *testing.T) {
	type want struct {
		state      *State
		opsHistory OperationHistory
		ops        map[OperationType]Operation
	}

	ops := map[OperationType]Operation{
		ADD:       &AddOps{state: &State{}},
		SUBSTRACT: &SubstractOps{state: &State{}},
		MULTIPLY:  &MultiplyOps{state: &State{}},
		DIVIDE:    &DivideOps{state: &State{}},
		CANCEL:    &CancelOps{state: &State{}},
		ABS:       &AbsOps{state: &State{}},
		NEG:       &NegOps{state: &State{}},
		SQRT:      &SqrtOps{state: &State{}},
		SQR:       &SqrOps{state: &State{}},
		CUBERT:    &CubertOps{state: &State{}},
		CUBE:      &CubeOps{state: &State{}},
		EXIT:      &ExitOps{},
	}
	ops[REPEAT] = &RepeatOps{
		ops:        ops,
		opsHistory: &OperationHistoryImpl{inArr: []Input{}},
		state:      &State{},
	}

	testCases := []struct {
		desc string
		want want
	}{
		{
			desc: "happy case",
			want: want{
				state:      &State{},
				opsHistory: &OperationHistoryImpl{inArr: []Input{}},
				ops:        ops,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			st, oph, op := setup()

			assert.Equal(t, tC.want.state, st)
			assert.Equal(t, tC.want.opsHistory, oph)
			assert.Equal(t, tC.want.ops, op)
		})
	}
}

func Test_exec(t *testing.T) {
	testCases := []struct {
		desc   string
		in     Input
		doMock func(op *MockOperation, oph *MockOperationHistory)
	}{
		{
			desc: "happy case",
			in:   Input{Operation: ADD, Number: 1},
			doMock: func(op *MockOperation, oph *MockOperationHistory) {
				op.EXPECT().Do(gomock.Eq(Input{Operation: "add", Number: 1})).AnyTimes().MinTimes(1)
				oph.EXPECT().Store(Input{Operation: "add", Number: 1}).AnyTimes().MinTimes(1)
			},
		},
		{
			desc: "happy case insensitive case",
			in:   Input{Operation: "ADD", Number: 2},
			doMock: func(op *MockOperation, oph *MockOperationHistory) {
				op.EXPECT().Do(gomock.Eq(Input{Operation: "add", Number: 2})).AnyTimes().MinTimes(1)
				oph.EXPECT().Store(Input{Operation: "add", Number: 2}).AnyTimes().MinTimes(1)
			},
		},
		{
			desc: "negative case not in operations",
			in:   Input{Operation: "INVALID", Number: 2},
			doMock: func(op *MockOperation, oph *MockOperationHistory) {
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			op := NewMockOperation(ctrl)
			oph := NewMockOperationHistory(ctrl)

			ops := map[OperationType]Operation{
				ADD: op,
			}
			// do mock and exec func
			tC.doMock(op, oph)
			exec(tC.in, oph, ops)
		})
	}
}
