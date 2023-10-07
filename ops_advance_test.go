package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_NewAbs(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *AbsOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &AbsOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &AbsOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewAbs(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}

func Test_Abs_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case",
			want:    &State{Number: 100},
			initial: &State{Number: -100},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case inf positif",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case nan",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: ABS},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &AbsOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(tC.initial.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewNeg(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *NegOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &NegOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &NegOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewNeg(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}

func Test_Neg_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case negative num",
			want:    &State{Number: 100},
			initial: &State{Number: -100},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case positive num",
			want:    &State{Number: -100},
			initial: &State{Number: 100},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case inf positif",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case inf negative",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case nan",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: ABS},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &NegOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(tC.initial.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewSqrt(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *SqrtOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &SqrtOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &SqrtOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewSqrt(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}

func Test_Sqrt_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case positive num",
			want:    &State{Number: 10},
			initial: &State{Number: 100},
			input:   Input{Operation: SQRT},
		},
		{
			desc:    "happy case negative num",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: -100},
			input:   Input{Operation: SQRT},
		},
		{
			desc:    "happy case inf positif",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: SQRT},
		},
		{
			desc:    "happy case nan",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: SQRT},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &SqrtOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(tC.initial.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewSqr(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *SqrOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &SqrOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &SqrOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewSqr(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}

func Test_Sqr_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case negative num",
			want:    &State{Number: 10000},
			initial: &State{Number: -100},
			input:   Input{Operation: SQR},
		},
		{
			desc:    "happy case positive num",
			want:    &State{Number: 10000},
			initial: &State{Number: 100},
			input:   Input{Operation: SQR},
		},
		{
			desc:    "happy case inf negative",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: SQR},
		},
		{
			desc:    "happy case inf positive",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: SQR},
		},
		{
			desc:    "happy case nan",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: SQR},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &SqrOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(tC.initial.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewCubert(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *CubertOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &CubertOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &CubertOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewCubert(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}

func Test_Cubert_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case positive num",
			want:    &State{Number: 10},
			initial: &State{Number: 1000},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case negative num",
			want:    &State{Number: -10},
			initial: &State{Number: -1000},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case inf negative",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case inf positive",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: ABS},
		},
		{
			desc:    "happy case nan",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: ABS},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &CubertOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(tC.initial.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewCube(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *CubeOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &CubeOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &CubeOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewCube(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}

func Test_Cube_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case negative num",
			want:    &State{Number: -1000},
			initial: &State{Number: -10},
			input:   Input{Operation: CUBE},
		},
		{
			desc:    "happy case positive num",
			want:    &State{Number: 1000},
			initial: &State{Number: 10},
			input:   Input{Operation: CUBE},
		},
		{
			desc:    "happy case inf negative",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: CUBE},
		},
		{
			desc:    "happy case inf positive",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: CUBE},
		},
		{
			desc:    "happy case nan",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: CUBE},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &CubeOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(tC.initial.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewRepeat(t *testing.T) {
	type input struct {
		state      *State
		opsHistory OperationHistory
		ops        map[OperationType]Operation
	}

	testCases := []struct {
		desc  string
		want  *RepeatOps
		input input
	}{
		{
			desc: "happy case",
			want: &RepeatOps{
				state: &State{},
				ops:   map[OperationType]Operation{},
				opsHistory: &OperationHistoryImpl{
					inArr: []Input{},
				}},
			input: input{
				state:      &State{},
				opsHistory: NewOperationHistory(),
				ops:        map[OperationType]Operation{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewRepeat(tC.input.state, tC.input.opsHistory, tC.input.ops)

			assert.Equal(t, tC.want, op)
		})
	}
}

func Test_Repeat_Do(t *testing.T) {
	type initial struct {
		state *State
		ops   map[OperationType]Operation
	}

	testCases := []struct {
		desc    string
		want    *State
		doMock  func(opsHistory *MockOperationHistory, ops *MockOperation)
		initial initial
		input   Input
	}{
		{
			desc: "happy case",
			want: &State{Number: 3},
			doMock: func(opsHistory *MockOperationHistory, ops *MockOperation) {
				opsHistory.EXPECT().GetLastRecord(gomock.Eq(float64(1)), gomock.Eq(0)).AnyTimes().MinTimes(1).Return(
					[]Input{
						{Operation: ADD, Number: 1, Offset: 0},
					},
				)
			},
			initial: initial{
				state: &State{Number: 2},
				ops:   map[OperationType]Operation{},
			},
			input: Input{Operation: REPEAT, Number: 1, Offset: 0},
		},
		{
			desc: "happy case repeat 4",
			want: &State{Number: 6},
			doMock: func(opsHistory *MockOperationHistory, ops *MockOperation) {
				opsHistory.EXPECT().GetLastRecord(gomock.Eq(float64(4)), gomock.Eq(0)).AnyTimes().MinTimes(1).Return(
					[]Input{
						{Operation: ADD, Number: 1, Offset: 0},
						{Operation: ADD, Number: 1, Offset: 1},
						{Operation: ADD, Number: 1, Offset: 2},
						{Operation: ADD, Number: 1, Offset: 3},
					},
				)
			},
			initial: initial{
				state: &State{Number: 2},
				ops:   map[OperationType]Operation{},
			},
			input: Input{Operation: REPEAT, Number: 4, Offset: 0},
		},
		{
			desc: "happy case repeat on offset 0",
			want: &State{Number: 3},
			doMock: func(opsHistory *MockOperationHistory, ops *MockOperation) {
				opsHistory.EXPECT().GetLastRecord(gomock.Eq(float64(100)), gomock.Eq(0)).AnyTimes().MinTimes(1).Return(
					[]Input{
						{Operation: REPEAT, Number: 1, Offset: 0},
						{Operation: ADD, Number: 1, Offset: 1},
					},
				)
			},
			initial: initial{
				state: &State{Number: 2},
				ops:   map[OperationType]Operation{},
			},
			input: Input{Operation: REPEAT, Number: 100, Offset: 0},
		},
		{
			desc: "happy case repeat on offset 0 with 2 repeat",
			want: &State{Number: 3},
			doMock: func(opsHistory *MockOperationHistory, ops *MockOperation) {
				opsHistory.EXPECT().GetLastRecord(gomock.Eq(float64(100)), gomock.Eq(0)).AnyTimes().MinTimes(1).Return(
					[]Input{
						{Operation: REPEAT, Number: 1, Offset: 0},
						{Operation: REPEAT, Number: 1, Offset: 1},
						{Operation: ADD, Number: 1, Offset: 2},
					},
				)
				// second call
				opsHistory.EXPECT().GetLastRecord(gomock.Eq(float64(1)), gomock.Eq(1)).AnyTimes().MinTimes(1).Return(
					[]Input{
						{Operation: REPEAT, Number: 1, Offset: 0},
					},
				)
			},
			initial: initial{
				state: &State{Number: 2},
				ops:   map[OperationType]Operation{},
			},
			input: Input{Operation: REPEAT, Number: 100, Offset: 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// do mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			opsHistory := NewMockOperationHistory(ctrl)
			ops := NewMockOperation(ctrl)

			// add ops to map of ops
			tC.initial.ops[ADD] = NewAdd(tC.initial.state)
			op := &RepeatOps{
				state:      tC.initial.state,
				opsHistory: opsHistory,
				ops:        tC.initial.ops}
			tC.initial.ops[REPEAT] = op

			// do mock and exec function
			tC.doMock(opsHistory, ops)
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(tC.initial.state.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial.state)
		})
	}
}
