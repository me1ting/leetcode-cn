package main

// https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/solution/tong-ji-quan-wei-1-de-zheng-fang-xing-zi-ju-zhen-2/
// 全都是套路，其实和解数学题一样，如果陷入死胡同就无法解出，而掌握套路就很容易解决
// 这种题应该使用递归求解，但困难的是如何找到推进方式，题解给出了答案：
// 1. 使用f[i][j]表示以(i,j)为右下角的正方形的个数，不考虑边界情况下，存在递归式:
//	f[i][j] = (matrix[i][j])(min(matrix[i-1][j],matrix[i-1][j-1],matrix[i][j-1])+matrix[i-1][j])

func countSquares(matrix [][]int) int {
	// 考虑性能，原地更新
	sum:=0
	for i,line:=range matrix{
		for j,point:=range line{
			if i==0||j==0{
				sum+=point
			}else{
				if point>0{
					min:=matrix[i-1][j-1]
					if min>0{
						if left:=matrix[i-1][j];left<min {
							min=left
						}
					}
					if min>0{
						if top:=matrix[i][j-1];top<min {
							min=top
						}
					}
					matrix[i][j]= min + 1
					sum+=min+1
				}
			}
		}
	}
	return sum
}
