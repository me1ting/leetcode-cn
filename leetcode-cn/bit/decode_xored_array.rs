struct Solution();

// https://leetcode-cn.com/problems/decode-xored-array/
impl Solution {
    pub fn decode(encoded: Vec<i32>, first: i32) -> Vec<i32> {
        let mut ans = Vec::with_capacity(encoded.len() + 1);
        ans.push(first);

        let mut prev = first;
        for cipher in encoded.into_iter() {
            let plain = cipher ^ prev;
            ans.push(plain);
            prev = plain;
        }

        ans
    }
}
