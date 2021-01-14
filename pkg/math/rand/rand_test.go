package rand

import (
	std_rand "math/rand"
	"testing"
	"time"
)

// 测试
func Test_RandomSliceInt(t *testing.T) {

	t.Run("RandomSliceInt", func(t *testing.T) {
		var length = std_rand.Int() % 10000
		var lo = std_rand.Int() % 1000000
		var hi = std_rand.Int() % 10000000

		t.Logf("length = %d, lo = %d, hi = %d", length, lo, hi)
		if lo > hi {
			lo, hi = hi, lo
		}

		t.Logf("length = %d, lo = %d, hi = %d", length, lo, hi)

		var nums = RandomSliceInt(length, lo, hi)
		checkSliceInt(t, nums, length, lo, hi)
	})

}

func checkSliceInt(t *testing.T, nums []int, length, lo, hi int) {
	if len(nums) != length {
		t.Errorf("len(nums)  = %d, want = %d", len(nums), length)
	}
	for i, v := range nums {
		if v < lo {
			t.Errorf("nums[%d]  = %d less than lo = %d", i, nums[i], lo)
		}

		if v > hi {
			t.Errorf("nums[%d]  = %d greater than hi = %d", i, nums[i], hi)
		}
	}
}


func Benchmark_UniformInt(b *testing.B) {
	std_rand.Seed(time.Now().UnixNano())
	
	for i:=0; i< b.N; i++ {
		var lo = std_rand.Int()
		var hi = std_rand.Int()
		UniformInt(lo, hi)
	}
}


func Benchmark_Discrete(b *testing.B) {
	std_rand.Seed(time.Now().UnixNano())
	


	for i:=0; i< b.N; i++ {
		var lo = UniformInt(0, 1000000)
		var hi = UniformInt(0, 1000000)
		var length = UniformInt(0, 1000)
		var nums = RandomSliceInt(length, lo, hi)
		Discrete(nums)
	}
}

func Benchmark_Shuffle(b *testing.B) {
	std_rand.Seed(time.Now().UnixNano())

	for i:=0; i< b.N; i++ {
		var lo = UniformInt(0, 1000000)
		var hi = UniformInt(0, 1000000)
		var length = UniformInt(0, 1000)
		var nums = RandomSliceInt(length, lo, hi)
		Shuffle(nums)
	}
}
