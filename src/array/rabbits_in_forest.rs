use std::convert::TryInto;

struct Solution {}

// https://leetcode-cn.com/problems/rabbits-in-forest/
// 脑筋急转弯
// 不同的n必然是不同的颜色
// 对于相同的n，最多有(n+1)只兔子提到额外有n只相同颜色兔子
impl Solution {
    pub fn num_rabbits(answers: Vec<i32>) -> i32 {
        if answers.len() == 0 {
            return 0;
        }
        use std::collections::HashMap;
        let mut counts: HashMap<i32, i32> = HashMap::new();
        for i in answers.into_iter() {
            let count = counts.entry(i).or_insert(0);
            *count += 1;
        }

        let mut sum = 0;
        for (i, count) in counts.into_iter() {
            if i == 0 {
                sum += count;
            } else {
                sum += ((count as f32 / (i + 1) as f32).ceil() as i32) * (i + 1);
            }
        }

        sum
    }
}
