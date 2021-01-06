package canplaceflowers

import "testing"

// https://leetcode-cn.com/problems/can-place-flowers/
// 贪婪法
func canPlaceFlowers(flowerbed []int, n int) bool {
	//这里单独处理边界情况，当然可以和后面迭代混在一起使得代码更加简洁，但多了更多的判断，得不偿失
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		flowerbed[0] = 1
		n--
	}
	if len(flowerbed) >= 2 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		n--
	}

	if n <= 0 {
		return true
	}

	last := len(flowerbed) - 1
	for i := 1; i < last; {
		if flowerbed[i] == 1 {
			i += 2
		} else {
			if flowerbed[i+1] == 1 {
				i += 3
			} else if flowerbed[i-1] == 1 {
				i++
			} else {
				flowerbed[i] = 1
				n--
				//只有n变化时才进行检查，使得效率更高
				if n <= 0 {
					break
				}
				i += 2
			}
		}
	}

	if last > 0 && flowerbed[last-1] == 0 && flowerbed[last] == 0 {
		//这一步从代码可读性来讲需要做，从代码效率上来讲，不需要做
		//flowerbed[last]=1
		n--
	}

	return n <= 0
}

func TestIt(t *testing.T) {
	flowerbeds := [][]int{
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 0},
	}
	ns := []int{
		1,
		1,
		1,
	}

	expecteds := []bool{
		true,
		true,
		true,
	}

	for i := 0; i < len(flowerbeds); i++ {
		if canPlaceFlowers(flowerbeds[i], ns[i]) != expecteds[i] {
			t.Errorf("error")
		}
	}
}
