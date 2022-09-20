package maps_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/gotility/pkg/maps"
)

func newMap(t *testing.T) *maps.ConcurrentMap[string, uint64] {
	t.Helper()
	return maps.NewConcurrentMap[string, uint64]()
}

func TestNewConcurrentMap(t *testing.T) {
	cm := newMap(t)
	assert.NotNil(t, cm)
}

type getArgs struct {
	key string
}

type setArgs struct {
	key string
	val uint64
}

type argsType interface{ getArgs | setArgs }

type testCase[ArgsType argsType] struct {
	name   string
	args   ArgsType
	want   uint64
	wantOk bool
}

func TestConcurrentMap_Get(t *testing.T) {
	tests := []testCase[getArgs]{
		{
			name:   "valid_key",
			args:   getArgs{key: "1"},
			want:   1,
			wantOk: true,
		},
		{
			name:   "invalid_key",
			args:   getArgs{key: "2"},
			want:   0,
			wantOk: false,
		},
	}

	cm := newMap(t)
	cm.Set("1", 1)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := cm.Get(tt.args.key)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantOk, ok)
		})
	}
}

func TestConcurrentMap_Concurrent_Get(t *testing.T) {
	tests := []testCase[getArgs]{
		{
			name:   "valid_key",
			args:   getArgs{key: "1"},
			want:   1,
			wantOk: true,
		},
		{
			name:   "invalid_key",
			args:   getArgs{key: "2"},
			want:   0,
			wantOk: false,
		},
	}

	cm := newMap(t)
	cm.Set("1", 1)

	var wg sync.WaitGroup
	for i := range tests {
		wg.Add(1)
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T) {
			go func(t *testing.T, wg *sync.WaitGroup, tc testCase[getArgs]) {
				t.Helper()
				defer wg.Done()
				got, ok := cm.Get(tc.args.key)
				assert.Equal(t, tc.want, got)
				assert.Equal(t, tc.wantOk, ok)
			}(t, &wg, tc)
		})
	}

	wg.Wait()
}

func TestConcurrentMap_HasKey(t *testing.T) {
	tests := []testCase[getArgs]{
		{
			name:   "valid_key",
			args:   getArgs{key: "1"},
			wantOk: true,
		},
		{
			name:   "invalid_key",
			args:   getArgs{key: "2"},
			wantOk: false,
		},
	}

	cm := newMap(t)
	cm.Set("1", 1)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := cm.Get(tt.args.key)
			assert.Equal(t, tt.wantOk, got)
		})
	}
}

func TestConcurrentMap_Set(t *testing.T) {
	tests := []testCase[setArgs]{
		{
			name:   "set/1",
			args:   setArgs{key: "1", val: 1},
			want:   1,
			wantOk: true,
		},
		{
			name:   "set/hello",
			args:   setArgs{key: "hello", val: 42},
			want:   42,
			wantOk: true,
		},
	}

	cm := newMap(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm.Set(tt.args.key, tt.args.val)
			check, checkOk := cm.Get(tt.args.key)
			assert.Equal(t, tt.want, check)
			assert.Equal(t, tt.wantOk, checkOk)
		})
	}
}
