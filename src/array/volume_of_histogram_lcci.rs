struct Solution;

// https://leetcode-cn.com/problems/volume-of-histogram-lcci/
// 参考题解：https://leetcode-cn.com/problems/volume-of-histogram-lcci/solution/zhi-fang-tu-de-shui-liang-by-leetcode-so-7rla/
// 重复题：https://leetcode-cn.com/problems/trapping-rain-water/
//
// 有多种解决该问题的算法，双指针是其中最直观的方法
// 我们从左右两边的低点往两边扫描，遇到更低的则高度差为当前位置的水量，遇到更高的则重新比较左右两边选择新扫描起点
impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
        if height.len() <= 2 {
            return 0;
        }

        let mut i = 0;
        let mut j = height.len() - 1;
        let mut l_height = height[i];
        let mut r_height = height[j];

        let mut sum = 0;
        unsafe {
            while i < j {
                if (height.get_unchecked(i) < height.get_unchecked(j)) {
                    let h = *height.get_unchecked(i);
                    if h <= l_height {
                        sum += l_height - h;
                    } else {
                        l_height = h;
                    }
                    i += 1;
                } else {
                    let h = *height.get_unchecked(j);
                    if h <= r_height {
                        sum += r_height - h;
                    } else {
                        r_height = h;
                    }
                    j -= 1;
                }
            }
        }
        sum
    }
}
