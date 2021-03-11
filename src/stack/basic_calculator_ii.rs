struct Solution();

/// https://leetcode-cn.com/problems/basic-calculator-ii/
/// 参考题解：https://leetcode-cn.com/problems/basic-calculator-ii/solution/ji-ben-ji-suan-qi-ii-by-leetcode-solutio-cm28/
/// 由于没有括号，因此可以针对简化代码：
/// 一次遍历，可能遇到以下几种情况：
///
impl Solution {
    pub fn calculate(s: String) -> i32 {
        let n = s.len();

        let mut stack = vec![0];
        let mut num = 0;
        let mut pre_sign = b'+';

        for (i, c) in s.bytes().enumerate() {
            let is_digit = c.is_ascii_digit();
            if is_digit {
                num = num * 10 + (c - b'0') as i32;
            }

            if !is_digit && c != b' ' || i == n - 1 {
                match pre_sign {
                    b'+' => stack.push(num),
                    b'-' => stack.push(-num),
                    b'*' => *stack.last_mut().unwrap() *= num,
                    b'/' => *stack.last_mut().unwrap() /= num,
                    _ => unreachable!(),
                }
                pre_sign = c;
                num = 0;
            }
        }

        stack.into_iter().sum()
    }
}

#[cfg(test)]
mod tests {
    use crate::stack::basic_calculator_ii::Solution;

    #[test]
    fn it_works() {
        let ss = vec!["3+2*2", " 3/2 ", " 3+5 / 2 "];
        let exceptations = vec![7, 1, 5];

        for (i, s) in ss.into_iter().enumerate() {
            let excepted = exceptations[i];
            let val = Solution::calculate(s.to_string());

            assert_eq!(val, excepted);
        }
    }
}
