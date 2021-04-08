struct Solution {}

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
//
// 使用二分法查找，分情况讨论：
// 1.nums[l]<nums[r]，说明当前slice是有序的，返回nums[l]
// 2.否则，查看mid，nums[mid]<nums[l]，则最小数位于nums[l..=mid]；nums[mid]>nums[r]，则最小数位于nums[mid..=r]
//
impl Solution {
    pub fn find_min(nums: Vec<i32>) -> i32 {
        use std::cmp::min;
        let n = nums.len();
        assert!(n > 0);

        if n == 1 {
            return nums[0];
        }
        if n == 2 {
            return min(nums[0], nums[1]);
        }

        let (mut l, mut r) = (0, n - 1);
        while l < r {
            let n_l = nums[l];
            let n_r = nums[r];
            if n_l < n_r {
                return nums[l];
            } else {
                if l + 1 == r {
                    return min(nums[l], nums[r]);
                } else {
                    let mid = (l + r) / 2;
                    let n_mid = nums[mid];
                    if n_mid < n_l {
                        r = mid;
                    } else {
                        l = mid;
                    }
                }
            }
        }
        nums[l]
    }
}

#[cfg(test)]
mod tests {
    use crate::search::find_minimum_in_rotated_sorted_array::Solution;
    #[test]
    fn test_search() {
        let nums = vec![2, 3, 1];
        assert_eq!(Solution::find_min(nums), 1);
        let nums = vec![5, 1, 2, 3, 4];
        assert_eq!(Solution::find_min(nums), 1);
    }
}
