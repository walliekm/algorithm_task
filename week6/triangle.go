//120. 三角形最小路径和
//https://leetcode-cn.com/problems/triangle/

package week6

//时间复杂度O(n*n)，其中n为三角形的行数
//空间复杂度O(n*n)，需要一个n * n的二维数组存放所有状态
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		//每一行的第一个元素，只能从上一行的第一个元素过来
		f[i][0] = f[i-1][0] + triangle[i][0]

		for j := 1; j < i; j++ {
			//对于 0 < j < i的情况，可能从上一行的j或j-1位置过来，需取两者的最小值
			f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
		}

		//当j=i时，只能从上一行的i-1位置过来
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}

	ans := f[n-1][0]
	for i := 1; i < n; i++ {
		ans = min(f[n-1][i], ans)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
