struct Solution();

/// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
/// 考察点：栈
/// 此题拿到的第一瞬间还以为需要什么复杂的算法，参考题解才发现自己想复杂了。
/// 删除所有相邻的重复元素至少需要遍历每个元素，即O(n)
/// 使用一个栈来模拟从左到右的“消消乐”的过程即可。
impl Solution {
    pub fn remove_duplicates(s: String) -> String {
        let mut stack = vec![];
        for c in s.bytes().into_iter() {
            if let Some(ch) = stack.last() {
                if *ch == c {
                    stack.pop();
                    continue;
                }
            }
            stack.push(c);
        }

        String::from_utf8(stack).unwrap()
    }
}

#[cfg(test)]
mod tests {
    use crate::stack::remove_all_adjacent_duplicates_in_string::Solution;
    #[test]
    fn test_remove_duplicates() {
        let ss = vec!["abbaca"];
        let expecteds = vec!["ca"];

        for (i, s) in ss.iter().enumerate() {
            let expected = expecteds[i];
            let val = Solution::remove_duplicates(s.to_string());
            assert_eq!(val, expected)
        }
    }
}
