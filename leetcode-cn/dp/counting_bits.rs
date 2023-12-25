struct Solution;

// https://leetcode-cn.com/problems/counting-bits/solution/
// 我理解错了，这道题不是编写求数的比特位数的函数
// 而是遍历范围内所有数的比特位数，因此可以使用dp加速
// 参考题解：https://leetcode-cn.com/problems/counting-bits/solution/bi-te-wei-ji-shu-by-leetcode-solution-0t1i/
// 几种方法都很巧妙
// 这里使用最低设置位
//
// 测试用例表明，对于该题，直接运算的效率并没有想象中的慢，甚至还快于基于dp的实现。
impl Solution {
    pub fn count_bits(num: i32) -> Vec<i32> {
        let mut counts: Vec<i32> = vec![0; num as usize + 1];

        for i in 1..=num as usize {
            counts[i] = counts[i & (i - 1)] + 1
        }

        counts
    }
}

//原实现，直接运算

/*
impl Solution {
    pub fn count_bits(num: i32) -> Vec<i32> {
        static TABLE: [i32; 256] = [
            0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3,
            4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4,
            4, 5, 4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4,
            5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5,
            4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2,
            3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5,
            5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4,
            5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 3, 4, 4, 5, 4, 5, 5, 6,
            4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
        ];
        static MASK: usize = 0xff;
        let mut counts: Vec<i32> = vec![0; num as usize + 1];

        for (i, count) in counts.iter_mut().enumerate() {
            let mut n = i as i32;
            let mut c = 0;
            while n > 0 {
                c += TABLE[n as usize & MASK];
                n >>= 8;
            }
            *count = c;
            //std::mem::replace(count, c);
        }

        counts
    }
}
*/
