package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewAdd(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *AddOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &AddOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &AddOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewAdd(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}
func Test_Add_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case",
			want:    &State{Number: 105},
			initial: &State{Number: 100},
			input:   Input{Operation: ADD, Number: 5},
		},
		{
			desc:    "happy case inf positif",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: 100},
			input:   Input{Operation: ADD, Number: math.Inf(1)},
		},
		{
			desc:    "happy case inf negative",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: 100},
			input:   Input{Operation: ADD, Number: math.Inf(-1)},
		},
		{
			desc:    "happy case inf positif initial",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: ADD, Number: 100},
		},
		{
			desc:    "happy case inf negative initial",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: ADD, Number: 100},
		},
		{
			desc:    "happy case nan initial",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: ADD, Number: 100},
		},
		{
			desc:    "happy case nan initial negative",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: ADD, Number: -100},
		},
		{
			desc:    "happy case inf add",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: ADD, Number: math.Inf(-1)},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &AddOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(op.state.Number) {
					t.Error("not NaN")
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewSubtract(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *SubstractOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &SubstractOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &SubstractOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewSubstract(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}
func Test_Subtract_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case",
			want:    &State{Number: 95},
			initial: &State{Number: 100},
			input:   Input{Operation: SUBSTRACT, Number: 5},
		},
		{
			desc:    "happy case inf positif",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: 100},
			input:   Input{Operation: SUBSTRACT, Number: math.Inf(1)},
		},
		{
			desc:    "happy case inf negative",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: 100},
			input:   Input{Operation: SUBSTRACT, Number: math.Inf(-1)},
		},
		{
			desc:    "happy case inf add",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: SUBSTRACT, Number: math.Inf(1)},
		},
		{
			desc:    "happy case inf add positive",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: SUBSTRACT, Number: math.Inf(-1)},
		},
		{
			desc:    "happy case inf positif initial",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: SUBSTRACT, Number: 100},
		},
		{
			desc:    "happy case inf negative initial",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: SUBSTRACT, Number: 100},
		},
		{
			desc:    "happy case nan initial",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: SUBSTRACT, Number: 100},
		},
		{
			desc:    "happy case nan initial negative",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: SUBSTRACT, Number: -100},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &SubstractOps{state: tC.initial}
			op.Do(tC.input)

			if math.IsNaN(tC.want.Number) {
				if !math.IsNaN(op.state.Number) {
					t.Errorf("not NaN: %v", op.state.Number)
				}
				return
			}

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewMultiply(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *MultiplyOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &MultiplyOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case",
			want:  &MultiplyOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewMultiply(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}
func Test_Multiply_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case",
			want:    &State{Number: 500},
			initial: &State{Number: 100},
			input:   Input{Operation: MULTIPLY, Number: 5},
		},
		{
			desc:    "happy case max",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: 100},
			input:   Input{Operation: MULTIPLY, Number: math.MaxFloat64},
		},
		{
			desc:    "happy case max",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: -100},
			input:   Input{Operation: MULTIPLY, Number: math.MaxFloat64},
		},
		{
			desc:    "happy case inf * 0",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: MULTIPLY, Number: 0},
		},
		{
			desc:    "happy case inf positif initial",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: MULTIPLY, Number: 100},
		},
		{
			desc:    "happy case inf negative initial",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: MULTIPLY, Number: 100},
		},
		{
			desc:    "happy case nan initial",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: MULTIPLY, Number: 100},
		},
		{
			desc:    "happy case nan initial negative",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: MULTIPLY, Number: -100},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &MultiplyOps{state: tC.initial}
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

func Test_NewDivide(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *DivideOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &DivideOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case large number",
			want:  &DivideOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewDivide(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}
func Test_Divide_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case",
			want:    &State{Number: 20},
			initial: &State{Number: 100},
			input:   Input{Operation: DIVIDE, Number: 5},
		},
		{
			desc:    "happy case",
			want:    &State{Number: 0},
			initial: &State{Number: 0},
			input:   Input{Operation: DIVIDE, Number: 100},
		},
		{
			desc:    "happy case divide 0",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: 100},
			input:   Input{Operation: DIVIDE, Number: 0},
		},
		{
			desc:    "happy case divide 0 negative",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: -100},
			input:   Input{Operation: DIVIDE, Number: 0},
		},
		{
			desc:    "happy case divide large num",
			want:    &State{Number: 0.0},
			initial: &State{Number: 1},
			input:   Input{Operation: DIVIDE, Number: math.Inf(1)},
		},
		{
			desc:    "happy case inf positif initial",
			want:    &State{Number: math.Inf(1)},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: DIVIDE, Number: 100},
		},
		{
			desc:    "happy case inf negative initial",
			want:    &State{Number: math.Inf(-1)},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: DIVIDE, Number: 100},
		},
		{
			desc:    "happy case nan initial",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: DIVIDE, Number: 100},
		},
		{
			desc:    "happy case nan initial negative",
			want:    &State{Number: math.NaN()},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: DIVIDE, Number: -100},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &DivideOps{state: tC.initial}
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

func Test_NewCancel(t *testing.T) {
	testCases := []struct {
		desc  string
		want  *CancelOps
		input *State
	}{
		{
			desc:  "happy case",
			want:  &CancelOps{state: &State{}},
			input: &State{},
		},
		{
			desc:  "happy case large number",
			want:  &CancelOps{state: &State{Number: 1}},
			input: &State{Number: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewCancel(tC.input)

			assert.Equal(t, tC.want, op)
		})
	}
}
func Test_Cancel_Do(t *testing.T) {
	testCases := []struct {
		desc    string
		want    *State
		initial *State
		input   Input
	}{
		{
			desc:    "happy case",
			want:    &State{Number: 0},
			initial: &State{Number: 100},
			input:   Input{Operation: CANCEL},
		},
		{
			desc:    "happy case negative",
			want:    &State{Number: 0},
			initial: &State{Number: -100},
			input:   Input{Operation: CANCEL},
		},
		{
			desc:    "happy case +inf",
			want:    &State{Number: 0},
			initial: &State{Number: math.Inf(1)},
			input:   Input{Operation: CANCEL},
		},
		{
			desc:    "happy case divide -inf",
			want:    &State{Number: 0},
			initial: &State{Number: math.Inf(-1)},
			input:   Input{Operation: CANCEL},
		},
		{
			desc:    "happy case NaN",
			want:    &State{Number: 0},
			initial: &State{Number: math.NaN()},
			input:   Input{Operation: CANCEL},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := &CancelOps{state: tC.initial}
			op.Do(tC.input)

			assert.Equal(t, tC.want, tC.initial)
		})
	}
}

func Test_NewExit(t *testing.T) {
	testCases := []struct {
		desc string
		want *ExitOps
	}{
		{
			desc: "happy case",
			want: &ExitOps{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := NewExit()

			assert.Equal(t, tC.want, op)
		})
	}
}
