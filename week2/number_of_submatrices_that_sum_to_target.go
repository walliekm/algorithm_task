//1074. 元素和为目标值的子矩阵数量
//https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/

package week2

//将矩阵逐行枚举，假设当前枚举的是第i行
//令j>=i，将i~j行的矩阵折叠为一个一维数组
//问题转化为求一维数组内和为k的子数组个数问题
//subarraySum为560题的题解，在subarray_sum_equals_k.go中实现
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	rowsCount := len(matrix)
	colsCount := len(matrix[0])

	ans := 0
	for i := 0; i < rowsCount; i++ {
		sumArr := make([]int, colsCount)
		for j := i; j < rowsCount; j++ {
			for k := 0; k < colsCount; k++ {
				sumArr[k] += matrix[j][k]
			}

			ans += subarraySum(sumArr, target)
		}
	}

	return ans
}
