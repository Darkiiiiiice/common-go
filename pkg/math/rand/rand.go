package rand

import (
	std_rand "math/rand"
	"time"
)

// UniformInt
// 返回 [lo...hi) 范围之间的int类型数据
func UniformInt(lo, hi int) int {
	return lo + std_rand.Int()%(hi-lo)
}

// DDiscrete
// 根据离散概率随机返回 int
// 离散概率 i = nums[i]
func Discrete(nums []int) int {
	var r = UniformInt(0,101)
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum >= r {
			return i
		}
	}
	return -1
}

// SShuffle
// 随机将 nums 中的值排序
func Shuffle(nums []int) {

	var n = len(nums)
	for i := 0; i < n; i++ {
		var r = i + UniformInt(0, n-1)
		nums[i], nums[r] = nums[r], nums[i]
	}
}

// RandomSliceInt
// 随机生成 长度为 length , 范围在[lo, hi] 之间的Int数组切片
func RandomSliceInt(length int, lo, hi int) []int {
	// 初始化随机种子
	std_rand.Seed(time.Now().UnixNano())

	var nums = make([]int, length)

	for i := 0; i < length; i++ {
		nums[i] = UniformInt(lo, hi) 
	}

	return nums
}

