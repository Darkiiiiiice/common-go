package math

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_IsPrime(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"custom-1",
			args{
				x: 0,
			},
			false,
		},
		{
			"custom-2",
			args{
				x: 1,
			},
			false,
		},
		{
			"custom-3",
			args{
				x: -1,
			},
			false,
		},
		{
			"custom-4",
			args{
				x: 2,
			},
			true,
		},
		{
			"custom-5",
			args{
				x: 3,
			},
			true,
		},
		{
			"custom-6",
			args{
				x: 6389,
			},
			true,
		},
		{
			"custom-7",
			args{
				x: 100,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsPrime() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Benchmark_IsPrime(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		IsPrime(rand.Int() % 100000000)
	}
}
