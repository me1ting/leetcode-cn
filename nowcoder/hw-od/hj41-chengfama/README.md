# 称砝码

现有n种砝码，重量互不相等，分别为 m1,m2,m3…mn ；
每种砝码对应的数量为 x1,x2,x3...xn 。现在要用这些砝码去称物体的重量(放在同一侧)，问能称出多少种不同的重量。

## 解析

这题本身很简单，但是容易想复杂了。思路是:

- 我们遍历每一种砝码类型，每一种砝码共有0~counts(m)种组合，我们尝试每种组合，来更新现有的weightSums，记录下来
- 使用set去重

```go
func Answer2(types int, weights []int, counts []int) int {
	records := map[int]bool{0: true}

	for i := 0; i < types; i++ {
		newRecords := map[int]bool{}
		for j := 0; j <= counts[i]; j++ {
			for n := range records {
				newRecords[n+j*weights[i]] = true
			}
		}
		records = newRecords
	}

	return len(records)
}
```