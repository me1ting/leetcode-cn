struct Solution {}

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
// 参考题解：https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-ii-by-l-0nmp/
//
// 二分搜索是最基本的想法，但是难点在于：如何罗列并处理所有情况，使用递归还是迭代。
// 存在一种特殊情况：nums[l] == nums[mid] == nums[r]，此时我们让l++,r--即可，最差不过退化到O(n)
// 其它情况下：必然存在nums[l]<=nums[mid] || nums[mid]<=nums[r]，此时我们可以丢弃掉一半的待查找数据
//
impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> bool {
        let n = nums.len();
        if n == 0 {
            return false;
        }

        if n == 1 {
            return nums[0] == target;
        }

        let (mut l, mut r) = (0, n - 1);
        while l < r {
            let mid = (l + r) / 2;
            if nums[mid] == target {
                return true;
            }

            if nums[l] == nums[mid] && nums[mid] == nums[r] {
                l += 1;
                r -= 1;
            } else if nums[l] <= nums[mid] {
                if nums[l] <= target && target <= nums[mid] {
                    return match nums[l..=mid].binary_search(&target) {
                        Ok(_) => true,
                        Err(_) => false,
                    };
                } else {
                    l = mid + 1;
                }
            } else if nums[mid] <= nums[r] {
                if nums[mid] <= target && target <= nums[r] {
                    return match nums[mid..=r].binary_search(&target) {
                        Ok(_) => true,
                        Err(_) => false,
                    };
                } else {
                    r = mid - 1;
                }
            }
        }

        nums[l] == target
    }
}

#[cfg(test)]
mod tests {
    use crate::search::search_in_rotated_sorted_array_ii::Solution;
    #[test]
    fn test_search() {
        let nums = vec![3, 1, 1];
        let target = 0;
        assert!(!Solution::search(nums, target));
    }
}
