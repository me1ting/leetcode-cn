struct Solution {}

// https://leetcode-cn.com/problems/search-a-2d-matrix/
// 两次二分查找
impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        match matrix.binary_search_by_key(&target, |v| *v.get(0).unwrap()) {
            Ok(i) => true,
            Err(i) => {
                if i == 0 {
                    return false;
                }
                let i = i - 1;
                let line = matrix.get(i).unwrap();
                match line.binary_search(&target) {
                    Ok(i) => true,
                    Err(_) => false,
                }
            }
        }
    }
}
