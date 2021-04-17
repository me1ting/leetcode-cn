struct Solution {}

// https://leetcode-cn.com/problems/contains-duplicate-iii/
// 参考题解：https://leetcode-cn.com/problems/contains-duplicate-iii/solution/cun-zai-zhong-fu-yuan-su-iii-by-leetcode-bbkt/
// 考点：滑动窗口、桶排序
//
// 一种朴素算法是计算所有可能的差值，时间复杂度为kN，空间复杂度为O(1)
//
// 事实上，我们可以用空间换时间来改善算法。将数域根据其值划分所属空间，使用hash表记录当前空间的一个值
// 在一次迭代所有值时，当前值可能：
// 1. 所属空间有值，返回true
// 2. 所属上一个空间有值，两者差满足条件，返回true
// 3. 所属下一个空间有值，两者差满足条件，返回true
// 4. 将当前值记录到其空间中，同时维护窗口为大小k：移动窗口并删除不在窗口内的值
//
//
// 算法时间复杂度O(N)，空间复杂度主要是hash表的大小：O(min(k,N))
//
// 可能存在溢出：转换为64位处理

impl Solution {
    fn get_id(n: i64, w: i64) -> i32 {
        if n >= 0 {
            (n / w) as i32 //0~w-1 属于 0
        } else {
            ((n + 1) / w - 1) as i32 //-w~-1 属于 -1
        }
    }
    pub fn contains_nearby_almost_duplicate(nums: Vec<i32>, k: i32, t: i32) -> bool {
        use std::collections::HashMap;
        let mut buckets: HashMap<i32, i32> = HashMap::new();
        let w: i64 = t as i64 + 1;

        for (i, n) in nums.iter().enumerate() {
            let id = Self::get_id(*n as i64, w);
            match buckets.get(&id) {
                Some(_) => return true,
                None => {
                    buckets.insert(id, *n);
                }
            }
            match buckets.get(&(id + 1)) {
                Some(m) => {
                    if *m as i64 <= *n as i64 + t as i64 {
                        return true;
                    }
                }
                None => {
                    buckets.insert(id, *n);
                }
            }
            match buckets.get(&(id - 1)) {
                Some(m) => {
                    if *n as i64 <= *m as i64 + t as i64 {
                        return true;
                    }
                }
                None => {
                    buckets.insert(id, *n);
                }
            }
            let i = i as i32;
            if i >= k {
                buckets.remove(&Self::get_id(nums[(i - k) as usize] as i64, w));
            }
        }

        false
    }
}

#[cfg(test)]
mod tests {
    use crate::array::contains_duplicate_iii::Solution;

    #[test]
    fn test() {
        let nums = vec![2147483647, -1, 2147483647];
        let k = 1;
        let t = 2147483647;

        let ans = Solution::contains_nearby_almost_duplicate(nums, k, t);
        assert!(!ans)
    }
}
