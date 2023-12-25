struct Solution {}
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
use std::cell::RefCell;
use std::rc::Rc;

use crate::sort;

// https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/
// 树中任意两不同节点值之间的最小差值：就是中序遍历结果中任意相邻节点值之间的最小差值
impl Solution {
    fn middle_order(node: &Option<Rc<RefCell<TreeNode>>>, sorted: &mut Vec<i32>) {
        if let Some(node) = node {
            let node = node.as_ref().borrow();
            Self::middle_order(&node.left, sorted);
            sorted.push(node.val);
            Self::middle_order(&node.right, sorted);
        }
    }
    pub fn min_diff_in_bst(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut sorted = Vec::new();
        Self::middle_order(&root, &mut sorted);
        let mut min_diff = i32::MAX;
        for i in 1..sorted.len() {
            let diff = sorted[i] - sorted[i - 1];
            if diff < min_diff {
                min_diff = diff;
            }
        }
        min_diff
    }
}
