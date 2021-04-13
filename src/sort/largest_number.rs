struct Solution {}

//Rust版本实现，see largest-number，由Go实现的版本
impl Solution {
    pub fn largest_number(nums: Vec<i32>) -> String {
        let mut nums: Vec<Vec<u8>> = nums.iter().map(|n| n.to_string().into_bytes()).collect();
        let cmp = |a: &Vec<u8>, b: &Vec<u8>| {
            use std::cmp::Ordering::Equal;
            let val = a[..].cmp(&b[..a.len()]);
            if val != Equal || a.len() == b.len() {
                return val;
            }
            let diff_len = b.len() - a.len();
            let val = b[..diff_len].cmp(&b[a.len()..]);
            if val != Equal {
                return val;
            }

            b[diff_len..].cmp(&a[..])
        };

        nums.sort_unstable_by(|a, b| {
            if a.len() <= b.len() {
                cmp(a, b)
            } else {
                cmp(b, a).reverse()
            }
            .reverse()
        });

        if nums[0][0] == b'0' {
            return "0".to_owned();
        }

        let mut nums: Vec<String> = nums
            .into_iter()
            .map(|n| String::from_utf8(n).unwrap())
            .collect();
        nums.join("")
    }
}
