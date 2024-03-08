struct Solution {}

impl Solution {
    // https://leetcode.cn/problems/find-the-minimum-possible-sum-of-a-beautiful-array/
    pub fn minimum_possible_sum(n: i32, target: i32) -> i32 {
        const MOD: i64 = 1_000_000_007;

        let (n, target) = (n as i64, target as i64);
        let mid = target / 2;

        if n <= mid {
            return ((1 + n) * n / 2 % MOD) as i32;
        }

        let mut sum: i64 = (1 + mid) * mid / 2 % MOD;
        sum = (sum + (n - mid) * (target + target + n - mid - 1) / 2 % MOD) % MOD;

        sum as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = Solution::minimum_possible_sum(13, 50);
        assert_eq!(result, 91);
    }
}
