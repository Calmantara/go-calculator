package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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
