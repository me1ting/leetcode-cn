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
      right: None
    }
  }
}

pub struct Solution {}

use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn is_same_tree(p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> bool {
      fn travel(p: &Option<Rc<RefCell<TreeNode>>>,q: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        if p.is_none()&&q.is_none() {
          return true;
        }
  
        if p.is_none()||q.is_none() {
          return false;
        }

        if let Some(p) = p {
          if let Some(q) = q {
            if p.borrow().val != q.borrow().val {
              return false;
            }else{
              return travel(&p.borrow().left, &q.borrow().left)&&travel(&p.borrow().right, &q.borrow().right);
            }
          }
        }
        
        return true
      }
      
      travel(&p,&q)
    }
}