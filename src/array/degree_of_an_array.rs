struct Solution;

struct Info {
    count: i32,
    l: usize,
    r: usize,
}

impl Solution {
    // https://leetcode-cn.com/problems/degree-of-an-array/
    // 参考题解：
    // https://leetcode-cn.com/problems/degree-of-an-array/solution/shu-zu-de-du-by-leetcode-solution-ig97/
    // 使用hash表记录每个字母的次数、第一次出现位置、最后一次出现位置
    // 遍历hash表，找到出现字母次数最多的，前后出现距离最短的字母，返回最短距离
    pub fn find_shortest_sub_array(nums: Vec<i32>) -> i32 {
        use std::collections::HashMap;
        use std::cmp::min;

        let mut info_list: HashMap<i32, Info> = HashMap::new();

        let len = nums.len();
        for i in 0..len {
            let n = nums[i];
            let info = info_list.entry(n).or_insert(Info { count: 0, l: i, r: 0 });
            info.count += 1;
            info.r = i;
        }

        let mut max_count = 0;
        let mut ans = 0;
        for (n, info) in info_list {
            let count = info.count;
            let len = (info.r - info.l) as i32 + 1;
            if count > max_count {
                max_count = count;
                ans = len;
            } else if count == max_count {
                ans = min(ans, len);
            }
        }

        ans
    }
}

#[cfg(test)]
mod tests {
    use crate::array::degree_of_an_array::Solution;

    #[test]
    fn test_find_shortest_sub_array() {
        let numss = vec![
            vec![1, 2, 2, 3, 1],
        ];
        let expecteds = vec![2, 6];

        for i in 0..numss.len() {
            let nums = &numss[i];
            let mut copy = vec![];
            for n in nums {
                copy.push(*n);
            }
            let expected = expecteds[i];
            let ans = Solution::find_shortest_sub_array(copy);
            assert_eq!(ans, expected);
        }
    }
}