struct Solution();

// https://leetcode-cn.com/problems/single-number-ii/
// https://leetcode-cn.com/problems/single-number-ii/solution/zhi-chu-xian-yi-ci-de-shu-zi-ii-by-leetc-23t6/
// 解1：哈希表，容易理解
// 解2：单独计算每一个位的多次遍历的位运算，容易理解
// 解3：模拟数字组合电路，仅需要一次遍历，只能部分理解（理解思想，对于真值表如何转换位运算表达式完全陌生)
// TODO：补充背景知识
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let (mut a, mut b) = (0, 0);
        for n in nums {
            let (c, d) = (!a & b & n | a & !b & !n, !a & (b ^ n));
            a = c;
            b = d;
        }
        b
    }
}
