struct Solution();

// https://leetcode-cn.com/problems/subarray-sum-equals-k/
// 前缀和+哈希表
impl Solution {
    pub fn subarray_sum(nums: Vec<i32>, k: i32) -> i32 {
        use std::collections::HashMap;
        let mut count = 0;
        let mut pre = 0;
        let mut m: HashMap<i32, i32> = HashMap::new();
        m.insert(0, 1);

        for n in nums {
            pre += n;
            if let Some(c) = m.get(&(pre - k)) {
                count += c;
            }
            *m.entry(pre).or_insert(0) += 1;
        }
        count
    }
}
