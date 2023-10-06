package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	testCases := []struct {
		desc string
		want string
	}{
		{
			desc: "happy case",
			want: "hello world",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			op := OperationType("hello world")
			// assert
			assert.Equal(t, tC.want, op.String())
		})
	}
}

func Test_NewOperationHistory(t *testing.T) {
	testCases := []struct {
		desc string
		want *OperationHistoryImpl
	}{
		{
			desc: "happy case",
			want: &OperationHistoryImpl{inArr: []Input{}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ohimpl := NewOperationHistory()
			// assert
			assert.Equal(t, tC.want, ohimpl)
		})
	}
}

func Test_Store(t *testing.T) {
	testCases := []struct {
		desc    string
		in      Input
		initial []Input
		want    []Input
	}{
		{
			desc: "happy case",
			in: Input{
				Operation: ADD,
				Number:    1,
			},
			want: []Input{
				{
					Operation: "add",
					Number:    1,
					Offset:    0,
				},
			},
			initial: []Input{},
		},
		{
			desc: "happy case multi exist",
			in: Input{
				Operation: SUBSTRACT,
				Number:    12,
			},
			want: []Input{
				{
					Operation: "add",
					Number:    1,
					Offset:    0,
				},
				{
					Operation: "subtract",
					Number:    12,
					Offset:    1,
				},
			},
			initial: []Input{
				{
					Operation: ADD,
					Number:    1,
					Offset:    0,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ohimpl := &OperationHistoryImpl{inArr: tC.initial}
			ohimpl.Store(tC.in)

			// assert
			assert.Equal(t, tC.want, ohimpl.inArr)
		})
	}
}

func Test_GetLastRecord(t *testing.T) {
	testCases := []struct {
		desc    string
		in      Input
		initial []Input
		want    []Input
	}{
		{
			desc: "happy case",
			in: Input{
				Operation: ADD,
				Number:    1,
				Offset:    0,
			},
			want: []Input{
				{
					Operation: "add",
					Number:    1,
					Offset:    2,
				},
			},
			initial: []Input{
				{
					Operation: ADD,
					Number:    1,
					Offset:    0,
				},
				{
					Operation: ADD,
					Number:    1,
					Offset:    1,
				},
				{
					Operation: ADD,
					Number:    1,
					Offset:    2,
				},
			},
		},
		{
			desc: "happy case offset",
			in: Input{
				Operation: ADD,
				Number:    1,
				Offset:    1,
			},
			want: []Input{
				{
					Operation: "add",
					Number:    1,
					Offset:    0,
				},
			},
			initial: []Input{
				{
					Operation: ADD,
					Number:    1,
					Offset:    0,
				},
				{
					Operation: ADD,
					Number:    1,
					Offset:    1,
				},
				{
					Operation: ADD,
					Number:    1,
					Offset:    2,
				},
			},
		},
		{
			desc: "happy case large number",
			in: Input{
				Operation: ADD,
				Number:    1000,
				Offset:    0,
			},
			want: []Input{
				{
					Operation: "add",
					Number:    1,
					Offset:    0,
				},
				{
					Operation: "add",
					Number:    1,
					Offset:    1,
				},
				{
					Operation: "add",
					Number:    1,
					Offset:    2,
				},
			},
			initial: []Input{
				{
					Operation: ADD,
					Number:    1,
					Offset:    0,
				},
				{
					Operation: ADD,
					Number:    1,
					Offset:    1,
				},
				{
					Operation: ADD,
					Number:    1,
					Offset:    2,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ohimpl := &OperationHistoryImpl{inArr: tC.initial}
			arr := ohimpl.GetLastRecord(tC.in.Number, tC.in.Offset)

			// assert
			assert.Equal(t, tC.want, arr)
		})
	}
}
