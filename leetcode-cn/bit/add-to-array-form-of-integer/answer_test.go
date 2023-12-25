package answer

// https://leetcode-cn.com/problems/add-to-array-form-of-integer
// 这里抄一下题解
// 原理很简单，假设A,K都是数组形式的数据，我们也只需要一个int来保存和与余数
// 当A、K长度不一致时，实践时要考虑足够精巧的实现
// 先记录的是结果的低位，因此需要反转
func addToArrayForm(A []int, K int) (ans []int) {
	for i := len(A) - 1; i >= 0 || K > 0; i-- {
		if i >= 0 {
			K += A[i]
		}
		ans = append(ans, K%10)
		K /= 10
	}
	reverse(ans)
	return
}

func reverse(A []int) {
	for i, n := 0, len(A); i < n/2; i++ {
		A[i], A[n-1-i] = A[n-1-i], A[i]
	}
}
