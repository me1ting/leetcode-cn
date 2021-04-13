struct Solution();

// https://leetcode-cn.com/problems/merge-sorted-array/
// 归并排序的子步骤
impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let m = m as usize;
        let n = n as usize;

        let mut ans: Vec<i32> = Vec::with_capacity(m + n);
        let mut i = 0_usize;
        let mut j = 0_usize;
        while i < m && j < n {
            unsafe {
                let i_n = nums1.get_unchecked(i);
                let j_n = nums2.get_unchecked(j);
                if i_n < j_n {
                    ans.push(*i_n);
                    i += 1;
                } else {
                    ans.push(*j_n);
                    j += 1;
                }
            }
        }
        while i < m {
            unsafe {
                let i_n = nums1.get_unchecked(i);
                ans.push(*i_n);
                i += 1;
            }
        }
        while j < n {
            unsafe {
                let j_n = nums2.get_unchecked(j);
                ans.push(*j_n);
                j += 1;
            }
        }
        nums1.clone_from(&ans);
    }
}
