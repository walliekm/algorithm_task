//74. 搜索二维矩阵
//https://leetcode-cn.com/problems/search-a-2d-matrix/

package week5

//二维矩阵的每一行单调递增，第一个数大于前一行最后一个数
//可将二维数组转换为一维数组进行二分查找，二维数组的元素个数为N，时间复杂度为O(LogN)
//假如二维数组为m行n列，则二维数组的martrix[i][j]位置对应到一维数组的位置为i * n + j
//一维数组的某个位置x，映射到二维数组的martrix[i][j]，则有i = x / n, j = x % n
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)    //矩阵的行数
	n := len(matrix[0]) //矩阵的列数

	//二分查找，确定其两端位置
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		i, j := mid/n, mid%n //将一维数组的位置转换为二维数组的位置
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
