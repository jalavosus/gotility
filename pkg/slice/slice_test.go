package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/gotility/pkg/slice"
)

var testSlice = []int64{
	6, 7, 8, 3, 2222,
}

type argsCmpFunc[T any] struct {
	data []T
	fn   slice.CmpFunc[T]
}

type argsValue[T any] struct {
	data []T
	val  T
}

type argsType[T any] interface {
	argsCmpFunc[T] | argsValue[T]
}

type wantType[T any] interface {
	[]T | bool
}

type cmpTestCase[T any, ArgsType argsType[T], WantType wantType[T]] struct {
	name string
	args ArgsType
	want WantType
}

type elemTestArgs[T any] struct {
	data []T
	fn   slice.CmpFunc[T]
}

type elemTestCaseSingle[T any] struct {
	name string
	args elemTestArgs[T]
	want T
}

type elemTestCaseMulti[T any] struct {
	name string
	args elemTestArgs[T]
	want []T
}

func TestFirst(t *testing.T) {
	tests := []elemTestCaseSingle[int64]{
		{
			name: "testCase",
			args: elemTestArgs[int64]{
				data: testSlice,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.First(tt.args.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLast(t *testing.T) {
	tests := []elemTestCaseSingle[int64]{
		{
			name: "testCase",
			args: elemTestArgs[int64]{
				data: testSlice,
			},
			want: 2222,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.Last(tt.args.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHead(t *testing.T) {
	tests := []elemTestCaseMulti[int64]{
		{
			name: "testCase",
			args: elemTestArgs[int64]{
				data: testSlice,
			},
			want: []int64{6, 7, 8, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				got := slice.Head(tt.args.data)
				assert.Equal(t, tt.want, got)
			})
		})
	}
}

func TestTail(t *testing.T) {
	tests := []elemTestCaseMulti[int64]{
		{
			name: "testCase",
			args: elemTestArgs[int64]{
				data: testSlice,
			},
			want: []int64{7, 8, 3, 2222},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				got := slice.Tail(tt.args.data)
				assert.Equal(t, tt.want, got)
			})
		})
	}
}

func TestContains(t *testing.T) {
	tests := []cmpTestCase[int64, argsCmpFunc[int64], bool]{
		{
			name: "contains=true",
			args: argsCmpFunc[int64]{
				data: []int64{1, 2, 3, 4, 5},
				fn:   func(v int64) bool { return v == 2 },
			},
			want: true,
		},
		{
			name: "contains=false",
			args: argsCmpFunc[int64]{
				data: []int64{1, 2, 3, 4, 5},
				fn:   func(v int64) bool { return v == 6 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slice.Contains(tt.args.data, tt.args.fn), "Contains(%v, %v)", tt.args.data, tt.args.fn)
		})
	}
}

func TestContainsValue(t *testing.T) {
	tests := []cmpTestCase[int64, argsValue[int64], bool]{
		{
			name: "contains=true",
			args: argsValue[int64]{
				data: []int64{1, 2, 3, 4, 5},
				val:  2,
			},
			want: true,
		},
		{
			name: "contains=false",
			args: argsValue[int64]{
				data: []int64{1, 2, 3, 4, 5},
				val:  6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slice.ContainsValue(tt.args.data, tt.args.val), "ContainsValue(%v, %v)", tt.args.data, tt.args.val)
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []cmpTestCase[int64, argsCmpFunc[int64], []int64]{
		{
			name: "filter evens",
			args: argsCmpFunc[int64]{
				data: []int64{1, 2, 3, 4, 5},
				fn:   func(v int64) bool { return v%2 == 0 },
			},
			want: []int64{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slice.Filter(tt.args.data, tt.args.fn), "Filter(%v, %v)", tt.args.data, tt.args.fn)
		})
	}
}

func TestFilterValue(t *testing.T) {
	tests := []cmpTestCase[int64, argsValue[int64], []int64]{
		{
			name: "filter fives",
			args: argsValue[int64]{
				data: []int64{5, 1, 2, 5, 3, 4, 5},
				val:  5,
			},
			want: []int64{5, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slice.FilterValue(tt.args.data, tt.args.val), "FilterValue(%v, %v)", tt.args.data, tt.args.val)
		})
	}
}

type findTestCase[T any, ArgsType argsType[T]] struct {
	name       string
	args       ArgsType
	want       T
	wantExists bool
}

func TestFind(t *testing.T) {
	tests := []findTestCase[int64, argsCmpFunc[int64]]{
		{
			name: "exists=true",
			args: argsCmpFunc[int64]{
				data: []int64{44, 22, 3, 7, 42, 5},
				fn:   func(v int64) bool { return v == 5 },
			},
			want:       5,
			wantExists: true,
		},
		{
			name: "exists=false",
			args: argsCmpFunc[int64]{
				data: []int64{44, 22, 3, 7, 42, 5},
				fn:   func(v int64) bool { return v == 6 },
			},
			want:       0,
			wantExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := slice.Find(tt.args.data, tt.args.fn)
			assert.Equalf(t, tt.want, got, "Find(%v, %v)", tt.args.data, tt.args.fn)
			assert.Equalf(t, tt.wantExists, got1, "Find(%v, %v)", tt.args.data, tt.args.fn)
		})
	}
}

func TestFindValue(t *testing.T) {
	tests := []findTestCase[int64, argsValue[int64]]{
		{
			name: "exists=true",
			args: argsValue[int64]{
				data: []int64{44, 22, 3, 7, 42, 5},
				val:  5,
			},
			want:       5,
			wantExists: true,
		},
		{
			name: "exists=false",
			args: argsValue[int64]{
				data: []int64{44, 22, 3, 7, 42, 5},
				val:  6,
			},
			want:       0,
			wantExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := slice.FindValue(tt.args.data, tt.args.val)
			assert.Equalf(t, tt.want, got, "FindValue(%v, %v)", tt.args.data, tt.args.val)
			assert.Equalf(t, tt.wantExists, got1, "FindValue(%v, %v)", tt.args.data, tt.args.val)
		})
	}
}
