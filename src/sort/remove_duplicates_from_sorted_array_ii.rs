struct Solution {}

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.len() <= 2 {
            return nums.len() as i32;
        }

        let mut count = 1;
        let mut last = nums[0];
        let mut pos = 1;

        for i in 1..nums.len() {
            let n = nums[i];
            if n == last {
                count += 1;
                if count <= 2 {
                    nums[pos] = n;
                    pos += 1;
                } else {
                    continue;
                }
            } else {
                count = 1;
                nums[pos] = n;
                pos += 1;
                last = n;
            }
        }
        pos as i32
    }
}
