pub struct Solution {}

// https://leetcode-cn.com/problems/subarrays-with-k-different-integers/
// 滑动窗口法：
// 对于某一个右边界，满足k个不同元素的左边界必然构成某一个连续区间
// 我们维护三个边界：l1,l2,r
// r随着遍历递增
// [l1,l2)为当前r时，满足k个不同元素的左边界的连续区间
// 当r移动时，l1,l2可能会向右移动，只要其不满足性质，就需要移动直到满足性质
//
// 题目条件`1 <= A[i] <= A.length`意味着可以使用桶计数
// 这里为了通用性使用的是hashmap，因此效率可能相对慢了几倍
impl Solution {
    pub fn subarrays_with_k_distinct(a: Vec<i32>, k: i32) -> i32 {
        use std::collections::HashMap;
        let mut count: i32 = 0;

        let mut l1: usize = 0;
        let mut l2: usize = 0;

        let mut k_ns1 = HashMap::new();
        let mut k_ns2 = HashMap::new();

        for r in 0..a.len() {
            let i = a[r];
            let i_n = k_ns1.get(&i).unwrap_or(&0);
            k_ns1.insert(i, i_n + 1);
            let i_n = k_ns2.get(&i).unwrap_or(&0);
            k_ns2.insert(i, i_n + 1);

            while k_ns1.len() > k as usize {
                let i = a[l1];
                let i_n = k_ns1.get(&i).unwrap_or(&0);
                if *i_n == 1 {
                    k_ns1.remove(&i);
                } else {
                    k_ns1.insert(i, i_n - 1);
                }
                l1 += 1;
            }

            while k_ns2.len() >= k as usize {
                let i = a[l2];
                let i_n = k_ns2.get(&i).unwrap_or(&0);
                if *i_n == 1 {
                    k_ns2.remove(&i);
                } else {
                    k_ns2.insert(i, i_n - 1);
                }
                l2 += 1;
            }
            count += (l2 - l1) as i32;
        }

        count
    }
}

#[cfg(test)]
mod tests {
    use crate::array::subarrays_with_k_different_integers;
    #[test]
    fn it_works() {
        let a = vec![1, 2, 1, 2, 3];
        let k = 2;
        let expected = 7;
        assert_eq!(
            expected,
            subarrays_with_k_different_integers::Solution::subarrays_with_k_distinct(a, k)
        );

        let a = vec![1, 1, 1, 1, 1];
        let k = 1;
        let expected = 15;
        assert_eq!(
            expected,
            subarrays_with_k_different_integers::Solution::subarrays_with_k_distinct(a, k)
        );
    }
}
