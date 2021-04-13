struct Solution {}

impl Solution {
    // https://leetcode-cn.com/problems/longest-repeating-character-replacement/
    // 滑动窗口：
    // 我们可以使用双指针维护一个窗口，初始满足：window-max_count_of_chars(window)<=k
    // 从位置0开始，我们最开始一定会得到一个初始窗口
    // 我们在移动右指针时可能遇到不满足的情况，由于我们需求的是最大窗口，
    // 我们只需要不缩小窗口的情况，如果不满足，则左右指针同时右移一步即可
    // 此时左右指针所表示的空间不一定满足条件
    pub fn character_replacement(s: String, k: i32) -> i32 {
        let mut counts = [0; 26];
        let (mut l, mut r) = (0, 0);

        let chars = s.as_bytes();
        const A: u8 = 'A' as u8;
        while r < chars.len() {
            counts[(chars[r] - A) as usize] += 1;
            r += 1;

            let mut max = &0;
            for count in &counts {
                if count > max {
                    max = count;
                }
            }

            if r - l > (max + k) as usize {
                counts[(chars[l] - A) as usize] -= 1;
                l += 1;
            }
        }

        (r - l) as i32
    }
}

#[cfg(test)]
mod tests {
    use crate::array::sliding_window::longest_repeating_character_replacement::Solution;

    #[test]
    fn test_character_replacement() {
        let ss = vec!["ABAB", "AABABBA", "AAAA", "ABAA"];
        let ks = vec![2, 1, 0, 0];
        let expectations = vec![4, 4, 4, 2];

        for i in 0..ss.len() {
            let s = ss[i];
            let k = ks[i];
            let excepted = expectations[i];
            let ans = Solution::character_replacement(String::from(s), k);
            assert_eq!(ans, excepted);
        }
    }
}
