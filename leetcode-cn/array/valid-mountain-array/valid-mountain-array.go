package valid_mountain_array

// https://leetcode-cn.com/problems/valid-mountain-array
// 一个简单题，主要是实现的方式，一开始我设想找上坡下坡的状态
// 后来发现使用状态迭代更加轻松，停止状态不对则说明不是有效的山脉数组
func validMountainArray(A []int) bool {
	if len(A) >= 3 {
		i := 1
		for ; i < len(A) && A[i-1] < A[i]; i++ {
		}
		if i == 1 || i == len(A) || A[i-1] == A[i] {
			return false
		}

		for i = i + 1; i < len(A) && A[i-1] > A[i]; i++ {
		}

		if i == len(A) {
			return true
		} else {
			return false
		}

	}
	return false
}
