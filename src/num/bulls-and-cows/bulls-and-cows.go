package bulls_and_cows

import "fmt"

// https://leetcode-cn.com/problems/bulls-and-cows/
// 统计匹配位的个数并将不匹配的数字及其个数进行统计
// 最后将0~9不匹配的数字取最小值求和
func getHint(secret string, guess string) string {
	sCount, gCount := [10]int{}, [10]int{}
	x, y := 0, 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			x++
		} else {
			sCount[secret[i]-'0']++
			gCount[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		if sCount[i] < gCount[i] {
			y += sCount[i]
		} else {
			y += gCount[i]
		}
	}
	return fmt.Sprintf("%dA%dB", x, y)
}
