package pointers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/gotility/pkg/pointers"
)

type args[T any] struct {
	val T
}

type testCase[T, U any] struct {
	name string
	args args[T]
	want U
}

func TestFromPointer(t *testing.T) {
	var (
		n1 int64 = 42
	)

	tests := []testCase[*int64, int64]{
		{
			name: "not-nil",
			args: args[*int64]{
				val: &n1,
			},
			want: n1,
		},
		{
			name: "nil",
			args: args[*int64]{
				val: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pointers.FromPointer(tt.args.val), "FromPointer(%v)", tt.args.val)
		})
	}
}

func TestToPointer(t *testing.T) {
	var (
		n1 int64 = 42
	)

	tests := []testCase[int64, *int64]{
		{
			name: "42",
			args: args[int64]{
				val: n1,
			},
			want: &n1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pointers.ToPointer(tt.args.val), "ToPointer(%v)", tt.args.val)
		})
	}
}
