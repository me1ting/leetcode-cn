package answer

import (
	"testing"
)

// https://leetcode-cn.com/problems/find-majority-element-lcci/
// 主要元素占比超过一半，也就是排序后的中间元素必然是主要元素
// 因此问题可以转换为求数组的第k个大小的数
func majorityElement(nums []int) int {
	//n, count := quickSelect(nums, len(nums)/2)
	n := quickSelect2(nums, len(nums)/2)
	count := 0
	for _, e := range nums {
		if e == n {
			count++
		}
	}
	if count > len(nums)/2 {
		return n
	}
	return -1
}

//这里的i为index，从0开始
//使用的是《算法4》p189 三向切分的快速排序算法，适合重复元素较多的情况
func quickSelect(nums []int, i int) (num int, count int) {
	v := nums[0]
	//lt,gt表示小于v的不包含的右边界，大于v的不包含的左边界
	lt, gt := 0, len(nums)-1
	for i := 1; i <= gt; {
		n := nums[i]
		if n < v {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		} else if n > v {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}

	if i >= lt && i <= gt {
		return v, gt - lt + 1
	}

	if i < lt {
		return quickSelect(nums[:lt], i)
	}

	return quickSelect(nums[gt+1:], i-gt-1)
}

// 标准版quickselect
// 《算法4》2.3 快速排序 p184
func quickSelect2(nums []int, index int) int {
	v := nums[0]
	i, j := 1, len(nums)-1
	for {
		for i < len(nums) && nums[i] < v {
			i++
		}
		for j >= 0 && nums[j] > v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	if j < 0 {
		j = 0
	} else if j > 0 {
		nums[0], nums[j] = nums[j], nums[0]
	}

	if j == index {
		return v
	}

	if j < index {
		return quickSelect2(nums[j+1:], index-j-1)
	}

	return quickSelect2(nums[:j], index)
}

func TestIt(t *testing.T) {
	nums := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	res := majorityElement(nums)
	expected := 5
	if res != expected {
		t.Errorf("%v,%v", expected, res)
	}

	nums = []int{3, 2}
	res = majorityElement(nums)
	expected = -1
	if res != expected {
		t.Errorf("%v,%v", expected, res)
	}

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	res = majorityElement(nums)
	expected = 2
	if res != expected {
		t.Errorf("%v,%v", expected, res)
	}
}
