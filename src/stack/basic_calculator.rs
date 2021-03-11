struct Solution();

/// https://leetcode-cn.com/problems/basic-calculator/
/// 参考题解：https://leetcode-cn.com/problems/basic-calculator/solution/ji-ben-ji-suan-qi-by-leetcode-solution-jvir/
/// 考点：栈、表达式求解
///
/// 如果是基本的加减乘除，个人比较推荐使用逆波兰表达式来计算：构建抽象语法树-->后序遍历得到表达式-->使用栈求解表达式，时间复杂度O(NN^{1/2})
/// 此题由于运算符只有+,-，因此可以取巧，使用一个标记记录当前操作符，使用栈记录每一层括号的正负
impl Solution {
    pub fn calculate(s: String) -> i32 {
        let s = s.as_bytes();
        let n = s.len();

        let mut ops = vec![1]; //1表示符号+
        let mut sign = 1;

        let mut res = 0;
        let mut i = 0;
        while i < n {
            match s[i] as char {
                ' ' => {}
                '+' => {
                    sign = *ops.last().unwrap();
                }
                '-' => {
                    sign = -*ops.last().unwrap();
                }
                '(' => {
                    ops.push(sign);
                }
                ')' => {
                    ops.pop();
                }
                _ => {
                    let mut n_i = s[i];
                    let mut num = 0;
                    while '0' as u8 <= n_i && n_i <= '9' as u8 {
                        num = num * 10 + (n_i - '0' as u8) as i32;
                        i += 1;
                        if i == n {
                            break;
                        }
                        n_i = s[i];
                    }

                    res += sign * num;
                    i -= 1;
                }
            }
            i += 1;
        }

        res
    }
}

#[cfg(test)]
mod tests {
    use crate::stack::basic_calculator::Solution;
    #[test]
    fn test_calculate() {
        let ss = vec!["1 + 1", " 2-1 + 2 ", "(1+(4+5+2)-3)+(6+8)"];
        let expectations = vec![2, 3, 23];

        for i in 2..ss.len() {
            let s = &ss[i];
            let expected = expectations[i];
            let val = Solution::calculate(s.to_string());

            assert_eq!(val, expected);
        }
    }
}
