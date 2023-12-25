package contains_duplicate_iii

import (
	"testing"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	return containsNearbyAlmostDuplicate3(nums, k, t)
}

// 问题：https://leetcode-cn.com/problems/contains-duplicate-iii
// 官方题解：https://leetcode-cn.com/problems/contains-duplicate-iii/solution/cun-zai-zhong-fu-yuan-su-iii-by-leetcode/

// 穷举法，或者称为滑动窗口
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j, limit := i+1, i+k; j < len(nums) && j <= limit; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// 使用二叉树：
// 初始化一颗空的二叉搜索树 set
// 对于每个元素x，遍历整个数组
//
//    在 set 上查找大于等于x的最小的数，如果s−x≤t则返回 true
//    在 set 上查找小于等于x的最大的数，如果x−g≤t则返回 true
//    在 set 中插入x
//    如果树的大小超过了k, 则移除最早加入树的那个数。
//
// 返回 false
//
// 这种方法的疑惑点在于是否存在遗漏？
// 因为set保存的是其前k个数，这保证了距离k；
// 其次是判断大于等于x的最小的数以及小于等于x的最大的数与x的差的绝对值是否小于等于t，因此不存在遗漏
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	// TODO golang标准库没有treeset，可以实现一个treeset
	return false
}

// 使用桶
// 将数据划分为...[0~(t+1)),[(t+1)~2(t+1)),[2(t+1)~3(t+1))...
// 假设nums[i]位于桶m，只需要判断桶m中是否存在元素,桶m-1,桶m+1是存在满足条件的元素
func containsNearbyAlmostDuplicate3(nums []int, k int, t int) bool {
	buckets := map[int]map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		bucketNum := bucketOf(num, t)

		if bucket, ok := buckets[bucketNum]; ok {
			if len(bucket) > 0 {
				return true
			}
		} else {
			buckets[bucketNum] = map[int]int{}
		}

		buckets[bucketNum][num]++

		if bucket, ok := buckets[bucketNum-1]; ok {
			for n, _ := range bucket {
				if abs(n-num) <= t {
					return true
				}
			}
		}
		if bucket, ok := buckets[bucketNum+1]; ok {
			for n, _ := range bucket {
				if abs(n-num) <= t {
					return true
				}
			}
		}

		if j := i - k; j >= 0 {
			num := nums[j]
			bucket:=buckets[bucketOf(num, t)]
			count:= bucket[num]
			if count <= 1 {
				delete(bucket, num)
			}else{
				bucket[num] = count-1
			}
		}
	}
	return false
}

func bucketOf(n, t int) int {
	if n > 0 {
		return n / (t + 1)
	}
	return n/(t+1) - 1
}

func TestContains(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	tt := 0
	if !containsNearbyAlmostDuplicate3(nums, k, tt) {
		t.Errorf("expected true but return false")
	}
}
