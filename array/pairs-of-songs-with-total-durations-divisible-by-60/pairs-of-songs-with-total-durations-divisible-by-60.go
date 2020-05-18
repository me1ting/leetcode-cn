package main

import (
	"log"
)
// https://leetcode-cn.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
// 解体思路：
// 求余数并求组合数
func numPairsDivisibleBy60(time []int) int {
	// 统计每个余数出现的次数
	count := [60]int{}
	for _, t := range time {
		count[t%60]++
	}

	//1~29,31~59的和为60的组合使用乘积
	sum := 0
	for i := 1; i < 30; i++ {
		sum += count[i] * count[60-i]
	}

	// 0和30需要求组合C_n^2的结果
	// https://baike.baidu.com/item/%E6%8E%92%E5%88%97%E7%BB%84%E5%90%88/706498
	sum += count[0]*(count[0]-1)/2
	sum += count[30]*(count[30]-1)/2

	return sum
}

func main() {
	/*
		if numPairsDivisibleBy60([]int{418, 204, 77, 278, 239, 457, 284, 263, 372, 279, 476, 416, 360, 18}) != 1 {
			log.Fatal("error")
		}
	*/

	if numPairsDivisibleBy60([]int{14, 161, 302, 232, 270, 428, 155, 64, 499, 400, 25, 349, 434, 427, 5, 265, 20, 346, 463, 10, 1, 163, 189, 345, 390, 212, 498, 281, 308, 482, 447, 217, 318, 239, 374, 449, 298, 213, 2, 230, 5, 500, 300, 390, 139, 484, 464, 477, 111, 88, 93, 198, 181, 113, 393, 283, 383, 205, 42, 292, 335, 392, 384, 268, 361, 462, 413, 176, 118, 111, 143, 242, 166, 286, 177, 52, 126, 342, 415, 302, 210, 48, 369, 148, 209, 87, 212, 53, 296, 95, 97, 45, 467, 47, 373, 404, 43, 439, 19, 64, 289, 218, 268, 460, 238, 456, 490, 155, 429, 218, 301, 225, 228, 237, 453, 267, 259, 327, 427, 169, 176, 322, 216, 451, 29, 327, 404, 177, 225, 44, 248, 174, 287, 326, 441, 354, 110, 4, 226, 324, 331, 158, 454, 469, 354, 383, 336, 211, 133, 500, 233, 330, 471, 327, 426, 478, 195, 232, 163, 312, 126, 51, 161, 248, 433, 348, 464, 206, 480, 478, 98, 354, 304, 424, 338, 382, 431, 379, 194, 488, 6, 494, 394, 469, 430, 1, 207, 307, 172, 47, 269, 472, 415, 163, 309, 68, 213, 175, 343, 187, 15, 232, 429, 389, 373, 143, 146, 88, 58, 320, 116, 82, 482, 125, 343, 329, 115, 116, 183, 389, 112, 294, 74, 89, 62}) != 426 {
		log.Fatal("error")
	}
}
