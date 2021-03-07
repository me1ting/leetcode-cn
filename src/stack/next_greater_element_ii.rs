struct Solution();
/// https://leetcode-cn.com/problems/next-greater-element-ii/
/// 参考题解：https://leetcode-cn.com/problems/next-greater-element-ii/solution/xia-yi-ge-geng-da-yuan-su-ii-by-leetcode-bwam/
/// 使用技巧：单调栈
///
/// 单调栈：使用栈记录从底到顶单调递减的数组元素的索引
///        每次遍历到一个数据，出栈比它元素更小的数组元素的索引，这些位置的下一个更大元素就是当前位置上的元素
/// 循环数组：一次遍历是不够的，因为是循环数组，还需要循环遍历一次
///         事实上，位置n-1上的元素是位置0的重复，一些题型要考虑重复带来的问题。本题不需要考虑重复位置的问题，因此直接暴力做两次遍历即可
impl Solution {
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![-1; nums.len()];
        let mut stack: Vec<usize> = Vec::new();
        for c in 0..2 {
            for (i, n) in nums.iter().enumerate() {
                while stack.len() > 0 {
                    let last = stack.last().unwrap();
                    if nums[*last] < *n {
                        ans[*last] = *n;
                        stack.pop();
                    } else {
                        break;
                    }
                }
                stack.push(i);
            }
        }
        ans
    }
}

#[cfg(test)]
mod tests {
    use crate::stack::next_greater_element_ii::Solution;
    #[test]
    fn test_next_greater_elements() {
        let numss = vec![vec![1, 2, 1]];
        let expecteds = vec![vec![2, -1, 2]];

        for i in 0..numss.len() {
            let nums = numss[i].clone();
            let expected = expecteds[i].clone();
            let ans = Solution::next_greater_elements(nums);

            assert_eq!(ans, expected);
        }
    }
}
