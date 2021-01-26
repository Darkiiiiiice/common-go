package container

func reverse(arr []int) {

	var l = 0
	var r = len(arr) - 1

	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

