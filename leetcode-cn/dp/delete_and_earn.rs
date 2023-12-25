struct Solution();

// https://leetcode-cn.com/problems/delete-and-earn/
//
// 首先对数组进行统计和去重，然后排序，可以使用hashmap统计个数称为count，假设去重且排序后的数组依然称为nums
// nums[i]的点数即为nums[i]*count[nums[i]]
//
// 可以使用动态规划，但是如何找到状态和转移方程一直是动态规划的一个难点：
//
// 考虑最极端的情况下，所有数字连续，这要求我们要考虑两种可能值：奇数项的总和、偶数项的总和，这意味着只有遍历到最后一个数才知道答案
// 因此我们可以使用dp[i][0]表示获取数字nums[i]得到的最大点数，
// 使用dp[i][1]表示丢弃数字nums[i]得到的最大点数
//
// 可以使用滚动数组优化
//
// 考虑题目条件：1=<n<=10000，也可以不去重不排序，使用数字0~10000的桶来处理。
impl Solution {
    pub fn delete_and_earn(nums: Vec<i32>) -> i32 {
        use std::cmp::max;
        use std::collections::HashMap;
        let mut count: HashMap<i32, i32> = HashMap::new();

        for n in nums.iter() {
            *count.entry(*n).or_insert(0) += 1;
        }

        let mut nums: Vec<i32> = count.keys().map(|n| *n).collect();
        nums.sort();

        let mut prev_n = -1; // n>=1
        let mut contains = 0;
        let mut uncontains = 0;

        for n in nums.into_iter() {
            let c = count[&n];
            if prev_n + 1 == n {
                let temp = uncontains;
                uncontains = max(contains, uncontains);
                contains = temp + c * n;
            } else {
                uncontains = max(contains, uncontains);
                contains = uncontains + c * n;
            }
            prev_n = n;
        }

        max(contains, uncontains)
    }
}
