package main

// https://leetcode-cn.com/problems/container-with-most-water/
// 第一眼，个人认为和best-sightseeing-pair一样，是一个吸引力随距离改变的问题，因此O(n)应该可以解决问题，但后来发现该思路是错误的
// 因为究竟哪个更好，与第二根垂线的高度有关。
//
// 然后个人想到了双指针法，但对于具体怎么移动，以及正确性无法判断。经过题解解释想通了：
// 移动短的那一根，因为木桶效应，容量受短板限制，只有移动短的那一根才能改善容量，而移动长的那一根无法改变，只会缩小。

func maxArea(height []int) int {
	max:=0
	for i,j:=0, len(height)-1;i<j;{
		lH,rH:=height[i],height[j]
		minH := lH
		if rH < lH{
			minH = rH
		}
		if cap:=minH*(j-i);cap > max {
			max = cap
		}

		if lH < rH {
			i++
		}else{
			j--
		}
	}
	return max
}