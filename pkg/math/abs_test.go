package math

import (
	"math"
	"reflect"
	"testing"
)

func Test_AbsInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"custom-1",
			args{
				x: 0,
			},
			0,
		},
		{
			"custom-2",
			args{
				x: 1,
			},
			1,
		},
		{
			"custom-3",
			args{
				x: -1,
			},
			1,
		},
		{
			"custom-4",
			args{
				x: math.MaxInt64,
			},
			math.MaxInt64,
		},
		{
			"custom-5",
			args{
				x: math.MinInt64,
			},
			math.MinInt64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AbsInt() = %d, want = %d", got, tt.want)
			}
		})
	}
}

