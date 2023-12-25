pub struct Solution {}

#[derive(Debug, PartialEq, Eq)]
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

// https://leetcode-cn.com/problems/balanced-binary-tree/
// 深度优先搜索，每个根节点返回其深度，并由根节点判断其两个子树是否平衡
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn is_balanced(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let (is, depth) = Solution::is_balanced_and_depth(&root);
        is
    }

    fn is_balanced_and_depth(node: &Option<Rc<RefCell<TreeNode>>>) -> (bool, i32) {
        match node {
            Some(val) => {
                let (left_balance, left_depth) =
                    Solution::is_balanced_and_depth(&val.borrow().left);
                if !left_balance {
                    return (false, -1);
                }
                let (right_balance, right_depth) =
                    Solution::is_balanced_and_depth(&val.borrow().right);
                if !right_balance {
                    return (false, -1);
                }

                if left_depth - right_depth > 1 || left_depth - right_depth < -1 {
                    return (false, -1);
                }
                return (true, Solution::max(left_depth, right_depth) + 1);
            }
            None => (true, 0),
        }
    }

    #[inline]
    fn max(x: i32, y: i32) -> i32 {
        if x > y {
            return x;
        }
        y
    }
}
