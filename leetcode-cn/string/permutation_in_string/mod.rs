pub struct Solution {}

// https://leetcode-cn.com/problems/permutation-in-string/
// 滑动窗口
use std::collections::HashMap;
use std::collections::HashSet;
impl Solution {
    pub fn check_inclusion(s1: String, s2: String) -> bool {
        if s1.len() > s2.len() {
            return false;
        }

        let s1 = s1.as_bytes();
        let s2 = s2.as_bytes();

        let mut char_set = HashSet::new();
        let mut char_count = HashMap::new();

        for b in s1 {
            char_set.insert(*b);
            let n = *char_count.get(b).unwrap_or(&0);
            char_count.insert(*b, n + 1);
        }

        let include = |record: &HashMap<u8, i32>, count: &HashMap<u8, i32>| -> bool {
            if record.len() != count.len() {
                return false;
            }

            for b in &char_set {
                let n1 = count.get(b).unwrap_or(&0);
                let n2 = record.get(b).unwrap_or(&0);

                if n1 != n2 {
                    return false;
                }
            }

            true
        };

        let mut record = HashMap::new();
        let k = s1.len();
        let diff_k = char_count.len();
        let mut l = 0;
        let mut r = 0;

        while r < s2.len() {
            if r < l + k {
                let i = s2[r];
                if char_set.contains(&i) {
                    let n = *record.get(&i).unwrap_or(&0);
                    record.insert(i, n + 1);
                } else {
                    l = r + 1;
                    record.clear();
                }
                r += 1;
            }

            if r == l + k {
                if record.len() == diff_k && include(&record, &char_count) {
                    return true;
                }

                let i = s2[l];
                let n = *record.get(&i).unwrap_or(&0);
                if n == 1 {
                    record.remove(&i);
                } else {
                    record.insert(i, n - 1);
                }

                l += 1;
            }
        }

        false
    }
}

#[cfg(test)]
mod tests {
    use crate::string::permutation_in_string;
    #[test]
    fn it_works() {
        let s1s = vec!["ab", "ab", "abcdxabcde"];
        let s2s = vec!["eidbaooo", "eidboaoo", "abcdeabcdx"];
        let expecteds = vec![true, false, true];

        for i in 2..s1s.len() {
            let s1 = s1s[i];
            let s2 = s2s[i];
            let expected = expecteds[i];
            assert_eq!(
                expected,
                permutation_in_string::Solution::check_inclusion(
                    String::from(s1),
                    String::from(s2)
                )
            );
        }
    }
}
