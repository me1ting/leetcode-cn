struct NumMatrix {
    pre_sums: Vec<Vec<i32>>,
}

impl NumMatrix {
    fn new(matrix: Vec<Vec<i32>>) -> Self {
        let m = matrix.len();
        if m == 0 {
            return NumMatrix { pre_sums: vec![] };
        }
        let n = matrix[0].len();
        let mut pre_sums = vec![vec![0; n + 1]; m + 1];

        for i in 0..m {
            for j in 0..n {
                pre_sums[i + 1][j + 1] =
                    pre_sums[i][j + 1] + pre_sums[i + 1][j] - pre_sums[i][j] + matrix[i][j];
            }
        }

        NumMatrix { pre_sums }
    }

    fn sum_region(&self, row1: i32, col1: i32, row2: i32, col2: i32) -> i32 {
        let row1 = row1 as usize;
        let col1 = col1 as usize;
        let row2 = row2 as usize;
        let col2 = col2 as usize;

        return self.pre_sums[row2 + 1][col2 + 1]
            - self.pre_sums[row2 + 1][col1]
            - self.pre_sums[row1][col2 + 1]
            + self.pre_sums[row1][col1];
    }
}

#[cfg(test)]
mod tests {
    use crate::array::pre_sum::range_sum_query_2d_immutable::NumMatrix;

    #[test]
    fn test_sum_region() {
        let matrixs = vec![vec![
            vec![3, 0, 1, 4, 2],
            vec![5, 6, 3, 2, 1],
            vec![1, 2, 0, 1, 5],
            vec![4, 1, 0, 1, 7],
            vec![1, 0, 3, 0, 5],
        ]];
        let region_lists = vec![vec![vec![2, 1, 4, 3], vec![1, 1, 2, 2], vec![1, 2, 2, 4]]];
        let expectedss = vec![vec![8, 11, 12]];

        for (i, matrix) in matrixs.into_iter().enumerate() {
            let nm = NumMatrix::new(matrix);
            let region_list = &region_lists[i];
            let expecteds = &expectedss[i];

            for (j, region) in region_list.iter().enumerate() {
                let expected = expecteds[j];
                let ans = nm.sum_region(region[0], region[1], region[2], region[3]);
                assert_eq!(ans, expected);
            }
        }
    }
}
