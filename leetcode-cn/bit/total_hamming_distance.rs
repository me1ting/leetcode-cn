struct Solution();

// https://leetcode-cn.com/problems/total-hamming-distance/
// 技巧1：位于第i位，只需统计为1的为m，那么为0的自然为n-m，根据数学原理，  汉明距离（位不同的总数）为C_m^1*C_(n-m)^1 = m(n-m)
// 技巧2：10^9<2^30，但直接按照32位来思考也并没有什么问题
//
//
impl Solution {
    pub fn total_hamming_distance(nums: Vec<i32>) -> i32 {
        let mut count: Vec<i32> = vec![0; 30];
        let nums_len = nums.len() as i32;

        for n in nums.into_iter() {
            let mut mask = 1;
            for j in 0..30 {
                if n & mask > 0 {
                    count[j] += 1;
                }
                mask <<= 1;
            }
        }

        let mut ans = 0;
        for c in count.into_iter() {
            ans += c * (nums_len - c);
        }

        ans
    }
}

#[cfg(test)]
mod tests {
    use crate::bit::total_hamming_distance::Solution;

    #[test]
    fn test() {
        let nums = vec![4, 14, 2];
        let expected = 6;

        let ans = Solution::total_hamming_distance(nums);
        assert_eq!(ans, expected)
    }
}
