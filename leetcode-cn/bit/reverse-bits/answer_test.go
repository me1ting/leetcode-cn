package answer

// https://leetcode-cn.com/problems/reverse-bits/
// 参考题解：https://leetcode-cn.com/problems/reverse-bits/solution/dian-dao-er-jin-zhi-wei-by-leetcode-solu-yhxz/
// 算法：分治位翻转
// 技巧：如何对比特位进行位翻转

const (
	m1  = 0b01010101010101010101010101010101
	m2  = 0b00110011001100110011001100110011
	m4  = 0b00001111000011110000111100001111
	m8  = 0b00000000111111110000000011111111
	m16 = 0b00000000000000001111111111111111
)

func reverseBits(num uint32) uint32 {
	num = num>>1&m1 | num&m1<<1
	num = num>>2&m2 | num&m2<<2
	num = num>>4&m4 | num&m4<<4
	num = num>>8&m8 | num&m8<<8
	//num = num>>16&m16 | num&m16<<16 优化，不需要&运算
	return num>>16 | num<<16
}
