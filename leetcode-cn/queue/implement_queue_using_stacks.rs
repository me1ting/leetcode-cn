/// https://leetcode-cn.com/problems/implement-queue-using-stacks/
/// 用两个栈，一个push，一个存储逆序的数据用于pop
struct MyQueue {
    en: Vec<i32>,
    out: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyQueue {
    /** Initialize your data structure here. */
    fn new() -> Self {
        MyQueue {
            en: Vec::new(),
            out: Vec::new(),
        }
    }

    /** Push element x to the back of queue. */
    fn push(&mut self, x: i32) {
        self.en.push(x);
    }

    /** Removes the element from in front of queue and returns that element. */
    fn pop(&mut self) -> i32 {
        if self.out.len() == 0 {
            std::mem::swap(&mut self.en, &mut self.out);
            self.out.reverse();
        }

        return self.out.pop().unwrap();
    }

    /** Get the front element. */
    fn peek(&mut self) -> i32 {
        if self.out.len() == 0 {
            std::mem::swap(&mut self.en, &mut self.out);
            self.out.reverse();
        }

        *self.out.last().unwrap()
    }

    /** Returns whether the queue is empty. */
    fn empty(&self) -> bool {
        self.en.len() == 0 && self.out.len() == 0
    }
}

#[cfg(test)]
mod tests {
    use crate::queue::implement_queue_using_stacks::MyQueue;

    #[test]
    fn test_my_queue() {
        let mut obj = MyQueue::new();
        obj.push(1);
        assert_eq!(obj.empty(), false);
        obj.push(2);
        obj.push(3);
        assert_eq!(obj.peek(), 1);
        assert_eq!(obj.pop(), 1);
        assert_eq!(obj.pop(), 2);
        obj.push(4);
        assert_eq!(obj.pop(), 3);
        assert_eq!(obj.pop(), 4);
        assert_eq!(obj.empty(), true);
    }
}
