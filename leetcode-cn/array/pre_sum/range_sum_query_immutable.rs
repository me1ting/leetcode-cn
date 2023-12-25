struct NumArray {
    sums: Vec<i32>,
}

// https://leetcode-cn.com/problems/range-sum-query-immutable/
// 了解前缀和
impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        //使用n+1个空间，位置n表示nums[0]~nums[n-1]的前缀和，更加方便计算
        let mut sums: Vec<i32> = vec![0; nums.len() + 1];
        for i in 0..nums.len() {
            sums[i + 1] = sums[i] + nums[i];
        }
        NumArray { sums }
    }

    fn sum_range(&self, i: i32, j: i32) -> i32 {
        self.sums[j as usize + 1] - self.sums[i as usize]
    }
}

#[cfg(test)]
mod tests {
    use crate::array::pre_sum::range_sum_query_immutable::NumArray;

    #[test]
    fn test_sum_range() {
        let numss = vec![vec![-2, 0, 3, -5, 2, -1]];
        let range_lists = vec![vec![vec![0, 2], vec![2, 5], vec![0, 5]]];
        let expectedss = vec![vec![1, -1, -3]];

        for (i, nums) in numss.into_iter().enumerate() {
            let num_arr = NumArray::new(nums);
            let range_list = &range_lists[i];
            let expecteds = &expectedss[i];

            for (j, range) in range_list.iter().enumerate() {
                let expected = expecteds[j];
                let (l, r) = (range[0], range[1]);
                let ans = num_arr.sum_range(l, r);
                assert_eq!(ans, expected);
            }
        }
    }
}
