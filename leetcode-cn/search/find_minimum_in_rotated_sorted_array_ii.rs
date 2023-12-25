struct Solution();
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
//
// 前一题的答案改一改，需要处理nums[l]==nums[mid]==nums[r]的情况
//
impl Solution {
    pub fn find_min(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        assert!(n > 0);

        use std::cmp::min;
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
                    if n_l == n_mid && n_mid == n_r {
                        l += 1;
                        r -= 1;
                    } else if n_mid < n_l {
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
