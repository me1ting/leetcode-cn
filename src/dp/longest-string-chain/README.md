# 简介
原题：https://leetcode.cn/problems/longest-string-chain/submissions/

# 解析
字符串a如果是支付b的前身，那么：

- `len(a)+1==len(b)`

可以使用简单的办法验证前身关系：
```go
func isPre(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) && a[i] == b[j] {
		i++
		j++
	}
    //仅需要跳过一次，其它部分的字符都是一一相同的
	j++
	for i < len(a) && j < len(b) && a[i] == b[j] {
		i++
		j++
	}
	return i == len(a) && j == len(b)
}
```

我们可以将所有字符串排序后，使用动态规划解决问题，`dp[i]`表示`word[0:i+1]`以`word[i]`为结尾的链的最大长度。那么：
```go
dp[i] = max(dp[i],dp[j]+1)//其中，word[j]是word[i]的前身
```
最后，dp中的最大值即为本题结果。