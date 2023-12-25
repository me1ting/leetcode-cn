struct Solution();

// https://leetcode-cn.com/problems/reverse-integer/
// https://leetcode-cn.com/problems/reverse-integer/solution/zheng-shu-fan-zhuan-by-leetcode-solution-bccn/
//
// 得到避免溢出的条件需要一定的数学技巧，参考题解
//
impl Solution {
    pub fn reverse(x: i32) -> i32 {
        let mut x = x;
        let mut ans = 0;
        while x != 0 {
            if ans < i32::MIN / 10 || ans > i32::MAX / 10 {
                return 0;
            }
            let digit = x % 10;
            x /= 10;
            ans = ans * 10 + digit;
        }
        ans
    }
}
