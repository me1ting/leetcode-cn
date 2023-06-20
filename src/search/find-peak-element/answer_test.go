package findpeakelement

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[mid+1] { //因为mid总是小于r的，因此可以放心的使用mid+1，不会出现数组越界的情况
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
