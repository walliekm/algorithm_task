//47. 全排列 II
//https://leetcode-cn.com/problems/permutations-ii/

package week3

import "sort"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	used := make([]bool, n)  //标记每个元素的使用情况
	arr := make([]int, 0, n) //进入排列的元素列表

	//对数组排序，方便后续去重判断
	sort.Ints(nums)
	var recurFunc func(i int)
	recurFunc = func(i int) {
		if i == n {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		for j := 0; j < n; j++ {
			//如果当前元素已被使用，则跳过
			//或者当前元素与其前一元素相等，并且其前一元素尚未使用，也跳过
			//确保重复元素是由左至右依次填入的，以此去重
			if used[j] || (j > 0 && nums[j] == nums[j-1] && !used[j-1]) {
				continue
			}
			used[j] = true
			arr = append(arr, nums[j])
			recurFunc(i + 1)
			arr = arr[:len(arr)-1]
			used[j] = false
		}
	}

	recurFunc(0)
	return ans
}
